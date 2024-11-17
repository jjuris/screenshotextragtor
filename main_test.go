package main

import (
	"flag"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

var debug = flag.Bool("d", false, "Enable debug messages")

func TestExtractScreenshot(t *testing.T) {
	// Create a temporary directory for test output
	outputDir := t.TempDir()
	outputFile := outputDir + "/test_screenshot.png"

	// Mock video file (assuming a small test video file is available)
	videoFile := "test_video.mp4"

	// Test case: valid screenshot extraction
	err := extractScreenshot(videoFile, outputFile, 1, *debug)
	if err != nil {
		t.Errorf("Failed to extract screenshot: %v", err)
	}

	// Check if the output file exists
	if _, err := os.Stat(outputFile); os.IsNotExist(err) {
		t.Errorf("Screenshot file was not created")
	}
}

func TestMainFunction(t *testing.T) {
	// Mock video file (assuming a small test video file is available)
	videoFile := "test_video.mp4"

	// Create a temporary output directory
	outputDir, err := os.MkdirTemp("", "output")
	if err != nil {
		t.Fatalf("Failed to create temp output directory: %v", err)
	}
	defer os.RemoveAll(outputDir)

	// Mock command-line arguments
	os.Args = []string{"cmd", "-i", videoFile, "-o", outputDir, "-c", "1", "-t", "1", "-d"}

	// Reset command-line flags
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	// Run the main function
	main()

	// Check if the screenshot was created
	outputFile := filepath.Join(outputDir, filepath.Base(videoFile)+"_1.png")
	if _, err := os.Stat(outputFile); os.IsNotExist(err) {
		t.Fatalf("Expected screenshot file %s does not exist", outputFile)
	}
}

func TestMainFunction_EmptyInput(t *testing.T) {
	// Mock command-line arguments with empty input file
	os.Args = []string{"cmd", "-i", "", "-o", outputDir, "-c", "1", "-t", "1"}

	// Reset command-line flags
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	// Capture the exit code
	exitCode := 0
	exitFunc := func(code int) {
		exitCode = code
	}
	defer func() { os.Exit = exitFunc }()

	// Run the main function
	main()

	// Check if the exit code is non-zero
	if exitCode == 0 {
		t.Fatalf("Expected non-zero exit code for empty input file")
	}
}

func TestMain(m *testing.M) {
	flag.Parse()

	// Setup code: create a small test video file using ffmpeg
	cmd := exec.Command("ffmpeg", "-f", "lavfi", "-i", "testsrc=duration=5:size=1280x720:rate=30", "test_video.mp4")
	err := cmd.Run()
	if err != nil {
		panic("Failed to create test video file")
	}

	// Run tests
	code := m.Run()

	// Teardown code: remove the test video file
	os.Remove("test_video.mp4")

	os.Exit(code)
}
