package service

import (
	"io"
	"os"
	"fmt"
	"net/url"
	"os/exec"

	"github.com/kkdai/youtube/v2"
)

func DownloadYouTube(url string) string {
	videoID, err := GetVideoID(url)
	if err != nil {
		panic(err)
	}

	client := youtube.Client{}

	video, err := client.GetVideo(videoID)
	if err != nil {
		panic(err)
	}

	formats := video.Formats.WithAudioChannels() // only get videos with audio
	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		panic(err)
	}
	defer stream.Close()

	file, err := os.Create(fmt.Sprintf("./.music/%s.mp4", video.Title))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		panic(err)
	}

	return video.Title
}

func VideoConverter(name string) {
	inputFile := fmt.Sprintf("./.music/%s.mp4", name)
	outputFile := fmt.Sprintf("./.music/%s.mp3", name)

	os.Remove(outputFile)
	defer os.Remove(inputFile)

	cmd := exec.Command("ffmpeg", "-i", inputFile, "-vn", "-acodec", "libmp3lame", outputFile)
	cmd.Dir = "/Users/levakondratev/dev/bg-bot"

	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func GetVideoID(link string) (string, error) {
	parsedURL, err := url.Parse(link)
	if err != nil {
		return "", err
	}

	queryParams := parsedURL.Query()
	videoID := queryParams.Get("v")

	if videoID == "" {
		segments := parsedURL.Path
		if len(segments) > 1 && segments[0] == '/' {
			segments = segments[1:]
		}
		segments = segments[:len(segments)-1] // Remove the trailing slash
		videoID = segments
	}

	if videoID == "" {
		return "", fmt.Errorf("video ID not found in the URL")
	}

	return videoID, nil
}
