package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	simpleserver "github.com/JDinABox/simple-server"
	"github.com/JDinABox/simple-server/app"
	"github.com/allocamelus/allocamelus/pkg/logger"
	jsoniter "github.com/json-iterator/go"
	"k8s.io/klog/v2"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func setup() {
	// Select Package Manager
	var pkgManager string
	err := survey.AskOne(&survey.Select{
		Message: "Package Manager:",
		Options: []string{"yarn", "npm"},
		Default: "yarn",
	}, &pkgManager)
	if err != nil {
		klog.Fatal(err)
	}

	// Write example files
	err = writeFiles("")
	logger.Fatal(err)

	// Write Config
	conf := app.DefaultConfig()
	conf.Headers = append(conf.Headers, "ss.umd.js", "style.css")
	confJson, err := json.MarshalIndent(conf, "", "  ")
	logger.Fatal(err)
	writeFile("config.json", confJson)

	cmd := getCmd(pkgManager, "install")
	logger.Fatal(cmd.Run())

	cmd = getCmd(pkgManager, "run", "build")
	logger.Fatal(cmd.Run())

	fmt.Printf("\nSuccess! Run %s to start the server\n", os.Args[0])
}

func getCmd(name string, args ...string) *exec.Cmd {
	cmdPath, err := exec.LookPath(name)
	if err != nil {
		logger.Fatal(err)
	}
	return &exec.Cmd{
		Path:   cmdPath,
		Args:   append([]string{cmdPath}, args...),
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
}

func writeFiles(dir string) error {
	embeddedPath := filepath.Join("example", dir)

	fsList, err := simpleserver.Files.ReadDir(embeddedPath)
	if err != nil {
		return err
	}

	for _, f := range fsList {
		relPath := filepath.Join(dir, f.Name())
		if f.IsDir() {
			// Create folders if needed
			if err = os.MkdirAll(relPath, os.ModeSticky|os.ModePerm); err != nil {
				return err
			}

			// Iterate over deeper folder
			err = writeFiles(relPath)
			if err != nil {
				return err
			}
		} else {
			// Get embedded file
			fi, err := simpleserver.Files.ReadFile(filepath.Join(embeddedPath, f.Name()))
			if err != nil {
				return err
			}

			writeFile(relPath, fi)
		}
	}
	return nil
}

func writeFile(path string, data []byte) error {
	// Prompt if file exist
	override := false
	if _, err := os.Stat(path); err == nil {
		survey.AskOne(&survey.Confirm{
			Message: "Overide file " + path + "?",
			Default: false,
		}, &override)
	} else if os.IsNotExist(err) {
		override = true
	} else {
		return err
	}

	if override {
		if err := os.WriteFile(path, data, 0644); err != nil {
			return err
		}
	}
	return nil
}
