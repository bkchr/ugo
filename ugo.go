package main

import (
	"fmt"
	git "gopkg.in/src-d/go-git.v4"
	"os"
	"path"
)

const GOPKG_FILENAME = ".gopackage"

func find_gopackage_path() (string, error) {
	current_dir, err := os.Getwd()

	if err != nil {
		return "", err
	}

	for current_dir != "" {
		if _, err := os.Stat(path.Join(current_dir, GOPKG_FILENAME)); err == nil {
			return current_dir, nil
		}

		current_dir, _ = path.Split(current_dir)
	}

	return "", fmt.Errorf("Could not find a %s in any parent directory up to the root directory.", GOPKG_FILENAME)
}

func guess_package_name() (string, error) {
	current_dir, err := os.Getwd()

	if err != nil {
		return "", err
	}

	for current_dir != "" {
		repo, err := git.PlainOpen(current_dir)

		if err != nil {
			continue
		}

		for _, remote := range repo.Config().Remotes {
			
		}

		current_dir, _ = path.Split(current_dir)
	}

}

func main() {
	fmt.Printf("Hello, world.\n")
}
