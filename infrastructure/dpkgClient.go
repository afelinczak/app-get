package infrastructure

import (
	"fmt"
	"os/exec"
)

// InstallApp runs dpkg using sudo
func InstallApp(path string) {
	fmt.Println("Try to install deb package " + path)
	var cmd = exec.Command("/bin/sh", "-c", "sudo dpkg -i "+path)
	var err = cmd.Run()
	if err != nil {
		fmt.Println(err)
		fmt.Println("installed sucessfully")
	}
}
