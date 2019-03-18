# Rtc-SSH
##### SSH (Secure Shell) via WebRTC tunnel
Rtc-SSH enables connection with SSH  Raspberry PI, BeagleBone and other devices, from the browser using WebRTC. Solves the problem of the lack of public IP address, proxy server, servers behind NAT etc. You can connect to an SSH session: https://www.sqs.io

### Install
Download a binary from the releases page: https://github.com/mxseba/rtc-ssh/releases<br />
or get source using the Go compilator:
### Usage
```
$ go get -u github.com/mxseba/rtc-ssh
$ cd $GOPATH/bin
```
### First run
```
$ ./rtc-ssh newkey
uuid xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
Signal OK
```
Option <code>newkey</code> usage only first run, enter the <b>uuid</b> key on the website: https://www.sqs.io 
