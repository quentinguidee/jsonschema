package jsonschema

type Schema struct {
	Title       string        `json:"title,omitempty"       yaml:"title,omitempty"`
	Description string        `json:"description,omitempty" yaml:"description,omitempty"`
	Ref         string        `json:"$ref,omitempty"        yaml:"$ref,omitempty"`
	Type        Types         `json:"type,omitempty"        yaml:"type,omitempty"`
	Format      Format        `json:"format,omitempty"      yaml:"format,omitempty"`
	Default     interface{}   `json:"default,omitempty"     yaml:"default,omitempty"`
	Example     []interface{} `json:"examples,omitempty"    yaml:"examples,omitempty"`

	Const string   `json:"const,omitempty" yaml:"const,omitempty"`
	Enum  []string `json:"enum,omitempty"  yaml:"enum,omitempty"`

	Items            *Schema `json:"items,omitempty"            yaml:"items,omitempty"`
	UniqueItems      bool    `json:"uniqueItems,omitempty"      yaml:"uniqueItems,omitempty"`
	UnevaluatedItems bool    `json:"unevaluatedItems,omitempty" yaml:"unevaluatedItems,omitempty"`

	Properties            map[string]*Schema `json:"properties,omitempty"            yaml:"properties,omitempty"`
	PropertyNames         *Schema            `json:"propertyNames,omitempty"         yaml:"propertyNames,omitempty"`
	AdditionalProperties  *Schema            `json:"additionalProperties,omitempty"  yaml:"additionalProperties,omitempty"`
	UnevaluatedProperties bool               `json:"unevaluatedProperties,omitempty" yaml:"unevaluatedProperties,omitempty"`

	AllOf []Schema `json:"allOf,omitempty" yaml:"allOf,omitempty"`
	AnyOf []Schema `json:"anyOf,omitempty" yaml:"anyOf,omitempty"`
	OneOf []Schema `json:"oneOf,omitempty" yaml:"oneOf,omitempty"`
	Not   *Schema  `json:"not,omitempty"   yaml:"not,omitempty"`

	ReadOnly  bool `json:"readOnly,omitempty"  yaml:"readOnly,omitempty"`
	WriteOnly bool `json:"writeOnly,omitempty" yaml:"writeOnly,omitempty"`

	Deprecated bool `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	Required   bool `json:"required,omitempty"   yaml:"required,omitempty"`

	If   *Schema `json:"if,omitempty"   yaml:"if,omitempty"`
	Then *Schema `json:"then,omitempty" yaml:"then,omitempty"`
	Else *Schema `json:"else,omitempty" yaml:"else,omitempty"`

	Contains          *Schema             `json:"contains,omitempty"          yaml:"contains,omitempty"`
	ContentEncoding   ContentEncoding     `json:"contentEncoding,omitempty"   yaml:"contentEncoding,omitempty"`
	ContentMediaType  string              `json:"contentMediaType,omitempty"  yaml:"contentMediaType,omitempty"`
	ContentSchema     *Schema             `json:"contentSchema,omitempty"     yaml:"contentSchema,omitempty"`
	DependentRequired map[string][]string `json:"dependentRequired,omitempty" yaml:"dependentRequired,omitempty"`
	DependentSchemas  *Schema             `json:"dependentSchemas,omitempty"  yaml:"dependentSchemas,omitempty"`
	MinContains       *int                `json:"minContains,omitempty"       yaml:"minContains,omitempty"`
	MaxContains       *int                `json:"maxContains,omitempty"       yaml:"maxContains,omitempty"`
	Minimum           *float64            `json:"minimum,omitempty"           yaml:"minimum,omitempty"`
	Maximum           *float64            `json:"maximum,omitempty"           yaml:"maximum,omitempty"`
	ExclusiveMinimum  *float64            `json:"exclusiveMinimum,omitempty"  yaml:"exclusiveMinimum,omitempty"`
	ExclusiveMaximum  *float64            `json:"exclusiveMaximum,omitempty"  yaml:"exclusiveMaximum,omitempty"`
	MinLength         *int                `json:"minLength,omitempty"         yaml:"minLength,omitempty"`
	MaxLength         *int                `json:"maxLength,omitempty"         yaml:"maxLength,omitempty"`
	MinProperties     *int                `json:"minProperties,omitempty"     yaml:"minProperties,omitempty"`
	MaxProperties     *int                `json:"maxProperties,omitempty"     yaml:"maxProperties,omitempty"`
	MinItems          *int                `json:"minItems,omitempty"          yaml:"minItems,omitempty"`
	MaxItems          *int                `json:"maxItems,omitempty"          yaml:"maxItems,omitempty"`
	MultipleOf        *int                `json:"multipleOf,omitempty"        yaml:"multipleOf,omitempty"`
	Pattern           string              `json:"pattern,omitempty"           yaml:"pattern,omitempty"`
	PatternProperties map[string]Schema   `json:"patternProperties,omitempty" yaml:"patternProperties,omitempty"`
	PrefixItems       []Schema            `json:"prefixItems,omitempty"       yaml:"prefixItems,omitempty"`
}
