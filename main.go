package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v2"

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

func availableFrom() []string {
	versions := []string{"v1"}

	return versions
}

func availableTo() []string {
	versions := []string{"v3.2"}

	return versions
}

func validateVersion(ver string, available []string, txt string) error {
	for _, i := range available {
		if ver == i {
			return nil
		}
	}

	errmsg := fmt.Sprintf("Unknown version %s for %s, available ones are: %s", ver, txt, strings.Join(available, ", "))
	fmt.Println(errmsg)

	return errors.New(errmsg)
}

func main() {
	_, err := flags.Parse(&opts)

	if err != nil {
		printUsage()
		os.Exit(1)
	}

	errf := validateVersion(opts.Args.From, availableFrom(), "input")
	errt := validateVersion(opts.Args.To, availableTo(), "output")

	if errf != nil || errt != nil {
		os.Exit(1)
	}

	f := loadFile(opts.Args.File)

	convert(opts.Args.From, opts.Args.To, f)
}

func convert(from string, to string, compose []byte) {
	var out interface{}

	if from == "v1" {
		var c composev1

		err := yaml.Unmarshal(compose, &c)
		if err != nil {
			panic(err)
		}

		switch to {
		case "v3.2":
			out = v1tov32(&c)
		}
	}

	newver, err := yaml.Marshal(&out)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(newver))
}

func loadFile(path string) []byte {
	f, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return f
}
