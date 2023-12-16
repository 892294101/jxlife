GOCMD=go
export TAG=1.2.0
export BD=$(shell date '+%b %d %Y %T')
GOBUILD=${GOCMD} build -gcflags=all='-l -N' -ldflags '-s -w'

BUILD_DIR=../build/server
BINARY_DIR=$(BUILD_DIR)/bin

MYSQL_EXTRACT_FILE=$(BINARY_DIR)/jxserver
MYSQL_EXTRACT_SRC=./main.go


.PHONY: all clean build

all: clean build

clean:

build:
	${GOBUILD} -o ${MYSQL_EXTRACT_FILE} ${MYSQL_EXTRACT_SRC}

