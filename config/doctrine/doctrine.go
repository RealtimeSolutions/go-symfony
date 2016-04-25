package doctrine

type Doctrine struct {
	Dbal Dbal
}

func (d *Doctrine) UseDB(name string) *Database {
	if d.Dbal.Connections == nil {
		d.Dbal.Connections = make(map[string]*Database)
	}

	for key, db := range d.Dbal.Connections {
		if key == name {
			return db
		}
	}

	db := &Database{Name: name}
	d.Dbal.Connections[name] = db
	return db
}

type Dbal struct {
	Default_connection string
	Connections        map[string]*Database
}

type Database struct {
	Driver   string `yaml:"driver,omitempty"`
	Name     string `yaml:"dbname"`
	User     string `yaml:"user,omitempty"`
	Password string `yaml:"password,omitempty"`
	Charset  string `yaml:"charset,omitempty"`
	Host     string `yaml:"host,omitempty"`
	Port     uint16 `yaml:"port,omitempty"`
}
