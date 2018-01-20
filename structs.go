package main

import (
	"fmt"
	"io/ioutil"
)

type Input struct {
	From string
	To   string
	File string
	In   []byte
}

func (i *Input) loadFile() error {
	f, err := ioutil.ReadFile(i.File)

	if err != nil {
		return err
	}

	i.In = f

	return nil
}

func (i Input) Convert() (Output, error) {
	switch i.From {
	case "v1":
		from := composev1{Input: &i}
		return from.Convert()
	}

	return Output{}, unknownInputError()
}

type Output struct {
	Out []byte
}

func (o Output) Print() {
	fmt.Println(string(o.Out))
}
