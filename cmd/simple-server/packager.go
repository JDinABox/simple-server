package main

import (
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/allocamelus/allocamelus/pkg/logger"
	"k8s.io/klog/v2"
)

func runDev() *exec.Cmd {
	packagerType := getPkgType()
	cmd := getCmd(packagerType, "run", "dev")
	cmd.Start()
	return cmd
}

func runBuild() {
	packagerType := getPkgType()
	err := getCmd(packagerType, "run", "build").Run()
	logger.Fatal(err)
}

func getPkgType() (packagerType string) {
	logger.Fatal(filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if filepath.Dir(path) != "." || info.IsDir() {
			return nil
		}

		switch info.Name() {
		case "yarn.lock":
			packagerType = "yarn"
		case "package-lock.json":
			packagerType = "npm"
		}
		return nil
	}))
	if packagerType == "" {
		klog.Fatalf("Missing lock file, try running %s init", os.Args[0])
	}
	return
}
