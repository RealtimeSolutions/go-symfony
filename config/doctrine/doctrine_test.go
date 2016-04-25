package doctrine

import (
	"gopkg.in/yaml.v2"
	"log"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	var data = `
    dbal:
        connections:
            acl:
                dbname: test2
            web:
                dbname: test
`

	c := Doctrine{}
	err := yaml.Unmarshal([]byte(data), &c)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	if len(c.Dbal.Connections) != 2 {
		t.Fail()
	}

	if c.Dbal.Connections["web"].Name != "test" {
		t.Fail()
	}

	if c.Dbal.Connections["acl"].Name != "test2" {
		t.Fail()
	}
}

func TestMarshal(t *testing.T) {
	c := Doctrine{
		Dbal: Dbal{
			Connections: map[string]*Database{
				"acl": &Database{},
			},
		},
	}

	d, err := yaml.Marshal(&c)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	expected := `dbal:
  default_connection: ""
  connections:
    acl:
      dbname: ""
`
	if expected != string(d) {
		t.Fail()
	}
}
