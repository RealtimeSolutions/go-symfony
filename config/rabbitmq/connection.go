package rabbitmq

func (c *Config) UseConnection(name string) *Connection {
	if c.Connections == nil {
		c.Connections = make(map[string]*Connection)
	}

	for key, conn := range c.Connections {
		if key == name {
			return conn
		}
	}

	conn := &Connection{}
	c.Connections[name] = conn
	return conn
}

type Connection struct {
	Host     string
	Port     int `yaml:"port,omitempty"`
	User     string
	Password string
	Vhost    string
	Lazy     bool `yaml:"lazy,omitempty"`
}
