package utils

import (
	"fmt"
	"log"
	"os"
)

func CreateDir(dirPath string, reCreate bool) {
	log.Println("Create Dir : ", dirPath)
	if _, err := os.Stat(dirPath); !os.IsNotExist(err) && reCreate {
		fmt.Println("Dir is exists")
		err := os.RemoveAll(dirPath)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Create Environment
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.Mkdir(dirPath, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		log.Print("Dir created")
	}
}
