package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strconv"
	"strings"
)

const (
	cli = "release"
)

var (
	err            error
	message        string
	increment      string
	module         string
	isFirstRelease bool
)

var (
	increments = []string{
		"patch",
		"minor",
	}
)

func usage() {
	fmt.Printf("Usage: %s [flags]\n\n", cli)

	flag.PrintDefaults()
}

func fatal(where, what string) {
	fmt.Printf("Error(%s): %s.\n", where, what)
	os.Exit(1)
}

func latestTag(module string) (string, error) {
	output, err := exec.Command("bash", "-c", fmt.Sprintf("git tag --sort=-taggerdate | grep %s | head -n 1", module)).Output()

	switch {
	case err != nil:
		return "", err
	case len(output) == 0:
		return "", fmt.Errorf("no previous release found for %q", module)
	}

	return strings.TrimRight(string(output), "\n"), nil
}

func bump(latest, increment string) (string, error) {
	actualVersion := strings.Split(latest, "v")[1]

	semver := strings.Split(actualVersion, ".")

	if len(semver) != 3 {
		return "", fmt.Errorf("%q does not follow the semver convention", actualVersion)
	}

	patch, errPatch := strconv.Atoi(semver[2])
	minor, errMinor := strconv.Atoi(semver[1])
	major, errMajor := strconv.Atoi(semver[0])

	if err = errors.Join(errPatch, errMinor, errMajor); err != nil {
		return "", err
	}

	switch increment {
	case "patch":
		patch++
	case "minor":
		minor++
	}

	return fmt.Sprintf("%d.%d.%d", major, minor, patch), nil
}

func commit(module, version string) error {
	if err := exec.Command("git", "commit", "--allow-empty", "-m", fmt.Sprintf("chore(release): %s/v%s", module, version)).Run(); err != nil {
		return err
	}

	return nil
}

func tag(version, module string) error {
	if err := exec.Command("git", "tag", "-a", fmt.Sprintf("%s/v%s", module, version), "-m", fmt.Sprintf("%s %s", module, version)).Run(); err != nil {
		return err
	}

	return nil
}

func main() {
	flag.StringVar(&increment, "i", "", "Increment \"patch\" or \"minor\" (required)")

	flag.StringVar(&module, "m", "", "Module name (required)")

	flag.BoolVar(&isFirstRelease, "f", false, "Is the first release (default: false)")

	flag.Usage = usage

	flag.Parse()

	switch {
	case flag.NFlag() == 0:
		flag.Usage()
		println()
		message = "define required flags"
	case !slices.Contains(increments, increment):
		message = fmt.Sprintf("%q is not valid, allowed values are \"patch\" or \"minor\"", increment)
	case strings.TrimSpace(module) == "":
		message = "module name is required"
	}

	if message != "" {
		fatal("flags", message)
		return
	}

	version := "0.1.0"

	if !isFirstRelease {
		latest, err := latestTag(module)

		if err != nil {
			fatal("latestTag", err.Error())
		}

		version, err = bump(latest, increment)

		if err != nil {
			fatal("bump", err.Error())
		}
	}

	err = commit(module, version)

	if err != nil {
		fatal("commit", err.Error())
	}

	err = tag(version, module)

	if err != nil {
		fatal("tag", err.Error())
	}

	fmt.Printf("Successfully released \"%s %s\"\n", module, version)
}
