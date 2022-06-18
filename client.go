package main

import (
	"fmt"
	"net"
	"time"

	"fyne.io/fyne/v2/widget"
)

var Clientstr string = ""
var netcon net.Conn

func Client(wd *widget.Label, Btn *widget.Button, needstop *bool, timeout float64, ip string, port string) {
	buf := make([]byte, 2048)
	var err error

	netcon, err = net.Dial("udp", ip+":"+port)
	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}

	for {
		netcon.SetDeadline(time.Now().Add(time.Second * time.Duration(timeout)))

		_, err = netcon.Read(buf)
		if err != nil {
			println(err.Error())
			break
		}

		serverstr += fmt.Sprintf("%s, %s", ip+":"+port, buf)
		wd.Text = serverstr
		wd.Refresh()
		println("got here")
	}

	netcon.Close()
}

func ClientSendMsg(in *widget.Entry) {

	netcon.Write([]byte(in.Text + "\n"))
}
