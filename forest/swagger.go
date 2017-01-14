package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/emicklei/go-restful/swagger"
)

func getListing() *swagger.ResourceListing {
	r, err := http.Get(*oSwaggerJsonUrl)
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Fatalf("Getting swagger listing failed: unable to read response body:%v", err)
	}
	var listing swagger.ResourceListing
	err = json.Unmarshal(data, &listing)
	if err != nil {
		log.Fatalf("Parsing swagger listing failed:%v", err)
	}
	return &listing
}

func getApiDeclaration(path string) *ApiDeclaration {
	r, err := http.Get(*oSwaggerJsonUrl + path)
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Fatalf("Getting swagger api failed: unable to read response body:%v", err)
	}
	var api ApiDeclaration
	err = json.Unmarshal(data, &api)
	if err != nil {
		log.Fatalf("Parsing swagger api failed:%v", err)
	}
	return &api
}

type ApiDeclaration struct {
	SwaggerVersion string                  `json:"swaggerVersion"`
	ApiVersion     string                  `json:"apiVersion"`
	BasePath       string                  `json:"basePath"`
	ResourcePath   string                  `json:"resourcePath"` // must start with /
	Apis           []Api                   `json:"apis,omitempty"`
	Models         swagger.ModelList       `json:"models,omitempty"`
	Produces       []string                `json:"produces,omitempty"`
	Consumes       []string                `json:"consumes,omitempty"`
	Authorizations []swagger.Authorization `json:"authorizations,omitempty"`
}

type Api struct {
	Path        string      `json:"path"` // relative or absolute, must start with /
	Description string      `json:"description"`
	Operations  []Operation `json:"operations,omitempty"`
}

type Operation struct {
	swagger.DataTypeFields
	Method           string                    `json:"method"`
	HttpMethod       string                    `json:"httpMethod"` // old Swagger API
	Summary          string                    `json:"summary,omitempty"`
	Notes            string                    `json:"notes,omitempty"`
	Nickname         string                    `json:"nickname"`
	Authorizations   []swagger.Authorization   `json:"authorizations,omitempty"`
	Parameters       []swagger.Parameter       `json:"parameters"`
	ResponseMessages []swagger.ResponseMessage `json:"responseMessages,omitempty"` // optional
	Produces         []string                  `json:"produces,omitempty"`
	Consumes         []string                  `json:"consumes,omitempty"`
	Deprecated       string                    `json:"deprecated,omitempty"`
}

func (o Operation) GetMethod() string {
	if len(o.Method) == 0 {
		return o.HttpMethod
	}
	return o.Method
}
