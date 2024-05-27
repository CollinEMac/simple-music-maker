package main

import (
	"math"
	"math/rand"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
)

var SampleRate = beep.SampleRate(44100)

func main() {
	speaker.Init(SampleRate, SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(beep.Take(SampleRate.N(2*time.Second), SawtoothWave(440)), beep.Callback(func() {
		done <- true
	})))
	<-done
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
