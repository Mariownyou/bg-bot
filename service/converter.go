package service

import (
	"io"
	"os"
	"os/exec"

	"github.com/jonas747/dca"
)

func VideoConverter() {
	inputFile := "./service/video.mp4"
	outputFile := "./service/output.mp3"

	os.Remove(outputFile)

	cmd := exec.Command("ffmpeg", "-i", inputFile, "-vn", "-acodec", "libmp3lame", outputFile)
	// ffmpeg -i ./service/video.mp4 -vn -acodec libmp3lame ./service/output.mp3
	// set cwd to ./service
	cmd.Dir = "/Users/levakondratev/dev/bg-bot"

	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func AudioConverter() {
	base := "/Users/levakondratev/dev/bg-bot/"
	encodeSession, err := dca.EncodeFile("music.mp3", dca.StdEncodeOptions)
	if err != nil {
		panic(err)
	}
	// Make sure everything is cleaned up, that for example the encoding process if any issues happened isnt lingering around
	defer encodeSession.Cleanup()

	output, err := os.Create(base + "cmd/bot/output.dca")
	if err != nil {
		panic(err)
		// Handle the error
	}

	io.Copy(output, encodeSession)
}
