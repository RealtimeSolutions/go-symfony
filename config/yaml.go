package config

import (
	"io"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func (c *Config) readFromYaml(r io.Reader) error {
	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(bytes, &c)
}

func (c *Config) writeToYaml(w io.Writer) error {
	data, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	n, err := w.Write(data)
	if err != nil {
		return err
	}

	if n != len(data) {
		return io.ErrUnexpectedEOF
	}

	return nil
}
