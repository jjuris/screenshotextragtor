# ScreenshoteXTraGtor

![GO worklfow](https://github.com/jjuris/sxtg/actions/workflows/go.yaml/badge.svg)

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

To verify the installation, you can run:

```shell
sxtg
```

## Usage

To use the tool, run the following command:

```shell
sxtg -i example.mp4 -o screenshots -c 5 -t 10
```

This command will extract 5 screenshots from `example.mp4` at 10-second intervals and save them in the `screenshots` directory.
