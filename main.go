/*
Copyright © 2022 Arkadiusz Felinczak <arek_felinczak@proton.me>
*/

package main

import (
	"github.com/afelinczak/app-get/cmd"
	"github.com/afelinczak/app-get/infrastructure"
)

func main() {
	infrastructure.CreateInstalledAppListFile()
	cmd.Execute()
}
