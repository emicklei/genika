package {{.Package}}

import "testing"

func Test{{.TypeName}}Map(t *testing.T) {
	l := &{{.TypeName}}Map{}
	e := {{.TypeName}}{}
	l.Put("a {{.TypeName}}", e)
	data, err := l.MarshalJSON()
	if err != nil {
		t.Fatal("marshal failed", err)
	}
	back := new({{.TypeName}}Map)
	err = back.UnmarshalJSON(data)
	if err != nil {
		t.Fatal("unmarshal failed", err)
	}
	if got, want := len(back.List), 1; got != want {
		t.Errorf("got %v want %v", got, want)
	}
	t.Log(string(data))
}
