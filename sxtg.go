package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

func main() {
	// Define command-line arguments
	count := flag.Int("c", 1, "Count of screenshots")
	interval := flag.Int("i", 1, "Time interval between screenshots in seconds")
	flag.Parse()

	// Check for required positional arguments
	if len(flag.Args()) < 1 {
		fmt.Println("Usage: sxtg <video_file> [output_directory] [-c count] [-i interval]")
		os.Exit(1)
	}

	videoFile := flag.Arg(0)
	outputDir := "."
	if len(flag.Args()) > 1 {
		outputDir = flag.Arg(1)
	}

	// Ensure the output directory exists
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		fmt.Printf("Output directory %s does not exist\n", outputDir)
		os.Exit(1)
	}

	// Print the parsed arguments (for debugging purposes)
	fmt.Printf("Video File: %s\n", videoFile)
	fmt.Printf("Output Directory: %s\n", outputDir)
	fmt.Printf("Count: %d\n", *count)
	fmt.Printf("Interval: %d\n", *interval)

	// Extract screenshots
	for i := 0; i < *count; i++ {
		timestamp := i * *interval
		outputFile := filepath.Join(outputDir, fmt.Sprintf("%s_%d.png", filepath.Base(videoFile), i+1))
		err := extractScreenshot(videoFile, outputFile, timestamp)
		if err != nil {
			fmt.Printf("Error extracting screenshot at %d seconds: %v\n", timestamp, err)
		}
	}
}

func extractScreenshot(videoFile, outputFile string, timestamp int) error {
	fmt.Printf("Timestamp: %s\n", strconv.Itoa(timestamp))
	cmd := exec.Command("ffmpeg", "-ss", strconv.Itoa(timestamp), "-i", videoFile, "-vframes", "1", outputFile)
	return cmd.Run()
}
