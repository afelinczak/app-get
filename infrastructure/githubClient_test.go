package infrastructure

import (
	"testing"
)

func Is64BitDebFile_686File_ReturnsFalse(t *testing.T) {
	var fileName = "bat_0.22.1_arm64.deb"

	if is64BitDebFile(fileName) != false {
		t.Fatalf(fileName + ` should be recognized as arm deb package`)
	}
}

func Is64BitDebFile_amd64ile_ReturnsFalse(t *testing.T) {
	var fileName = "bat_0.22.1_amd64.deb"

	if is64BitDebFile(fileName) != true {
		t.Fatalf(fileName + ` should be recognized as amd deb package`)
	}
}

func getShortAppName_returnNameForBat(t *testing.T) {
	var fileName = "bat-musl_0.22.1_amd64.deb"

	if getShortAppName(fileName) != "bat-musl" {
		t.Fatalf(fileName + ` should be recognized as bat`)
	}
}

func getShortAppName_returnNameForDBeaver(t *testing.T) {
	var fileName = "dbeaver-ce_22.3.3_amd64.deb"

	if getShortAppName(fileName) != "dbeaver-ce" {
		t.Fatalf(fileName + ` should be recognized as dbeaver-ce`)
	}
}
