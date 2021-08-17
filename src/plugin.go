package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type plugin string

var introMsg = "Cleaning up MY_FILES.." // This is the message that will be displayed to the user

func (g plugin) Cleanup() {
	fmt.Println(introMsg)
	DeleteFiles("~/DIR_YOU_WANT_TO_BE_CLEANED") // DIR_YOU_WANT_TO_BE_CLEANED is the path to the directory you want to be cleaned
}

func DeleteFiles(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}

func ShellCommand(command string) {
	out, err := exec.Command("/bin/sh", "-c", command).Output()
	if err != nil {
		fmt.Printf("error %s", err)
	}
	fmt.Printf("%s", out)
}

var Cleanup plugin
