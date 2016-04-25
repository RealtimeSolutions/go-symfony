package datatable

type Datatable struct {
	TicketStore TicketStore
	Cache       Cache
}

type TicketStore struct {
	Backend string
	Host    string `yaml:"port,omitempty"`
}

type Cache struct {
	Backend string
	Host    string `yaml:"port,omitempty"`
}
