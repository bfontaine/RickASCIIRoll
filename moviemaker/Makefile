URL=https://www.youtube.com/watch?v=dQw4w9WgXcQ
# Change this based on your storage space
#  8 = 1698 frames, 4.4GB, sleep ~0.15
# 10 = 2122 frames, 5.5GB, sleep ???
RATE=10
SLEEP=0.10

video.mp4:
	youtube-dl -f mp4 -o $@ $(URL)

frames: video.mp4
	mkdir -p frames
	ffmpeg -i video.mp4 -r $(RATE) 'frames/frame%05d.bmp'

# note some parameters below won't work with another video
txtframes: frames
	mkdir -p txtframes
	for f in frames/frame*.bmp; do \
		echo "==> $${f%.bmp}"; \
		img2txt.py --maxLen=86 --targetAspect=0.5 $$f | \
			python ./process_frame.py --lpad 12 --rpad 11 > txt$${f%bmp}txt; \
	done

play: txtframes
	for f in txtframes/frame*; do clear; cat $$f; sleep $(SLEEP); done
