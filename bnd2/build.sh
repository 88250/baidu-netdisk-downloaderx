#!/bin/bash

export GOPROXY=https://goproxy.io
echo Building Kernel
export GOOS=windows
export GOARCH=amd64
go version
go build -v -o electron/bnd2.exe -ldflags "-s -w -H=windowsgui"

export GOOS=darwin
export GOARCH=amd64
go build -v -o electron/bnd2 -ldflags "-s -w"

echo Building UI
cd electron
node -v
npm -v
npm install && npm run dist
cd ..
