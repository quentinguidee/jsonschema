package jsonschema

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestTypes_Unmarshal(t *testing.T) {
	type out struct {
		Type Types `yaml:"type"`
	}

	cases := []struct {
		inYAML []byte
		inJSON []byte
		out    out
		len    int
	}{
		{
			inYAML: []byte(`type: object`),
			inJSON: []byte(`{"type":"object"}`),
			out: out{
				Type: NewTypes(TypeObject),
			},
			len: 1,
		},
		{
			inYAML: []byte(`type: 'null'`),
			inJSON: []byte(`{"type":"null"}`),
			out: out{
				Type: NewTypes(TypeNull),
			},
			len: 1,
		},
		{
			inYAML: []byte(`type: [object, 'null']`),
			inJSON: []byte(`{"type":["object","null"]}`),
			out: out{
				Type: NewTypes(TypeObject, TypeNull),
			},
			len: 2,
		},
	}

	for _, c := range cases {
		o := out{}
		err := yaml.Unmarshal(c.inYAML, &o)
		assert.NoError(t, err)
		assert.Equal(t, c.out.Type, o.Type)
		assert.Equal(t, c.len, o.Type.Len())

		o = out{}
		err = json.Unmarshal(c.inJSON, &o)
		assert.NoError(t, err)
		assert.Equal(t, c.out.Type, o.Type)
	}
}
