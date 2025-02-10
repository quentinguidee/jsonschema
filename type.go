package jsonschema

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

type Type string

const (
	TypeString  Type = "string"
	TypeNumber  Type = "number"
	TypeInteger Type = "integer"
	TypeObject  Type = "object"
	TypeArray   Type = "array"
	TypeBoolean Type = "boolean"
	TypeNull    Type = "null"
)

type Types struct {
	values []Type
}

func NewTypes(values ...Type) Types {
	return Types{
		values: values,
	}
}

func (t *Types) Values() []Type {
	return t.values
}

func (t *Types) Len() int {
	return len(t.values)
}

func (t *Types) UnmarshalYAML(value *yaml.Node) error {
	var err error
	if value.Kind == yaml.SequenceNode {
		err = value.Decode(&t.values)
	} else {
		var s string
		err = value.Decode(&s)
		t.values = append(t.values, Type(s))
	}
	return err
}

func (t *Types) UnmarshalJSON(data []byte) error {
	var err error
	if len(data) != 0 && data[0] == '[' {
		err = json.Unmarshal(data, &t.values)
	} else {
		var s string
		err = json.Unmarshal(data, &s)
		t.values = append(t.values, Type(s))
	}
	return err
}
