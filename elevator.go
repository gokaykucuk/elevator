package elevator

import (
	"fmt"
	"os"
	"os/exec"
)

func CheckAndRequestElevation() (err error) {
	if os.Getuid() != 0 && os.Getenv("ELEVATOR_SECOND_ROUND") != "true" {
		// The binary will restart itself but with sudo
		environment_password := os.Getenv("ELEVATOR_SUDO_PASSWORD")
		if environment_password != "" { // Try to supply sudo password if it's in environment
			password_relay := exec.Command("echo", environment_password)
			cmd := exec.Command("/usr/bin/sudo", "-S", os.Args[0])
			cmd.Stdin, err = password_relay.StdoutPipe()
			password_relay.Start()
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Env = os.Environ()
			err = cmd.Run()
			//err := syscall.Exec("/usr/bin/sudo", []string{"-S",os.Args[0], environment_password + "\n"}, []string{"ELEVATOR_SECOND_ROUND=true"})
			//fmt.Println(err)
		} else {
			cmd := exec.Command("/usr/bin/sudo", os.Args[0])
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Env = os.Environ()
			err = cmd.Run()
		}
		os.Exit(0)
	} else {
		fmt.Println("!!! Sudo access granted !!!")
	}
	return
}
