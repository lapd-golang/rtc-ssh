package main

import (
	"fmt"
	"log"
	"net"
	"io"
	"github.com/pion/webrtc/v2"
	"github.com/gorilla/websocket"
)

type Wrap struct {
	*webrtc.DataChannel
}

func (rtc *Wrap) Write(data []byte) (int, error) {
	err := rtc.DataChannel.Send(data)
	return len(data), err
}

func interpreter(ws *websocket.Conn, data Json, conf Config) error {
	if data.Error != "" {
		return fmt.Errorf(data.Error)
	}
	
	switch data.Type {	
		case "signal_OK":
			log.Println("Signal OK")
		
		case "offer":
			pc, err := webrtc.NewPeerConnection(configRTC)
			if err != nil {
				return err
			}
									
			pc.OnICEConnectionStateChange(func(state webrtc.ICEConnectionState) {
				log.Println("ICE Connection State has changed:", state.String())
			})
			 					
			pc.OnDataChannel(func(dc *webrtc.DataChannel) {
				if dc.Label() == "SSH" {
					ssh, err := net.Dial("tcp", fmt.Sprintf("%s:%d", conf.Host, conf.Port))
					if err != nil {
						log.Println("ssh dial failed:", err)
						pc.Close()
					} else {
						log.Println("Connect SSH socket")
						DataChannel(dc, ssh)
					}
				}
			})
						
			if err := pc.SetRemoteDescription(webrtc.SessionDescription{
				Type: webrtc.SDPTypeOffer,
				SDP:  data.Sdp,
			}); err != nil {
				pc.Close()
				return err
			}
					
			answer, err := pc.CreateAnswer(nil)
			if err != nil {
				pc.Close()
				return err
			}
		
			err = pc.SetLocalDescription(answer)
			if err != nil {
				pc.Close()
				return err
			}
		
			if err = ws.WriteJSON(answer); err != nil {
				return err
			}
		
			
		default:
			return fmt.Errorf("unknown signaling message %v", data.Type)
	}
	return nil
}


func DataChannel(dc *webrtc.DataChannel, ssh net.Conn) {
	dc.OnOpen(func() {	
		err := dc.SendText("OPEN_RTC_CHANNEL")
		if err != nil{
			log.Println("write data error:", err)
		}
		io.Copy(&Wrap{dc}, ssh)
	})
	dc.OnMessage(func(msg webrtc.DataChannelMessage) {
		ssh.Write(msg.Data)
	})
	dc.OnClose(func() {
		log.Printf("Close SSH socket")
		ssh.Close()
	})
}
