package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

func isnumber(str string) bool {
	arr := "1234567890."

	var total int = 0
	for _, char := range str {

		for _, v := range arr {
			if char == v {
				total++
			}
		}

	}

	return total == len(str)

}

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}
