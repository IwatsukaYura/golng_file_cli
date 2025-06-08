package pkg

import (
	"fmt"
	"os"
)

func GetFileInfo(filename string) error {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		fmt.Printf("ERROR : could not get file info for %s\n", filename)
		return err
	}

	fmt.Println("File Name : ", fileInfo.Name())
	fmt.Println("Size in Bytes : ", fileInfo.Size())
	fmt.Println("Permissions : ", fileInfo.Mode())
	fmt.Println("Last Modified : ", fileInfo.ModTime())
	fmt.Println("Is Directory : ", fileInfo.IsDir())
	return nil
}
