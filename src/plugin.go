package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	color "github.com/fatih/color"
)

type plugin string

var introMsg = "Cleaning up MY_FILES.." // This is the message that will be displayed to the user

func (g plugin) Cleanup() {
	fmt.Println(introMsg)
	DeleteFiles("~/DIR_YOU_WANT_TO_BE_CLEANED") // DIR_YOU_WANT_TO_BE_CLEANED is the path to the directory you want to be cleaned
}

func DeleteFiles(dir string, message string) error {
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

	color.Cyan(message + "...")
	return nil
}

func main() {
	plugin := plugin("")
	plugin.Cleanup()
}



var Cleanup plugin
