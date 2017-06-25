# kpass
password management tool for golang

## Download

### kpass 1.0

[windows_386.zip](http://or5mbu6lx.bkt.clouddn.com/1.0/windows_386.zip)
|[windows_amd64.zip](http://or5mbu6lx.bkt.clouddn.com/1.0/windows_amd64.zip)
|[linux_arm.tar.gz](http://or5mbu6lx.bkt.clouddn.com/1.0/linux_arm.tar.gz)
|[linux_386.tar.gz](http://or5mbu6lx.bkt.clouddn.com/1.0/linux_386.tar.gz)
|[linux_amd64.tar.gz](http://or5mbu6lx.bkt.clouddn.com/1.0/linux_amd64.tar.gz)
|[darwin.tar.gz](http://or5mbu6lx.bkt.clouddn.com/1.0/darwin.tar.gz)

## Install kpass
	
	cd $GOPATH/src
	git clone https://github.com/nanjishidu/kpass.git
	cd  kpass
	go build -o bin/kpass *.go

## Cross Compiling

    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/linux_amd64/kpass *.go 
    CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o bin/linux_386/kpass *.go 
    CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o bin/linux_arm/kpass *.go
    CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/windows_amd64/kpass *.go 
    CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o bin/windows_386/kpass *.go 
    CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/darwin/kpass *.go


## Use kpass

```    
./kpass create -kp=nanjishidu -kf=account.db -kc=aes-cfb
./kpass run 
```

## kpass
```
NAME:
   kpass - password management tool for golang

USAGE:
   kpass [global options] command [command options] [arguments...]

VERSION:
   1.0

COMMANDS:
     create, c  create kpassfile by kpassword
     run, r     kpass run
     help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```
## kpass create

kpass create command options
```
NAME:
   kpass create - create kpassfile by kpassword

USAGE:
   kpass create [command options] [arguments...]

OPTIONS:
   --kpassdir value, --kd value   kpass dir (default: "./data")
   --kpassfile value, --kf value  kpass encrypt file (default: "./kpass-encrypt")
   --kpassword value, --kp value  kpass password
   --kcrypto value, --kc value    kpass crypto (default: "aes-cfb")
```
## kpass run

kpass run command options
```
NAME:
   kpass run - kpass run

USAGE:
   kpass run [command options] [arguments...]

OPTIONS:
   --kpassdir value, --kd value  kpass dir (default: "./data")
   --host value                  kpass web host (default: "127.0.0.1")
   --port value                  kpass web port (default: 9988)   
```


