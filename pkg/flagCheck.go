package pkg

import (
	"errors"
	"flag"
)

func ValidateNoExtraArgs() error {
	if len(flag.Args()) > 0 {
		return errors.New("invalid number of arguments")
	}
	return nil
}
