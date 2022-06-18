package main

import (
	"fmt"
	"os"
	"os/user"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var globalapp fyne.App
var mainpage fyne.Window

func main() {

	globalapp = app.New()
	mainpage = globalapp.NewWindow("UDP Client & Server")

	mainmenu()
	mainpage.ShowAndRun()
}

func mainmenu() {
	Client := widget.NewButton("Client", Clientgui)

	currentUser, err := user.Current()

	username := currentUser.Username

	hname, err := os.Hostname()

	if err != nil {
		panic(err)
	}

	ServerBtn := widget.NewButton("Server", servergui)

	hello := widget.NewLabel(fmt.Sprintf("Hello, %s (%s)", username, hname))

	mainpage.SetContent(container.NewVBox(hello, container.NewHBox(
		Client,
		ServerBtn,
		widget.NewButton("Github", func() { openbrowser("https://github.com/ElecTwix") }),
	)))

}

func servergui() {

	needsstop := false

	hello := widget.NewLabel("Console Log")
	//timeout := widget.NewEntry()
	timeout := widget.NewSlider(1, 300)
	timeoutlabel := widget.NewLabel("1")
	timeout.Value = 60
	timeout.Refresh()
	timeout.OnChanged = func(f float64) {

		timeoutlabel.Text = fmt.Sprintf("Time out %f", f)
		timeoutlabel.Refresh()
	}

	timeoutlabel.Text = fmt.Sprintf("Time out %f", timeout.Value)
	timeoutlabel.Refresh()

	msg := widget.NewEntry()
	msg.PlaceHolder = "Packet Message"

	BackBtn := widget.NewButton("Back", mainmenu)

	SndMsg := widget.NewButton("Send", func() { sendmsg(msg) })
	SndMsg.Disable()

	var ServerBtn *widget.Button
	ServerBtn = widget.NewButton("Server", func() {

		go Server(hello, ServerBtn, &needsstop, timeout.Value)
		SndMsg.Enable()
		ServerBtn.Disable()

	})

	Github := widget.NewButton("Github", func() { openbrowser("https://github.com/ElecTwix") })

	bottombox := container.NewVBox(hello, msg, SndMsg)
	middlebox := container.NewHBox(ServerBtn, Github)
	timeoutbox := container.NewVBox(timeoutlabel, timeout)
	mainbox := container.NewVBox(BackBtn, middlebox, timeoutbox, bottombox)

	mainpage.SetContent(mainbox)

}

func Clientgui() {
	needsstop := false

	hello := widget.NewLabel("Console Log")

	timeout := widget.NewSlider(1, 300)
	timeoutlabel := widget.NewLabel("1")
	timeout.Value = 60
	timeout.Refresh()
	timeout.OnChanged = func(f float64) {

		timeoutlabel.Text = fmt.Sprintf("Time out %f", f)
		timeoutlabel.Refresh()
	}

	timeoutlabel.Text = fmt.Sprintf("Time out %f", timeout.Value)
	timeoutlabel.Refresh()

	ip := widget.NewEntry()
	ip.PlaceHolder = "ip"
	ip.OnChanged = func(s string) {
		if !isnumber(s) {
			ip.Text = "only numbers"
		}
	}

	port := widget.NewEntry()
	port.PlaceHolder = "port"
	port.OnChanged = func(s string) {
		if !isnumber(s) {
			port.Text = "only numbers"
		}
	}

	msg := widget.NewEntry()
	msg.PlaceHolder = "Packet Message"
	BackBtn := widget.NewButton("Back", mainmenu)

	SndMsg := widget.NewButton("Send", func() { ClientSendMsg(msg) })
	SndMsg.Disable()

	var ServerBtn *widget.Button
	ServerBtn = widget.NewButton("Connect", func() {

		go Client(hello, ServerBtn, &needsstop, timeout.Value, ip.Text, port.Text)
		SndMsg.Enable()
		ServerBtn.Disable()

	})

	Github := widget.NewButton("Github", func() { openbrowser("https://github.com/ElecTwix") })

	bottombox := container.NewVBox(hello, msg, SndMsg)

	upbox := container.NewGridWithColumns(2, container.NewVBox(ip), container.NewVBox(port))

	middlebox := container.NewHBox(ServerBtn, Github)
	timeoutbox := container.NewVBox(timeoutlabel, timeout)
	mainbox := container.NewVBox(BackBtn, upbox, middlebox, timeoutbox, bottombox)

	mainpage.SetContent(mainbox)
}
