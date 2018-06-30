package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/golang/glog"
)

// image stores the name of the image we pull help from
var image string

// init houses flag related functionality
func init() {
	flag.Parse()

	if flag.NArg() == 1 {
		image = flag.Args()[0]
	} else {
		fmt.Println("You must provide an image:tag")
		flag.PrintDefaults()
		os.Exit(1)
	}
}

// cleanUp removes the temporary container used
func cleanUp(container string) {
	output, err := execCommand("docker", "rm", container)
	if err != nil {
		glog.Errorf("Unable to remove temporary container %s: %s %s", container, output, err.Error())
	}
}

// execCommand takes care of execing local commands and returning th results
func execCommand(command ...string) (string, error) {
	if glog.V(2) {
		fmt.Println(command[0], command[1:len(command)])
	}
	cmd := exec.Command(command[0], command[1:len(command)]...)
	out, err := cmd.Output()
	return strings.TrimRight(string(out), "\n"), err
}

// main is the main entry point for the command
func main() {
	container, err := execCommand("docker", "run", "-ti", "-d", "--entrypoint=/bin/sh", image, "-c \"sleep 1000\"")
	defer cleanUp(container)
	if err != nil {
		panic(err.Error())
	}
	file, err := ioutil.TempFile(os.TempDir(), "iman")
	file.Close()
	_, err = execCommand("docker", "cp", fmt.Sprintf("%s:/help.1", container), file.Name())
	if err != nil {
		panic(err.Error())
	}
	manContent, err := execCommand("/usr/bin/bash", "-c", fmt.Sprintf("man -P /bin/cat -l %s", file.Name()))
	fmt.Println(manContent)
}
