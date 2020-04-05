CFLAGS = -Wall -g -pedantic
LDFLAGS = $(shell pkg-config --cflags --libs x11)

.PHONY: all
all: bouncy

bouncy: bouncy.c
	cc $(CFLAGS) $^ $(LDFLAGS) -o $@

.PHONY: clean
clean:
	rm -f bouncy
