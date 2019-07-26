#!/bin/bash

export CGO_ENABLED=1

go version
go build -o bnd -ldflags "-s -w"
