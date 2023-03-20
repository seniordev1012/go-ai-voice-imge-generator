package aigenAudioAutoPlay

import (
	"fmt"
	"github.com/hajimehoshi/go-mp3"
	"github.com/youpy/go-wav"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gen2brain/malgo"
)

func PlayAudioPlayback(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("failed to close file: %v", err)
		}
	}(file)

	var reader io.Reader
	var channels, sampleRate uint32
	var audioLength int64

	switch strings.ToLower(filepath.Ext(filename)) {
	case ".wav":
		w := wav.NewReader(file)
		f, err := w.Format()
		if err != nil {
			return fmt.Errorf("failed to read WAV format: %w", err)
		}

		reader = w
		channels = uint32(f.NumChannels)
		sampleRate = f.SampleRate

	case ".mp3":
		m, err := mp3.NewDecoder(file)
		if err != nil {
			return fmt.Errorf("failed to decode MP3: %w", err)
		}

		reader = m
		channels = 2
		audioLength = m.Length()
		sampleRate = uint32(m.SampleRate())
	default:
		return fmt.Errorf("unsupported audio format")
	}

	ctx, err := malgo.InitContext(nil, malgo.ContextConfig{}, func(message string) {
		//fmt.Printf("LOG <%v>\n", message)
	})
	if err != nil {
		return fmt.Errorf("failed to init context: %w", err)
	}
	defer func() {
		_ = ctx.Uninit()
		ctx.Free()
		return
	}()

	deviceConfig := malgo.DefaultDeviceConfig(malgo.Playback)
	deviceConfig.Playback.Format = malgo.FormatS16
	deviceConfig.Playback.Channels = channels
	deviceConfig.SampleRate = sampleRate
	deviceConfig.Alsa.NoMMap = 1

	// This is the function that's used for sending more data to the device for playback.
	onSamples := func(pOutputSample, pInputSamples []byte, framecount uint32) {
		io.ReadFull(reader, pOutputSample)
	}

	deviceCallbacks := malgo.DeviceCallbacks{
		Data: onSamples,
	}
	device, err := malgo.InitDevice(ctx.Context, deviceConfig, deviceCallbacks)
	if err != nil {
		return fmt.Errorf("failed to init device: %w", err)
	}
	defer device.Uninit()

	if err = device.Start(); err != nil {
		return fmt.Errorf("failed to start device: %w", err)
	}

	durationMillis := audioLength / 50 // Set loop duration to 10 seconds (in milliseconds)
	for i := durationMillis; i >= 0; i-- {
		time.Sleep(1 * time.Millisecond)
		if i == 0 {

			break // Exit the loop if duration has elapsed
		}

		// Do something here (e.g. print a message)
		fmt.Println(i/1000, "seconds remaining")

	}
	return nil
}
