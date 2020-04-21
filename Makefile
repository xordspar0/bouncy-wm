CFLAGS = -Wall -g -pedantic
LDFLAGS = $(shell pkg-config --cflags --libs x11)

.PHONY: all
all: bouncy

.PHONY:
bouncy:
	go build

.PHONY: run
run: bouncy
	Xephyr -ac -screen 1820x980 -br -reset -terminate :2 &
	sleep 0.1
	xterm -display :2 &
	env DISPLAY=:2 ./bouncy-wm

.PHONY: clean
clean:
	rm -f bouncy-wm
