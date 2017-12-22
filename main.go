package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jessevdk/go-flags"
)

const VERSION = "0.2"

var opts struct {
	Args struct {
		From string `positional-arg-name:"from" description:"version to convert from"`
		To   string `positional-arg-name:"to" description:"version to convert to"`
		File string `positional-arg-name:"file" description:"docker-compose yaml file"`
	} `positional-args:"true" required:"2"`
}

func printUsage() {
	fmt.Println(VERSION)
	fmt.Println("")
	fmt.Printf("Usage: %s from to file\n", os.Args[0])
	fmt.Println("\tfrom - version to convert from")
	fmt.Println("\tto - version to convert to")
	fmt.Println("\tfile - path to docker-compose file")
	fmt.Println("")
	fmt.Printf("Example: %s v1 v3.2 docker-compose.yml\n", os.Args[0])
}

func validateInput(from string, to string) error {
	convs := getConverters()

	_, ok := convs[from]
	if ok {
		_, ok := convs[from][to]
		if ok {
			return nil
		}
	}

	return unknownInputError()
}

func main() {
	_, err := flags.Parse(&opts)

	if err != nil {
		printUsage()
		os.Exit(1)
	}

	err = validateInput(opts.Args.From, opts.Args.To)
	if err != nil {
		printAndExit(err)
	}

	f, err := loadFile(opts.Args.File)
	if err != nil {
		printAndExit(err)
	}

	if err = convert(opts.Args.From, opts.Args.To, f); err != nil {
		printAndExit(err)
	}
}

func convert(from string, to string, compose []byte) error {
	f := getConverters()[from][to]

	out, err := f(&compose)
	if err != nil {
		return err
	}

	fmt.Println(string(out))
	return nil
}

func loadFile(path string) ([]byte, error) {
	f, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	return f, nil
}
