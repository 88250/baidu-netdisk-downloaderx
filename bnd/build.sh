#!/bin/bash

export CGO_ENABLED=1

go build -o bnd -ldflags "-s -w"