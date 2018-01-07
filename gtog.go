package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func dirContain(dir string, f string) (path string) {

	if dir == "/" {
		return ""
	}
	isPres, err := exists(filepath.Join(dir, f))
	if err != nil {
		panic("Error checking the existence of file")
	}
	if !isPres {
		return dirContain(filepath.Dir(dir), f)
	}

	return dir

}

func rename(dir, oldFile, newFile string) error {
	return os.Rename(filepath.Join(dir, oldFile), filepath.Join(dir, newFile))
}
func main() {
	cwd, err := os.Getwd()
	if err != nil {
		panic("Cannot get os.Getwd()")
	}

	fileA := ".git"
	fileB := ".xxxgit"

	fileADir := dirContain(cwd, fileA)
	fileBDir := dirContain(cwd, fileB)

	// does not exist
	if fileADir == "" && fileBDir == "" {
		fmt.Printf("%s or %s do not exist in your path\n", fileA, fileB)
		os.Exit(1)
	}

	// when git is off
	if fileBDir != "" {
		err = rename(fileBDir, fileB, fileA)
		if err != nil {
			panic("Cannot change name!")
		}
		os.Exit(0)
	}

	// when git is on
	err = rename(fileADir, fileA, fileB)
	if err != nil {
		panic("Cannot change name!")
	}

}
