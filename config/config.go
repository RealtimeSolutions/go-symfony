package config

import (
	"os"
	"realtime.eu/symfony/config/datatable"
	"realtime.eu/symfony/config/doctrine"
	"realtime.eu/symfony/config/monolog"
	"realtime.eu/symfony/config/rabbitmq"
)

type Config struct {
	filename string
	Doctrine doctrine.Doctrine
	Common   Common
	Wms      Wms
	Rabbitmq rabbitmq.Config `yaml:"old_sound_rabbit_mq"`
	Monolog  monolog.Config
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
