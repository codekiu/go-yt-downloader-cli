package main

import (
	"flag"
	"fmt"
)

type Config struct {
	ytUrl  string
	format string
}

var supportedFormats = []string{"mp3", "mp4", "wav"}

func main() {
	fmt.Println("Verifying input...")
	config, err := parseInput()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Downloading YouTube video: %s and formatting it as: %s\n", config.ytUrl, config.format)
}

func checkStringInSlice(slice []string, target string) bool {
	for _, str := range slice {
		if str == target {
			return true
		}
	}
	return false
}

func parseInput() (Config, error) {
	ytUrl := flag.String("u", "", "Enter YouTube URL to download and parse")
	format := flag.String("f", "mp4", "Enter format")
	flag.Parse()

	if err := validation(*ytUrl, *format); err != nil {
		return Config{}, err
	}

	return Config{
		ytUrl:  *ytUrl,
		format: *format,
	}, nil
}

func validation(url string, format string) error {
	if url == "" || format == "" {
		return fmt.Errorf("all flags must be set")
	}
	if !checkStringInSlice(supportedFormats, format) {
		return fmt.Errorf("unsupported format: %s", format)
	}
	return nil
}
