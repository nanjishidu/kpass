#!/bin/sh
#go get -u github.com/jteeuwen/go-bindata/...
#go generate
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/linux_amd64/kpass *.go 
CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o bin/linux_386/kpass *.go 
CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o bin/linux_arm/kpass *.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/windows_amd64/kpass *.go 
CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o bin/windows_386/kpass *.go 
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/darwin/kpass *.go
cd bin
tar -zcvf linux_amd64.tar.gz linux_amd64
tar -zcvf linux_386.tar.gz linux_386
tar -zcvf linux_arm.tar.gz linux_arm
tar -zcvf darwin.tar.gz darwin
zip -r windows_amd64.zip windows_amd64
zip -r windows_386.zip windows_386
rm -rf linux_amd64 linux_386 linux_arm darwin windows_amd64 windows_386