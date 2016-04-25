package rabbitmq

func (c *Config) UseProducer(name string) *Producer {
	if c.Producers == nil {
		c.Producers = make(map[string]*Producer)
	}

	for key, conn := range c.Producers {
		if key == name {
			return conn
		}
	}

	conn := &Producer{}
	c.Producers[name] = conn
	return conn
}

type Producer struct {
	Connection string
	Exchange   ProducerExchangeOptions `yaml:"exchange_options"`
}

type ProducerExchangeOptions struct {
	Name string
	Type string
}
