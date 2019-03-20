package main

import (
	"github.com/pions/webrtc"
)
	

var (
	configRTC = webrtc.RTCConfiguration{
		IceServers: []webrtc.RTCIceServer{
			{
				URLs: []string{
					"stun:stun.l.google.com:19302",
				},
			},
		},
	}
	defaultHost = "127.0.0.1"
	defaultPort = 22
)
