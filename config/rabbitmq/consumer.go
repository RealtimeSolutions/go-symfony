package rabbitmq

// UseConsumer ensures a consumer with given name is
// available and returns a pointer to that consumer.
func (c *Config) UseConsumer(name string) *Consumer {
	if c.Consumers == nil {
		c.Consumers = make(map[string]*Consumer)
	}

	for key, conn := range c.Consumers {
		if key == name {
			return conn
		}
	}

	conn := &Consumer{}
	c.Consumers[name] = conn
	return conn
}

type Consumer struct {
	Connection string
	Exchange   ConsumerExchangeOptions `yaml:"exchange_options"`
	Queue      QueueOptions            `yaml:"queue_options"`
	Qos        QosOptions              `yaml:"qos_options"`
	Callback   string
}

type ConsumerExchangeOptions struct {
	Name string
	Type string
}

type QueueOptions struct {
	Name string
}

type QosOptions struct {
	PrefetchSize  int  `yaml:"prefetch_size"`
	PrefetchCount int  `yaml:"prefetch_count"`
	Global        bool `yaml:"global"`
}
