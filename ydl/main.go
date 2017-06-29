package main

import (
	"log"

	"github.com/BrianAllred/goydl/ydl"
)

func main() {
	ydl := ydl.NewYoutubeDl()
	ydl.Options.Output.Value = "/home/brian/Downloads/test.mp3"
	ydl.Options.AudioFormat.Value = "mp3"
	ydl.Options.ExtractAudio.Value = true

	cmd, err := ydl.Download("")

	if err != nil {
		log.Fatal(err)
	}

	cmd.Wait()
}
