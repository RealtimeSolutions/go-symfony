package rabbitmq

type Config struct {
	Connections map[string]*Connection
	Producers   map[string]*Producer
	Consumers   map[string]*Consumer
}
