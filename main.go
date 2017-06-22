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

var availableFrom []string = []string{"v1"}
var availableTo []string = []string{"v3.2"}

func printAndExit(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func validateVersion(ver string, available []string, txt string) error {
	for _, i := range available {
		if ver == i {
			return nil
		}
	}

	errmsg := fmt.Sprintf("Unknown version %s for %s, available ones are: %s", ver, txt, strings.Join(available, ", "))

	return errors.New(errmsg)
}

func main() {
	_, err := flags.Parse(&opts)

	if err != nil {
		printUsage()
		os.Exit(1)
	}

	err = func() error {
		if e := validateVersion(opts.Args.From, availableFrom, "input"); e != nil {
			return e
		}

		if e := validateVersion(opts.Args.To, availableTo, "output"); e != nil {
			return e
		}

		return nil
	}()

	if err != nil {
		printAndExit(err)
	}

	f, err := loadFile(opts.Args.File)
	if err != nil {
		printAndExit(err)
	}

	convert(opts.Args.From, opts.Args.To, f)
}

func convert(from string, to string, compose []byte) error {
	var out interface{}

	if from == "v1" {
		var c composev1

		err := yaml.Unmarshal(compose, &c)
		if err != nil {
			return err
		}

		switch to {
		case "v3.2":
			out = v1tov32(&c)
		default:
			return fmt.Errorf("Wrong version %s", to)
		}
	}

	newver, err := yaml.Marshal(&out)

	if err != nil {
		return err
	}

	fmt.Println(string(newver))
	return nil
}

func loadFile(path string) ([]byte, error) {
	f, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	return f, nil
}
