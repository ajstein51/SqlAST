package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	fmt.Println("Test Program Start")

	err := filepath.Walk("test", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if info.Name() == "main.go" {
			cmd := exec.Command("go", "run", path)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			err := cmd.Run()

			if err != nil {
				panic(err)
			}
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	fmt.Println("Test Program End")
}