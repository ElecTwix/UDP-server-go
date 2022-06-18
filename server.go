package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"fyne.io/fyne/v2/widget"
)

var serverstr string = ""
var lastconnect net.Addr

var globalnet net.PacketConn

func sendResponse(conn *net.UDPConn, addr *net.UDPAddr) {
	_, err := conn.WriteToUDP([]byte("From server: Hello I got your message "), addr)
	if err != nil {
		fmt.Printf("Couldn't send response %v", err)
	}
}

func Server(wd *widget.Label, Btn *widget.Button, needstop *bool, timeout float64) {

	var err error
	if globalnet == nil {
		println("started listen")

		// listen to incoming udp packets
		globalnet, err = net.ListenPacket("udp", ":7777")
		if err != nil {
			log.Fatal(err)
		}
	}

	for {

		if *needstop {

			break
		}

		buf := make([]byte, 1024)

		globalnet.SetDeadline(time.Now().Add(time.Second * time.Duration(timeout)))
		_, lastconnect, err = globalnet.ReadFrom(buf)
		if err != nil {
			println(err.Error())
			break
		}

		serverstr += fmt.Sprintf("%s, %s", lastconnect, buf)
		wd.Text = serverstr
		wd.Refresh()
		println("got here")

	}
	println("finish")
	Btn.Enable()

}

func sendmsg(in *widget.Entry) {

	globalnet.WriteTo([]byte(in.Text+"\n"), lastconnect)
}
