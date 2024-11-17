# ScreenshoteXTraGtor

A screenshot extractor tool written in GO lang. The script automates the process of extracting multiple screenshots from a video file at specified intervals using [ffmpeg](https://www.ffmpeg.org).

## Installation

To install the tool, you need to have Go installed on your system. Additionally, the tool depends on `ffmpeg`, which must be installed separately.

### Install Go

Follow the instructions on the [official Go website](https://golang.org/doc/install) to install Go on your system.

Or simply use Homebrew on macOS:

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
go install github.com/jjuris/sxtg
```

To verify the installation, you can run the tests:

```shell
sxtg
```

## Usage

To use the tool, run the following command:

```shell
sxtg -i example.mp4 -o screenshots -c 5 -t 10
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
