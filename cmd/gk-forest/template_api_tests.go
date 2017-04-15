package main

import "text/template"

type testData struct {
	HeaderData
	Description      string
	Nickname         string
	ParametersValues string
}

var testHeader = template.Must(template.New("testHeader").Parse(`
// This file contains examples of Test functions that use the generated functions per operation.
// It will NOT be overwritten ; you can add, change and remove whatever you like.
//
// Code generated by forestgen <https://github.com/emicklei/forestgen> on {{.Today}} from {{.Source}}

package main

import (
	"net/http"
	"testing"

	. "github.com/emicklei/forest"
)`))

var testTemplate = template.Must(template.New("testTemplate").Parse(`
/* 
{{.Description}}
*/
func Test_{{.Nickname}}(t *testing.T) {
	r := {{.Nickname}}(t{{.ParametersValues}})
	ExpectStatus(t, r, http.StatusOK)
}
`))