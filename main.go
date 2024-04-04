package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func walkDir(path string) error {
	return filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			fmt.Println("Directory: ", path)
		} else {
			fmt.Println("File: ", path)
		}
		return nil
	})
}

func main() {
	publicDir := "./public"

	err := walkDir(publicDir)
	if err != nil {
		fmt.Println("Error walking directory: ", err)
	}

}
