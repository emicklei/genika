package {{.Package}}

// generated with github.com/emicklei/genika/gtg/{{.TemplateName}}

import (
	"bytes"
	"encoding/json"
)

// named{{.TypeName}} associates a name with a {{.TypeName}}
type named{{.TypeName}} struct {
	Name   string
	{{.TypeName}} {{.TypeName}}
}

// {{.TypeName}}Map encapsulates a list of named{{.TypeName}} (association) and maintains the insertion order.
type {{.TypeName}}Map struct {
	List []named{{.TypeName}}
}

// Put adds or replaces a {{.TypeName}} by its name
func (l *{{.TypeName}}Map) Put(name string, model {{.TypeName}}) {
	for i, each := range l.List {
		if each.Name == name {
			// replace
			l.List[i] = named{{.TypeName}}{name, model}
			return
		}
	}
	// add
	l.List = append(l.List, named{{.TypeName}}{name, model})
}

// At returns a {{.TypeName}} by its name, ok is false if absent
func (l {{.TypeName}}Map) At(name string) (m {{.TypeName}}, ok bool) {
	for _, each := range l.List {
		if each.Name == name {
			return each.{{.TypeName}}, true
		}
	}
	return m, false
}

// Do enumerates all the headers, each with its assigned name
func (l {{.TypeName}}Map) Do(block func(name string, value {{.TypeName}})) {
	for _, each := range l.List {
		block(each.Name, each.{{.TypeName}})
	}
}

// MarshalJSON writes the {{.TypeName}}Map as if it was a map[string]{{.TypeName}}
func (l {{.TypeName}}Map) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString("{\n")
	for i, each := range l.List {
		buf.WriteString("\"")
		buf.WriteString(each.Name)
		buf.WriteString("\": ")
		data, err := json.MarshalIndent(each.{{.TypeName}},"","\t")
		if err != nil {
			return buf.Bytes(), err
		}
		buf.Write(data)
		if i < len(l.List)-1 {
			buf.WriteString(",\n")
		}
	}
	buf.WriteString("}")
	return buf.Bytes(), nil
}

// UnmarshalJSON reads back a {{.TypeName}}Map. This is an expensive operation.
func (l *{{.TypeName}}Map) UnmarshalJSON(data []byte) error {
	raw := map[string]interface{}{}
	json.NewDecoder(bytes.NewReader(data)).Decode(&raw)
	for k, v := range raw {
		// produces JSON bytes for each value
		data, err := json.Marshal(v)
		if err != nil {
			return err
		}
		var m {{.TypeName}}
		json.NewDecoder(bytes.NewReader(data)).Decode(&m)
		l.Put(k, m)
	}
	return nil
}
