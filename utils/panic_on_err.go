package utils

import "fmt"

func CheckForErrPanic(err error) {
	if err != nil {
		panic(fmt.Errorf("error found and panic raised: %v", err))
	}
}
