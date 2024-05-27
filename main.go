package main

import (

	"fmt"
	"os/exec"
	
	
	"math"
	"math/rand"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
)

var SampleRate = beep.SampleRate(44100)

func main() {
	// // Initialize the speaker with a buffer size
	speaker.Init(SampleRate, SampleRate.N(time.Second/10))

	// Define the kick drum pattern (every second)
	kickPattern := LoopStreamer(beep.Seq(
		beep.Take(SampleRate.N(1*time.Second), KickDrum()),
		beep.Silence(SampleRate.N(1*time.Second)),
	))

	// Define the wave patterns to play sequentially
	wavePatterns := beep.Seq(
		beep.Take(SampleRate.N(2*time.Second), SineWave(400)),
		beep.Take(SampleRate.N(2*time.Second), SineWave(440)),
		beep.Take(SampleRate.N(2*time.Second), SineWave(480)),
	)

	// Mix the kick pattern with the wave patterns
	music := beep.Mix(kickPattern, wavePatterns)

	// Play the mixed music
	speaker.Play(music)

	// Wait for the music to finish
	time.Sleep(10 * time.Second)

	// Close the speaker
	speaker.Close()
}

// LoopStreamer is a custom streamer that loops the given streamer indefinitely
func LoopStreamer(s beep.Streamer) beep.Streamer {
	buf := make([][2]float64, 44100) // Buffer for 1 second of audio
	_, ok := s.Stream(buf)
	if !ok {
		return beep.StreamerFunc(func(samples [][2]float64) (int, bool) {
			return 0, false
		})
	}

	return beep.StreamerFunc(func(samples [][2]float64) (n int, ok bool) {
		for i := range samples {
			samples[i] = buf[i%len(buf)]
		}
		return len(samples), true
	})
}

func Noise() beep.Streamer {
	return beep.StreamerFunc(func(samples [][2]float64) (n int, ok bool) {
		for i := range samples {
			samples[i][0] = rand.Float64()*2 - 1
			samples[i][1] = rand.Float64()*2 - 1
		}
		return len(samples), true
	})
}

func SineWave(freq float64) beep.Streamer {
	phase := 0.0
	return beep.StreamerFunc(func(samples [][2]float64) (n int, ok bool) {
		for i := range samples {
			samples[i][0] = math.Sin(2 * math.Pi * phase)
			samples[i][1] = samples[i][0]
			phase += freq / float64(beep.SampleRate(44100))
			if phase >= 1.0 {
				phase -= 1.0
			}
		}
		return len(samples), true
	})
}

func SquareWave(freq float64) beep.Streamer {
	phase := 0.0
	return beep.StreamerFunc(func(samples [][2]float64) (n int, ok bool) {
		for i := range samples {
			step := int(math.Round(phase))
			if step == 0 {
				samples[i][0] = -1
			} else {
				samples[i][0] = 1
			}
			samples[i][1] = samples[i][0]
			phase += freq / float64(beep.SampleRate(44100))
			if phase >= 1.0 {
				phase -= 1.0
			}
		}
		return len(samples), true
	})
}

func SawtoothWave(freq float64) beep.Streamer {
	phase := 0.0
	return beep.StreamerFunc(func(samples [][2]float64) (n int, ok bool) {
		for i := range samples {
			samples[i][0] = phase
			samples[i][1] = samples[i][0]
			phase += freq / float64(beep.SampleRate(44100))
			if phase >= 1.0 {
				phase -= 1.0
			}
		}
		return len(samples), true
	})
}

func KickDrum() beep.Streamer {
	return beep.StreamerFunc(func(samples [][2]float64) (int, bool) {
		for i := range samples {
			// Sharp attack
			attack := float64(i) / float64(len(samples))
			if attack < 0.01 { // Short attack
				samples[i][0] = 0.5 // Volume
				samples[i][1] = 0.5 // Volume
			} else {
				// Decay and sustain
				samples[i][0] = 0.25 * math.Exp(-attack*20) // Decreasing amplitude over time
				samples[i][1] = 0.25 * math.Exp(-attack*20) // Decreasing amplitude over time
			}
		}
		return len(samples), true
	})
}

func Vocals() {
	text := "Welcome to the 1 X Developer Podcast"
	cmd := exec.Command("espeak", text)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
	}
}