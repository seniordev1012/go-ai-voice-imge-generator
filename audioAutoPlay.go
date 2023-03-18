package main

import (
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"os"
	"time"
)

func playAudio(filename string) error {
	// Open the audio file
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	// Decode the audio file
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		return err
	}
	defer streamer.Close()

	// Initialize the audio player
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	// Play the audio stream
	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	// Wait for the audio to finish playing
	<-done

	fmt.Println("Audio playback complete.")
	return nil
}
