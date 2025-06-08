package main

import (
	"flag"
	"fmt"
	"golang_file_info/pkg"
	"os"
)

func main() {
	infoFlag := flag.String("info", "", "Show file info for a given file")
	listFlag := flag.String("list", "", "List files in a directory")
	flag.Parse()

	if *infoFlag != "" {
		err := pkg.ValidateNoExtraArgs()
		if err != nil {
			exitWithError(err)
		}
		err = pkg.GetFileInfo(*infoFlag)
		if err != nil {
			exitWithError(err)
		}
	} else if *listFlag != "" {
		err := pkg.ValidateNoExtraArgs()
		if err != nil {
			exitWithError(err)
		}
		err = pkg.GetListInDir(*listFlag)
		if err != nil {
			exitWithError(err)
		}
	}

}

func exitWithError(err error) {
	fmt.Printf("ERROR : %v\n", err)
	os.Exit(1)
}
