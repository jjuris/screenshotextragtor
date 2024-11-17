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
	inputFile := flag.String("i", "", "Input video file (required)")
	outputDir := flag.String("o", ".", "Output directory")
	count := flag.Int("c", 1, "Count of screenshots")
	interval := flag.Int("t", 1, "Time interval between screenshots in seconds")
	debug := flag.Bool("d", false, "Enable debug messages")
	flag.Usage = func() {
		fmt.Println("Usage: sxtg -i <video_file> [-o output_directory] [-c count] [-t interval] [-d]")
		flag.PrintDefaults()
	}
	flag.Parse()

	// Check for required arguments
	if *inputFile == "" {
		flag.Usage()
		os.Exit(1)
	}

	// Ensure the output directory exists
	if _, err := os.Stat(*outputDir); os.IsNotExist(err) {
		fmt.Printf("Output directory %s does not exist\n", *outputDir)
		os.Exit(1)
	}

	// Print the parsed arguments (for debugging purposes)
	if *debug {
		fmt.Printf("Video File: %s\n", *inputFile)
		fmt.Printf("Output Directory: %s\n", *outputDir)
		fmt.Printf("Count: %d\n", *count)
		fmt.Printf("Interval: %d\n", *interval)
	}

	// Extract screenshots
	for i := 0; i < *count; i++ {
		timestamp := i * *interval
		outputFile := filepath.Join(*outputDir, fmt.Sprintf("%s_%d.png", filepath.Base(*inputFile), i+1))
		err := extractScreenshot(*inputFile, outputFile, timestamp, *debug)
		if err != nil {
			fmt.Printf("Error extracting screenshot at %d seconds: %v\n", timestamp, err)
		}
	}
}

func extractScreenshot(videoFile, outputFile string, timestamp int, debug bool) error {
	if debug {
		fmt.Printf("Extracting screenshot at %d seconds\n", timestamp)
	}
	cmd := exec.Command("ffmpeg", "-y", "-ss", strconv.Itoa(timestamp), "-i", videoFile, "-vframes", "1", "-update", "1", outputFile)
	if debug {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	return cmd.Run()
}
