#!/bin/bash

mkdir -p ../bin
mkdir -p ../bin/win64
mkdir -p ../bin/linux64
mkdir -p ../bin/darwin64
mkdir -p ../bin/linuxARMv7

export GOOS=windows
export GOARCH=amd64
go build -o ../bin/win64/weybot.exe ../../main.go
cp ../../config.json ../bin/win64/config.json
echo "Built bot for Windows/amd64"

export GOOS=linux
export GOARCH=amd64
go build -o ../bin/linux64/weybot ../../main.go
cp ../../config.json ../bin/linux64/config.json
echo "Built bot for Linux/amd64"

export GOOS=darwin
export GOARCH=amd64
go build -o ../bin/darwin64/weybot ../../main.go
cp ../../config.json ../bin/darwin64/config.json
echo "Built bot for Darwin/amd64"

export GOOS=linux
export GOARCH=arm
export GOARM=7
go build -o ../bin/linuxARMv7/weybot ../../main.go
cp ../../config.json ../bin/linuxARMv7/config.json
echo "Built bot for Linux/ARMv7"