package {{.Package}}

// generated with github.com/emicklei/genika/gtg/{{.TemplateName}}

import "testing"

func TestOptional{{.TypeName}}(t *testing.T) {
	v := Optional{{.TypeName}}Of({{.TypeName}}{})
	t.Log(v.IsPresent(), v.Get())
}