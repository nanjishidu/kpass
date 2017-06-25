package main

import (
	"kpass/asset"
	"os"
	"path/filepath"

	"github.com/nanjishidu/gomini"
)

// initialize static resources ， /template  /static
func initSite() error {
	dirs := []string{"template", "static"}
	isSuccess := true
	var (
		err error
	)
	for _, dir := range dirs {
		// is exist static resources
		if gomini.IsExist(filepath.Join("./", dir)) {
			continue
		}
		if err = asset.RestoreAssets("./", dir); err != nil {
			isSuccess = false
			break
		}
	}
	if !isSuccess {
		for _, dir := range dirs {
			os.RemoveAll(filepath.Join("./", dir))
		}
		return err
	}
	return nil
}
