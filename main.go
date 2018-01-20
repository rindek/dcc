package main

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
)

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

func main() {
	_, err := flags.Parse(&opts)

	if err != nil {
		printUsage()
		os.Exit(1)
	}

	input := Input{From: opts.Args.From,
		To:   opts.Args.To,
		File: opts.Args.File,
	}

	if err = input.loadFile(); err != nil {
		printAndExit(err)
	}

	output, err := input.Convert()

	if err != nil {
		printAndExit(err)
	}

	output.Print()
}
