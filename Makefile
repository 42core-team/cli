VPATH	:=	cli \
			services

SRCS := $(foreach dir, $(VPATH), $(wildcard $(dir)/*.go))

all:
	go run $(SRCS)

build:
	go build -o core-cli $(SRCS)
