package elevator

import (
	"fmt"
	"os"
	"syscall"
)

func CheckAndRequestElevation(message_text string) (err error) {
	if os.Getuid() != 0 {
		fmt.Println(message_text)
		// The software will restart itself but with sudo
		syscall.Exec("/usr/bin/sudo", []string{"-S", os.Args[0]}, []string{})
		//_ = exec.Command(, "-c", "sudo -S " + os.Args[0]).Start()
	}
	os.Exit(1)
	return
}
