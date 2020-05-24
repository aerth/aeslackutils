export GOFLAGS ?= -ldflags=-w -ldflags=-s -tags=osusergo,netgo -v
slackdesc9: *.go
	go build -o $@
clean:
	rm -f slackdesc9
