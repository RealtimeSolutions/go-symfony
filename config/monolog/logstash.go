package monolog

type LogstashHandler struct {
	Type             string
	Level            string
	Persistent       bool
	Timeout          int
	ConnectionString string `yaml:"connection_string"`
	Formatter        string
}
