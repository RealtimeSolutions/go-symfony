package monolog

type Config struct {
	Handlers map[string]Handler
}

type Handler interface{}

func (c *Config) SetHandler(name string, h Handler) {
	if c.Handlers == nil {
		c.Handlers = make(map[string]Handler, 1)
	}

	c.Handlers[name] = h
}

func (c *Config) DeleteHandler(name string) {
	delete(c.Handlers, name)
}

func (c *Config) GetHandler(name string) Handler {
	if handler, has := c.Handlers[name]; has {
		switch v := handler.(type) {
		case LogstashHandler:
			return LogstashHandler(v)

		case map[interface{}]interface{}:
			m := make(map[string]interface{}, len(v))
			for key, value := range v {
				m[key.(string)] = value
			}
			return m
		}
	}

	return nil
}
