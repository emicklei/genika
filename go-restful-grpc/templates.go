package main

import "html/template"

type resourceData struct {
	Doc         string
	APIPackage  string
	ServiceName string
	Routes      []routeData
}

type routeData struct {
	ServiceName       string
	Doc               string
	HTTPMethod        string
	HTTPPath          string
	ResourceOperation string
	InputMessage      string
	OutputMessage     string
}

var resourceTemplate = template.Must(template.New("resourceTemplate").Parse(`
package rest

import (
	"context"
	"net/http"

	restful "github.com/emicklei/go-restful"
	api "{{.APIPackage}}"
)

// {{.ServiceName}}Resource implements REST operations for {{.ServiceName}}.
//{{.Doc}}
type {{.ServiceName}}Resource struct {
	Server api.{{.ServiceName}}Server
}
`))

var webserviceTemplate = template.Must(template.New("webserviceTemplate").Parse(`
// WebService returns a new restful.WebService with REST routes for {{.ServiceName}}.
func (r {{.ServiceName}}Resource) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Consumes(restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON)
	{{range .Routes}}
	ws.Route(ws.{{.HTTPMethod}}("{{.HTTPPath}}").To(r.{{.ResourceOperation}}).
		Doc("{{.Doc}}").
		Returns(200,"OK",api.{{.OutputMessage}}{}).
		DefaultReturns(500,"internal server error", restful.ServiceError{}).
		Reads(api.{{.InputMessage}}{}).
		Writes(api.{{.OutputMessage}}}{}))
	{{end}}
	return ws
}
`))

var operationsTemplate = template.Must(template.New("operationsTemplate").Parse(`
{{range .Routes}}
// {{.ResourceOperation}} is dispatched from {{.HTTPMethod}} {{.HTTPPath}}. 
//{{.Doc}}
func (r {{.ServiceName}}Resource) {{.ResourceOperation}}(request *restful.Request, response *restful.Response) {
	in := new(api.{{.InputMessage}})
	if err := request.ReadEntity(in); err != nil {
		out := restful.ServiceError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
		response.WriteHeaderAndEntity(http.StatusBadRequest, out)
		return
	}
	if out, err := r.Server.{{.ResourceOperation}}(request.Request.Context(), in); err != nil {
		response.WriteHeaderAndEntity(http.StatusInternalServerError, err)
		return
	} 
	response.WriteAsJson(out)	
}
{{end}}
`))
