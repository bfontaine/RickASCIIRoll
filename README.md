# RickASCIIRoll

## Setup

You need Python, `youtube-dl`, `ffmpeg`, and [`img2txt.py`][img2txt].

On macOS:

    brew install youtube-dl ffmpeg
    pip install img2txt.py

[img2txt]: https://github.com/hit9/img2txt

Then run `make txtframes` to get all video frames as text. Be aware that it
creates one 2MB `bmp` file for each frame in the video so you’ll need some
room.
Run `make play` for a preview.
