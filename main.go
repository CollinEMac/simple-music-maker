package main

import (
	"fmt"
	"os/exec"
	"sync"

	"math"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
)

const rate = 44100

var SampleRate = beep.SampleRate(rate)

func main() {
	// Initialize the speaker with a buffer size
	speaker.Init(SampleRate, SampleRate.N(time.Second/10))

	var wg sync.WaitGroup
	wg.Add(2)

	// // Play the kick drum in a separate goroutine
	// go func() {
	// 	speaker.Play(beep.Seq(
	// 		beep.Take(SampleRate.N(125*time.Millisecond), KickDrum()),
	// 		generators.Silence(SampleRate.N(125*time.Millisecond)),
	// 		beep.Take(SampleRate.N(125*time.Millisecond), KickDrum()),
	// 		generators.Silence(SampleRate.N(125*time.Millisecond)),
	// 		beep.Take(SampleRate.N(125*time.Millisecond), KickDrum()),
	// 		generators.Silence(SampleRate.N(125*time.Millisecond)),
	// 		beep.Take(SampleRate.N(125*time.Millisecond), KickDrum()),
	// 		generators.Silence(SampleRate.N(125*time.Millisecond)),
	// 		beep.Take(SampleRate.N(125*time.Millisecond), KickDrum()),
	// 		generators.Silence(SampleRate.N(125*time.Millisecond)),
	// 		beep.Take(SampleRate.N(125*time.Millisecond), KickDrum()),
	// 		generators.Silence(SampleRate.N(125*time.Millisecond)),
	// 		beep.Take(SampleRate.N(125*time.Millisecond), KickDrum()),
	// 		generators.Silence(SampleRate.N(125*time.Millisecond)),
	// 		beep.Take(SampleRate.N(125*time.Millisecond), KickDrum()),
	// 		generators.Silence(SampleRate.N(125*time.Millisecond)),
	// 		beep.Take(SampleRate.N(125*time.Millisecond), KickDrum()),
	// 		generators.Silence(SampleRate.N(125*time.Millisecond)),
	// 		beep.Take(SampleRate.N(125*time.Millisecond), KickDrum()),
	// 		generators.Silence(SampleRate.N(125*time.Millisecond)),
	// 		beep.Take(SampleRate.N(125*time.Millisecond), KickDrum()),
	// 		generators.Silence(SampleRate.N(125*time.Millisecond)),
	// 		beep.Callback(func() {
	// 			wg.Done()
	// 		}),
	// 	))
	// }()

	// Play the synths in the main goroutine
	speaker.Play(beep.Seq(
		beep.Take(SampleRate.N(1*time.Second), Chiptune(440, 0.5)),    // A4 note
		beep.Take(SampleRate.N(1*time.Second), Chiptune(494, 0.3)),    // B4 note (with a different pulse width)
		beep.Take(SampleRate.N(1*time.Second), Chiptune(523.25, 0.5)), // C5 note
	))

	wg.Done()

	// Wait for both goroutines to finish
	wg.Wait()

	// Close the speaker
	speaker.Close()
}

func KickDrum() beep.Streamer {
	return beep.StreamerFunc(func(samples [][2]float64) (int, bool) {
		for i := range samples {
			// Sharp attack
			attack := float64(i) / float64(len(samples))
			if attack < 0.01 { // Short attack
				samples[i][0] = 1.0 // Volume
				samples[i][1] = 1.0 // Volume
			} else {
				// Decay and sustain
				samples[i][0] = 0.5 * math.Exp(-attack*20) // Decreasing amplitude over time
				samples[i][1] = 0.5 * math.Exp(-attack*20) // Decreasing amplitude over time
			}

			// Add a low-frequency sine wave to simulate the "boom" of an 808 drum beat
			samples[i][0] += 0.25 * math.Sin(2*math.Pi*100*attack)
			samples[i][1] += 0.25 * math.Sin(2*math.Pi*100*attack)
		}
		return len(samples), true
	})
}

func Chiptune(freq float64, pulseWidth float64) beep.Streamer {
	phase := 0.0
	return beep.StreamerFunc(func(samples [][2]float64) (n int, ok bool) {
		for i := range samples {
			// Square wave
			squareWave := 0.0
			if phase < pulseWidth {
				squareWave = 1.0
			} else {
				squareWave = -1.0
			}

			// Pulse wave
			pulseWave := 0.0
			if phase < pulseWidth {
				pulseWave = 1.0
			}

			// Mix square wave and pulse wave
			samples[i][0] = squareWave*0.5 + pulseWave*0.5
			samples[i][1] = samples[i][0]

			phase += freq / float64(beep.SampleRate(rate))
			if phase >= 1.0 {
				phase -= 1.0
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
