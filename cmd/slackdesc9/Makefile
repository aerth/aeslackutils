export GOFLAGS ?= -ldflags=-w -ldflags=-s -tags=osusergo,netgo -v
PREFIX ?= /usr/local
slackdesc9: *.go
	go build -o $@
install: slackdesc9
	install $< ${PREFIX}/bin/
clean:
	rm -f slackdesc9
