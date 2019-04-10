#!/bin/bash

PLUGIN_NAME=set-weights-plugin

GOOS=linux GOARCH=amd64 go build -o ${PLUGIN_NAME}.linux64
GOOS=linux GOARCH=386 go build -o ${PLUGIN_NAME}.linux32
GOOS=windows GOARCH=amd64 go build -o ${PLUGIN_NAME}.win64
GOOS=windows GOARCH=386 go build -o ${PLUGIN_NAME}.win32
GOOS=darwin GOARCH=amd64 go build -o ${PLUGIN_NAME}.osx
