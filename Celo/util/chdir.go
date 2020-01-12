package util

import (
	"fmt"
	"log"
	"os"
)

func ChangeDir(dir string) {
	homeDir := os.Getenv("HOME")
	fullDir := homeDir + dir
	fmt.Println("\nChanging directory to \"", fullDir, "\"")
	if err := os.Chdir(fullDir); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("\n\u2713\u2713\u2713\u2713\u2713\u2713Ran successfully\u2713\u2713\u2713\u2713\u2713\u2713")
		fmt.Println("-----")
	}
}
