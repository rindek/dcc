package main

import (
	"errors"
	"fmt"
	"os"
)

func StringArray(str string) []string {
	out := []string{str}

	if str == "" {
		return nil
	}

	return out
}

func printAndExit(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func unknownInputError() error {
	errmsg := "Invalid converter versions"

	return errors.New(errmsg)
}
