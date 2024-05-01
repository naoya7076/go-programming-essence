package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func main() {
	files := []string{}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	err = filepath.WalkDir(cwd, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			if d.Name()[0] == '.' {
				return fs.SkipDir
			}
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(files)
}
