package {{.Package}}

// generated with github.com/emicklei/genika/gtg/{{.TemplateName}}

var AbsentOptional{{.TypeName}} = Optional{{.TypeName}}{hasValue: false}

// Optional{{.TypeName}} is an immutable value that may contain a non-empty {{.TypeName}}.
type Optional{{.TypeName}} struct {
	value    {{.TypeName}}
	hasValue bool
}

// Returns an Optional{{.TypeName}} instance containing the given non-nil reference
func Optional{{.TypeName}}Of(v {{.TypeName}}) Optional{{.TypeName}} {
	return Optional{{.TypeName}}{value: v, hasValue: true}
}

// Get returns the contained instance, which must be present.
func (o Optional{{.TypeName}}) Get() {{.TypeName}} {
	return o.value
}

// Returns true if this holder contains a (non-nil) instance.
func (o Optional{{.TypeName}}) IsPresent() bool {
	return o.hasValue
}

// Equals returns true if object is an Optional instance, and either the contained references are equal to each other or both are absent.
func (o Optional{{.TypeName}}) Equals(v Optional{{.TypeName}}) bool {
	return o.hasValue == v.hasValue && o.value == v.value
}