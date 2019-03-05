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
)