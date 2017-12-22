package main

import yaml "gopkg.in/yaml.v2"

// Convertable parses input as array of bytes into "from" version, then calls
// convert function which needs to be implemented in specific version conversion
// then it marshals yaml into "out" version and returns the array of bytes of
// that yml file
// Args:
// in interface{} - composev1, composev23 etc type
// out interface{} - composev1, composev23 etc type
// bytes - input docker-compose.yml as array of bytes
// f - func which populates out interface{}
// Returns []byte, error
// Example usage:
// return Convertable(&in, &out, bytes, func() {
//   out.Version = "2.3"
// })
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

type converter map[string]func(bytes *[]byte) ([]byte, error)

func getConverters() map[string]converter {
	var converters = make(map[string]converter)

	converters["v1"] = converter{
		"v2.3": v1tov23,
		"v3.2": v1tov32,
	}

	return converters
}
