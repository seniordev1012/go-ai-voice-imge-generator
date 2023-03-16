package main

//
//func PlayNotificationSound() {
//	// Open the MP3 file
//	f, err := os.Open("notification.mp3")
//	if err != nil {
//		fmt.Println("Failed to open notification sound file:", err)
//		return
//	}
//	defer f.Close()
//
//	// Decode the MP3 file
//	streamer, format, err := mp3.Decode(f)
//	if err != nil {
//		fmt.Println("Failed to decode notification sound file:", err)
//		return
//	}
//	defer streamer.Close()
//
//	// Initialize the speaker
//	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
//	if err != nil {
//		fmt.Println("Failed to initialize speaker:", err)
//		return
//	}
//
//	// Play the notification sound
//	done := make(chan bool)
//	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
//		done <- true
//	})))
//
//	// Wait for the sound to finish playing
//	<-done
//}
