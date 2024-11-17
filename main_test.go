package main

import (
	"flag"
	"os"
	"os/exec"
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
