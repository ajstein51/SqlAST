package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("Program Start")

	test := flag.Bool("test", false, "Run tests")
	flag.Parse()

	if *test {
		fmt.Println("Running tests")
		
		cmd := exec.Command("go", "test", "./test/main.go")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()

		if err != nil {
			panic(err)
		}

		return
	}

	fmt.Println("Program End")	
}