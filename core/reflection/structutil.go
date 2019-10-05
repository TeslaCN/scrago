package reflection

import "reflect"

type Setter interface {
	Set(fieldName string, value interface{}) error
}
type FieldSetter struct {
}

func (FieldSetter) Set(fieldName string, value interface{}) error {

	return nil
}

type MetaField struct {
	Name        string
	CssSelector string
	Attr        string
}

func ParseStruct(i interface{}) map[string]MetaField {
	fieldMaps := make(map[string]MetaField)
	t := reflect.TypeOf(i)
	for i := 0; i < t.NumField(); i++ {
		metaField := MetaField{}
		field := t.Field(i)
		jsonName, ok := field.Tag.Lookup("json")
		if ok {
			metaField.Name = jsonName
		} else {
			metaField.Name = field.Name
		}
		cssSelector, ok := field.Tag.Lookup("css_selector")
		if ok {
			metaField.CssSelector = cssSelector
		}
		attr, ok := field.Tag.Lookup("attr")
		if ok {
			metaField.Attr = attr
		}
		fieldMaps[metaField.Name] = metaField
	}
	return fieldMaps
}
