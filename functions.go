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
	convs := getConverters()

	errmsg := "Invalid converter versions, available are:\n"

	for from, tos := range convs {
		for to, _ := range tos {
			errmsg = errmsg + fmt.Sprintf("\t* %s -> %s\n", from, to)
		}
	}

	return errors.New(errmsg)
}
