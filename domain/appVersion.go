package domain

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type AppVersion struct {
	Major    int
	Minor    int
	Patch    int
	IsStable bool
}

// ParseVersion - parses standard version string like 1.2.31
func ParseVersion(version string) AppVersion {
	var chunks = strings.Split(version, ".")
	if len(chunks) < 3 {
		log.Fatal("Invalid app version " + version + ". Abort")
		os.Exit(110)
	}

	major, _ := strconv.Atoi(chunks[0])
	minor, _ := strconv.Atoi(chunks[1])

	var patchChunk = strings.Split(chunks[2], "-")
	patch, _ := strconv.Atoi(patchChunk[0])

	return AppVersion{Major: major, Minor: minor, Patch: patch, IsStable: len(patchChunk) == 1}
}

func IsAvailableVersionNewerAndStable(installedVersion string, availableVersion string) bool {
	var installedAppVersion = ParseVersion(installedVersion)
	var availableAppVersion = ParseVersion(availableVersion)

	if installedAppVersion.Major < availableAppVersion.Major && availableAppVersion.IsStable {
		return true
	}
	if installedAppVersion.Major == availableAppVersion.Major &&
		installedAppVersion.Minor < availableAppVersion.Minor &&
		availableAppVersion.IsStable {
		return true
	}
	if installedAppVersion.Major == availableAppVersion.Major &&
		installedAppVersion.Minor == availableAppVersion.Minor &&
		installedAppVersion.Patch < availableAppVersion.Patch &&
		availableAppVersion.IsStable {
		return true
	}
	return false
}

func (v AppVersion) String() string {
	return string(v.Major) + "." + string(v.Minor) + "." + string(v.Patch)
}
