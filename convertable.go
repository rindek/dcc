package main

import yaml "gopkg.in/yaml.v2"

func Convertable(in interface{}, out interface{}, bytes *[]byte, f func()) ([]byte, error) {
	err := yaml.Unmarshal(*bytes, in)
	if err != nil {
		return nil, err
	}

	f()

	bytesout, err := yaml.Marshal(&out)
	if err != nil {
		return nil, err
	}

	return bytesout, nil
}
