package pkg

import (
	"fmt"
	"os"
)

func GetListInDir(dirPath string) error {

	files, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Printf("ERROR : could not list files in %s\n", dirPath)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}

	return nil
}
