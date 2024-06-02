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

var keys = map[string]float32{
	"A0":  27.50000,
	"A0#": 29.13524,
	"B0":  30.86771,
	"C1":  32.70320,
	"C1#": 34.64783,
	"D1":  36.70810,
	"D1#": 38.89087,
	"E1":  41.20344,
	"F1":  43.65353,
	"F1#": 46.24930,
	"G1":  48.99943,
	"G1#": 51.91309,
	"A1":  55.00000,
	"A1#": 58.27047,
	"B1":  61.73541,
	"C2":  65.40639,
	"C2#": 69.29566,
	"D2":  73.41619,
	"D2#": 77.78175,
	"E2":  82.40689,
	"F2":  87.30706,
	"F2#": 92.49861,
	"G2":  97.99886,
	"G2#": 103.8262,
	"A2":  110.0000,
	"A2#": 116.5409,
	"B2":  123.4708,
	"C3":  130.8128,
	"C3#": 138.5913,
	"D3":  146.8324,
	"D3#": 155.5635,
	"E3":  164.8138,
	"F3":  174.6141,
	"F3#": 184.9972,
	"G3":  195.9977,
	"G3#": 207.6523,
	"A3":  220.0000,
	"A3#": 233.0819,
	"B3":  246.9417,
	"C4":  261.6256,
	"C4#": 277.1826,
	"D4":  293.6648,
	"D4#": 311.1270,
	"E4":  329.6276,
	"F4":  349.2282,
	"F4#": 369.9944,
	"G4":  391.9954,
	"G4#": 415.3047,
	"A4":  440.0000,
	"A4#": 466.1638,
	"B4":  493.8833,
	"C5":  523.2511,
	"C5#": 554.3653,
	"D5":  587.3295,
	"D5#": 622.2540,
	"E5":  659.2551,
	"F5":  698.4565,
	"F5#": 739.9888,
	"G5":  783.9909,
	"G5#": 830.6094,
	"A5":  880.0000,
	"A5#": 932.3275,
	"B5":  987.7666,
	"C6":  1046.502,
	"C6#": 1108.731,
	"D6":  1174.659,
	"D6#": 1244.508,
	"E6":  1318.510,
	"F6":  1396.913,
	"F6#": 1479.978,
	"G6":  1567.982,
	"G6#": 1661.219,
	"A6":  1760.000,
	"A6#": 1864.655,
	"B6":  1975.533,
	"C7":  2093.005,
	"C7#": 2217.461,
	"D7":  2349.318,
	"D7#": 2489.016,
	"E7":  2637.020,
	"F7":  2793.826,
	"F7#": 2959.955,
	"G7":  3135.963,
	"G7#": 3322.438,
	"A7":  3520.000,
	"A7#": 3729.310,
	"B7":  3951.066,
	"C8":  4186.009,
}

const rate = 44100

var SampleRate = beep.SampleRate(rate)

func main() {
	// Initialize the speaker with a buffer size
	speaker.Init(SampleRate, SampleRate.N(time.Second/10))

	var wg sync.WaitGroup
	wg.Add(3)

	// Play the kick drum in a separate goroutine
	// go func() {
	// 	speaker.Play(beep.Seq(
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
	// 		beep.Callback(func() {
	// 			wg.Done()
	// 		}),
	// 	))
	// }()

	// go func() {
	// 	Vocals("Welcome")
	// 	time.Sleep(125 * time.Millisecond)
	// 	Vocals("Welcome")
	// 	time.Sleep(125 * time.Millisecond)
	// 	Vocals("Welcome")
	// 	time.Sleep(125 * time.Millisecond)
	// 	Vocals("Welcome")
	// 	time.Sleep(125 * time.Millisecond)
	// 	Vocals("Welcome")
	// 	time.Sleep(125 * time.Millisecond)
	// 	Vocals("Welcome")
	// 	time.Sleep(125 * time.Millisecond)
	// 	Vocals("Welcome")
	// 	time.Sleep(125 * time.Millisecond)
	// 	Vocals("Welcome")
	// 	time.Sleep(125 * time.Millisecond)
	// 	Vocals("Welcome")
	// 	time.Sleep(125 * time.Millisecond)
	// 	Vocals("Welcome")
	// 	Vocals("to the 1 X Developer Podcast")
	// 	wg.Done()
	// }()

	// Play the synths in the main goroutine
	go func() {
		speaker.Play(beep.Seq(
			beep.Take(SampleRate.N(7*time.Second), ChiptuneModulated(300.0, 523.25, 0.5, 1*time.Second, 7*time.Second)),
		// beep.Take(SampleRate.N(1*time.Second), Chiptune(440, 0.5)),    // A4 note
		// beep.Take(SampleRate.N(1*time.Second), Chiptune(494, 0.3)),    // B4 note (with a different pulse width)
		// beep.Take(SampleRate.N(1*time.Second), Chiptune(523.25, 0.5)), // C5 note
		// beep.Callback(func() {
		// 	wg.Done()
		// }),
		))
	}()

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

func ChiptuneModulated(startFreq, endFreq, pulseWidth float64, duration, repeatDuration time.Duration) beep.Streamer {
	phase := 0.0
	start := time.Now()
	freq := startFreq

	return beep.StreamerFunc(func(samples [][2]float64) (n int, ok bool) {
		elapsed := time.Since(start)
		stage := 0

		if elapsed < duration {
			// Stage 1: Pitch down
			stage = 1
			freq = startFreq - (float64(elapsed)/float64(duration))*(startFreq-endFreq)
		} else if elapsed < 2*duration {
			// Stage 2: Pitch up
			stage = 2
			elapsed -= duration
			freq = endFreq + (float64(elapsed)/float64(duration))*(startFreq-endFreq)
		} else {
			// Both stages completed, check if repeatDuration is over
			if elapsed >= repeatDuration {
				return 0, false
			}
			// Reset the start time and frequency
			start = time.Now()
			elapsed = 0
			freq = startFreq
		}

		_ = stage

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

func Vocals(text string) {
	cmd := exec.Command("espeak", text)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
	}
}
