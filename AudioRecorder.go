package main

/*
  #include <stdio.h>
  #include <unistd.h>
  #include <termios.h>
  char getch(){
      char ch = 0;
      struct termios old = {0};
      fflush(stdout);
      if( tcgetattr(0, &old) < 0 ) perror("tcsetattr()");
      old.c_lflag &= ~ICANON;
      old.c_lflag &= ~ECHO;
      old.c_cc[VMIN] = 1;
      old.c_cc[VTIME] = 0;
      if( tcsetattr(0, TCSANOW, &old) < 0 ) perror("tcsetattr ICANON");
      if( read(0, &ch,1) < 0 ) perror("read()");
      old.c_lflag |= ICANON;
      old.c_lflag |= ECHO;
      if(tcsetattr(0, TCSADRAIN, &old) < 0) perror("tcsetattr ~ICANON");
      return ch;
  }
*/
import "C"
import (
	"fmt"
	"github.com/gordonklaus/portaudio"
	wave "github.com/zenwerk/go-wave"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

func errCheck(err error) {

	if err != nil {
		panic(err)
	}
}

func VoiceRecorder() (string, error) {
	var wg sync.WaitGroup
	wg.Add(2)
	audioFileName := "cache/audio.wav"

	fmt.Println("Recording. Press ESC to quit.")

	if !strings.HasSuffix(audioFileName, ".wav") {
		audioFileName += ".wav"
	}
	waveFile, err := os.Create(audioFileName)
	errCheck(err)

	inputChannels := 1
	outputChannels := 0
	sampleRate := 44100
	framesPerBuffer := make([]byte, 64)

	// init PortAudio

	portaudio.Initialize()
	//defer portaudio.Terminate()

	stream, err := portaudio.OpenDefaultStream(inputChannels, outputChannels, float64(sampleRate), len(framesPerBuffer), framesPerBuffer)
	errCheck(err)
	//defer stream.Close()

	// setup Wave file writer

	param := wave.WriterParam{
		Out:           waveFile,
		Channel:       inputChannels,
		SampleRate:    sampleRate,
		BitsPerSample: 8, // if 16, change to WriteSample16()
	}

	waveWriter, err := wave.NewWriter(param)
	errCheck(err)

	//defer waveWriter.Close()

	go func() {
		_, err := func() (string, error) {

			durationVoice := 5
			//use countdown timer to record for 15 seconds
			for i := durationVoice; i >= 0; i-- {
				time.Sleep(1 * time.Second)
				if i == 0 {

					//At the end of the countdown, stop recording
					//This is where you would do whatever you want to do after the countdown
					log.Printf("Time's up!")
					waveWriter.Close()
					stream.Close()
					portaudio.Terminate()
					fmt.Println("Play", audioFileName, "with a audio player to hear the result.")
					//Whisper(audioFileName)
					//TODO: Below is the code to "Whisper" the audio file ... risky stuff
					resultCh := make(chan string)

					// Start the Go routine with the channel as an argument
					go Whisper(audioFileName, resultCh)
					//go Whisper(audioFileName)
					//TODO: This needs to be returned to the calling function to be used in the "Whisper" function
					log.Println("Result: ", <-resultCh)
					log.Println("Done")
					wg.Wait()
					wg.Done()
					return audioFileName, nil

				}
			}
			return "", nil

		}()
		if err != nil {
		}
	}()
	//
	// start reading from microphone
	errCheck(stream.Start())
	for {
		errCheck(stream.Read())
		//fmt.Print(".") // show progress

		fmt.Print(".") // show progress
		// write to wave file
		_, err := waveWriter.Write([]byte(framesPerBuffer)) // WriteSample16 for 16 bits
		errCheck(err)
	}

	errCheck(stream.Stop())

	return "", nil
}
