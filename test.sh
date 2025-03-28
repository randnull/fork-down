#!/bin/bash

go build -o fork-down main.go

echo testing manifest similar to .bin

./fork-down -file "test15.bin" -manifest "manifest.json"

echo testing manifest not similar to .bin

./fork-down -file "test15_modified.bin" -manifest "manifest.json"
