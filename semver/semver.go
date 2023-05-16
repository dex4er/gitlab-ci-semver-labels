package semver

import (
	"log"
	"strconv"

	"github.com/Masterminds/semver/v3"
)

func incrementNumberAsString(number string) string {
	num, err := strconv.Atoi(number)
	if err != nil {
		num = 0
	}
	result := strconv.Itoa(num + 1)
	return result
}

func BumpPrerelease(version string) (string, error) {
	ver, err := semver.NewVersion(version)
	if err != nil {
		return "", err
	}
	log.Printf("[DEBUG] BumpPrerelease(version=%v)\n", version)

	newVer, err := ver.SetPrerelease(incrementNumberAsString(ver.Prerelease()))
	if err != nil {
		log.Fatalf("[ERROR] Can't bump semver: %v\n", err)
	}

	return newVer.String(), nil
}

func BumpPatch(version string) (string, error) {
	ver, err := semver.NewVersion(version)
	if err != nil {
		return "", err
	}
	log.Printf("[DEBUG] BumpPatch(version=%v)\n", version)

	newVer := ver.IncPatch()

	return newVer.String(), nil
}

func BumpMinor(version string) (string, error) {
	ver, err := semver.NewVersion(version)
	if err != nil {
		return "", err
	}
	log.Printf("[DEBUG] BumpMinor(version=%v)\n", version)

	newVer := ver.IncMinor()

	return newVer.String(), nil
}

func BumpMajor(version string) (string, error) {
	ver, err := semver.NewVersion(version)
	if err != nil {
		return "", err
	}
	log.Printf("[DEBUG] BumpMajor(version=%v)\n", version)

	newVer := ver.IncMajor()

	return newVer.String(), nil
}
