package domain

import (
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestStandardVersion_ReturnsStableVersion(t *testing.T) {
	var version = "0.22.1"
	var expected = AppVersion{Major: 0, Minor: 22, Patch: 1, IsStable: true}

	if ParseVersion(version) != expected {
		t.Fatalf(`Version %v should be parsed properly`, version)
	}
}

func TestBetaVersion_ReturnsNonStableVersion(t *testing.T) {
	var version = "1.233.0-beta1"
	var expected = AppVersion{Major: 1, Minor: 233, Patch: 0, IsStable: false}

	if ParseVersion(version) != expected {
		t.Fatalf(`Version %v should be parsed properly`, version)
	}
}

func Test_IsAvailableVersionNewerAndStable_MajorPatchHigher_True(t *testing.T) {
	var versionInstalled = "1.233.0"
	var versionAvailable = "1.234.1"

	if IsAvailableVersionNewerAndStable(versionInstalled, versionAvailable) == false {
		t.Fatalf(`Version %va should be recognised as newer than %vi`, versionAvailable, versionInstalled)
	}
}

func Test_IsAvailableVersionNewerAndStable_Same_False(t *testing.T) {
	var versionInstalled = "1.233.0"
	var versionAvailable = "1.233.0"

	if IsAvailableVersionNewerAndStable(versionInstalled, versionAvailable) == true {
		t.Fatalf(`Version %va should be recognised as same as %vi`, versionAvailable, versionInstalled)
	}
}

func Test_IsAvailableVersionNewerAndStable_NewerIsBeta_False(t *testing.T) {
	var versionInstalled = "1.233.0"
	var versionAvailable = "1.233.1-beta"

	if IsAvailableVersionNewerAndStable(versionInstalled, versionAvailable) == true {
		t.Fatalf(`Version %va should be recognised as beta `, versionAvailable)
	}
}
