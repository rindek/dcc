package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jessevdk/go-flags"
)

var opts struct {
	Args struct {
		From string `positional-arg-name:"version to convert from"`
		To   string `positional-arg-name:"version to convert to"`
		File string `positional-arg-name:"path to docker-compose file"`
	} `positional-args:"true" required:"2"`
}

func printUsage() {
	fmt.Println("")
	fmt.Printf("Usage: %s from to file\n", os.Args[0])
	fmt.Println("\tfrom - version to convert from")
	fmt.Println("\tto - version to convert to")
	fmt.Println("\tfile - path to docker-compose file")
	fmt.Println("")
	fmt.Printf("Example: %s v1 v3.2 docker-compose.yml\n", os.Args[0])
}

type converter map[string]func(bytes *[]byte) ([]byte, error)

func printAndExit(err error) {
	fmt.Println(err)
	os.Exit(1)
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

	return UnknownInputError()
}

func UnknownInputError() error {
	convs := getConverters()

	errmsg := "Invalid converter versions, available are:\n"

	for from, tos := range convs {
		for to, _ := range tos {
			errmsg = errmsg + fmt.Sprintf("\t* %s -> %s\n", from, to)
		}
	}

	return errors.New(errmsg)
}

func getConverters() map[string]converter {
	var converters = make(map[string]converter)

	converters["v1"] = converter{
		"v2.3": v1tov23,
		"v3.2": v1tov32,
	}

	return converters
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
