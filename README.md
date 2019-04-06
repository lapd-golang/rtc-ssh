# Rtc-SSH
##### SSH (Secure Shell) via WebRTC
Rtc-SSH enables connection with SSH  Raspberry PI, BeagleBone and other devices, from the browser using WebRTC. Solves the problem of the lack of public IP address, proxy server, servers behind NAT etc. You can connect to an SSH session: https://www.sqs.io
<!--
### Install from binary
```
wget https://github.com/mxseba/rtc-ssh/releases/download/v0.1.1/rtc-ssh_0.1.1_Linux_armv7.tar.gz
tar xvfz rtc-ssh_0.1.1_Linux_armv7.tar.gz
cd rtc-ssh_0.1.1_Linux_armv7
./rtc-ssh newkey
uuid xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
Signal OK
```
Other architectures and releases: https://github.com/mxseba/rtc-ssh/releases<br />

or get source using the Go compilator:
-->
### Usage
```
go get -u github.com/mxseba/rtc-ssh
cd $GOPATH/bin
```
### First run
```
./rtc-ssh newkey
uuid xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
Signal OK
```
Option <code>newkey</code> usage only first run, enter the <b>uuid</b> key on the website: https://www.sqs.io 

Rtc-SSH uses the pion-WebRTC library: https://github.com/pions/webrtc
