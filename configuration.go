package main

import (
	"os"
	"path/filepath"

	"code.google.com/p/gcfg"
)

type configuration struct {
	AccessKey string `gcfg:"access-key"`
	SecretKey string `gcfg:"secret-key"`
}

func getConfig(fileName string) (*configuration, error) {
	// first check directory where the executable is located
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return nil, err
	}
	path := dir + "/" + fileName
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// fallback to working directory. This is usefull when using `go run`
		path = fileName
	}

	var cfg struct {
		AWS configuration
	}

	err = gcfg.ReadFileInto(&cfg, path)
	if err != nil {
		return nil, err
	}

	return &cfg.AWS, nil
}
