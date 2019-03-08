# Rtc-SSH
SSH (Secure Shell) via WebRTC tunnel\n
Rtc-SSH enables connection with SSH  Raspberry PI, BeagleBone and other devices, from the browser using WebRTC. Solves the problem of the lack of public IP address, proxy server, servers behind NAT etc. You can connect to an SSH session: https://www.sqs.io
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
Copy uuid key and paste on the static page: https://www.sqs.io 
