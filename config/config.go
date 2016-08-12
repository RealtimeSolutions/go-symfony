package config

import (
	"os"
	"github.com/RealtimeSolutions/symfonyGoConf/config/datatable"
	"github.com/RealtimeSolutions/symfonyGoConf/config/doctrine"
	"github.com/RealtimeSolutions/symfonyGoConf/config/monolog"
	"github.com/RealtimeSolutions/symfonyGoConf/config/rabbitmq"
)

func (c *Config) SetParameter(key, value string) {
	if c.Parameters == nil {
		c.Parameters = make(map[string]string)
	}

	c.Parameters[key] = value
}

type Config struct {
	filename   string
	Doctrine   doctrine.Doctrine
	Common     Common
	Wms        Wms
	Rabbitmq   rabbitmq.Config `yaml:"old_sound_rabbit_mq"`
	Monolog    monolog.Config
	Parameters map[string]string
}

type Common struct {
	Mongodb Mongodb
}

type Mongodb struct {
	Server string
}

type Wms struct {
	Datatable datatable.Datatable
}

func Open(name string) (*Config, error) {
	file, err := os.Open(name)
	c := &Config{filename: name}
	if os.IsNotExist(err) {
		return c, nil
	}

	if err == nil {
		defer file.Close()
		err = c.readFromYaml(file)
	}

	return c, err
}

func (c *Config) Save() error {
	file, err := os.Create(c.filename)
	defer file.Close()
	if err != nil {
		return err
	}

	if err := c.writeToYaml(file); err == nil {
		return file.Sync()
	} else {
		return err
	}
}
