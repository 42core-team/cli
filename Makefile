VPATH	:=	cmd \
			cmd/core-cli \
			internal/bubble \
			internal/game \
			internal/github \
			internal/tournament \
			internal/utils \
			pkg \
			pkg/lib

SRCS := $(foreach dir, $(VPATH), $(wildcard $(dir)/*.go))

all:
	go run $(SRCS)

build:
	go build -o core-cli $(SRCS)
