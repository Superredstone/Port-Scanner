#!/bin/bash
Name=PortScanner

cd ..
rsrc.exe -arch amd64 -ico assets/icon.ico
rsrc.exe -arch arm -ico assets/icon.ico

env GOOS=windows GOARCH=amd64 go build -o $Name-windows-amd64.exe -ldflags="-H windowsgui"
echo Windows amd64 builded

mv $Name-windows-amd64.exe build/

cd build/
env GOOS=windows GOARCH=arm go build -o $Name-windows-arm.exe -ldflags="-H windowsgui" ../main.go
echo Windows arm builded

env GOOS=linux GOARCH=amd64 go build -o $Name-linux-amd64 ../main.go
echo Linux amd64 builded

env GOOS=android GOARCH=arm go build .
echo Android builded