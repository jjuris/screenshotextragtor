# ScreenshoteXTraGtor

A screenshot extractor tool written in GO lang.

## Installation

To install the tool, you need to have Go installed on your system. Additionally, the tool depends on `ffmpeg`, which must be installed separately.

### Install Go

Follow the instructions on the [official Go website](https://golang.org/doc/install) to install Go on your system.

Or use Homebrew on macOS:

```shell
brew install go
```

### Install ffmpeg

You can install `ffmpeg` using Homebrew on macOS:

```shell
brew install ffmpeg
```

### Install ScreenshoteXTraGtor

Once you have Go and `ffmpeg` installed, you can install ScreenshoteXTraGtor using the following commands:

```shell
go get github.com/jjuris/screenshotextragtor
cd $GOPATH/src/github.com/jjuris/screenshotextragtor
go install
```

To verify the installation, you can run the tests:

```shell
go test ./...
```

## Usage

To use the tool, run the following command:

```shell
./sxtg <video_file> [output_directory] [-c count] [-i interval]
```

* `<video_file>`: The path to the video file from which to extract screenshots.
* `[output_directory]`: (Optional) The directory where the screenshots will be saved. Defaults to the current directory.
* `-c count`: (Optional) The number of screenshots to extract. Defaults to 1.
* `-i interval`: (Optional) The time interval between screenshots in seconds. Defaults to 1.

Example:

```shell
./sxtg example.mp4 screenshots -c 5 -i 10
```

This command will extract 5 screenshots from `example.mp4` at 10-second intervals and save them in the `screenshots` directory.

## Development

To contribute to the development of this tool, follow these steps:

* Fork the repository on GitHub.
* Clone your forked repository to your local machine.
* Create a new branch for your feature or bugfix.
* Make your changes and commit them with descriptive messages.
* Push your changes to your forked repository.
* Create a pull request on the original repository.
* Make sure to run the tests before submitting your pull request:

Feel free to open issues for any bugs or feature requests. ```
