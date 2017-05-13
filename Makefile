URL=https://www.youtube.com/watch?v=dQw4w9WgXcQ

video.mp4:
	youtube-dl -f mp4 -o $@ $(URL)

frames: video.mp4
	mkdir -p frames
	ffmpeg -i video.mp4  'frames/frame%%03d.bmp'

txtframes: frames
	mkdir -p txtframes
	for f in frames/frame*.bmp; do \
		echo "==> $$f"; \
		img2txt.py --maxLen=80 --targetAspect=0.5 $$f | \
			ruby -e '$$stdin.read=~/<pre>(.+)<\/pre>/m;puts $$1' \
				> txt$${f%%bmp}txt; \
	done

play: txtframes
	for f in txtframes/frame*; do clear; cat $$f; sleep 0.02; done
