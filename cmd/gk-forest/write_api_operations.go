package main

import (
	"bytes"
	"fmt"
	"html"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-openapi/spec"
)

func writeApiOperations() {
	full, _ := filepath.Abs(*oTargetDirectory)
	where := path.Join(full, "api_operations.go")
	out, err := os.Create(where)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	header := HeaderData{Today: time.Now(), Source: *oSwaggerJsonUrl}
	apiopHeader.Execute(out, header)

	doc := getDocument()
	for k, item := range doc.Spec().Paths.Paths {
		writeApiOperation(k, item, item.Get, out)
		writeApiOperation(k, item, item.Put, out)
		writeApiOperation(k, item, item.Post, out)
		writeApiOperation(k, item, item.Patch, out)
		writeApiOperation(k, item, item.Head, out)
	}
	log.Printf("[forestgen] written %s\n", where)
}

func writeApiOperation(path string, item spec.PathItem, op *spec.Operation, out *os.File) {
	if op == nil {
		return
	}
	data := apiopData{
		Nickname:              getOperationName(op.ID),
		Description:           fmt.Sprintf("%s\n%s", op.ID, html.UnescapeString(op.Description)),
		ShortDescription:      strings.Replace(strings.Split(html.UnescapeString(op.Summary), "\n")[0], "\"", "'", -1),
		HttpMethod:            getMethod(item),
		Path:                  path,
		PathParameters:        pathParametersFrom(op),
		QueryAndHeaderCalls:   configBuildCallsFrom(op),
		ParametersDeclaration: parameterDeclarationFrom(op),
		ApiName:               *oApiName,
	}
	apiopTemplate.Execute(out, data)
}

func getOperationName(op string) string {
	if strings.Contains(op, "/") {
		return "do" + fmt.Sprintf("%v", &op)
	}
	return op
}

func getMethod(item spec.PathItem) string {
	if item.Get != nil {
		return "GET"
	}
	if item.Put != nil {
		return "PUT"
	}
	if item.Post != nil {
		return "POST"
	}
	if item.Head != nil {
		return "HEAD"
	}
	if item.Patch != nil {
		return "PATCH"
	}
	panic("unknown method")
}

func pathParametersFrom(op *spec.Operation) string {
	buf := new(bytes.Buffer)
	for _, each := range op.Parameters {
		pname := asParameterName(each.Name)
		switch each.Type {
		case "path":
			buf.WriteString(", ")
			buf.WriteString(pname)
		}
	}
	return buf.String()
}

func configBuildCallsFrom(op *spec.Operation) string {
	buf := new(bytes.Buffer)
	for _, each := range op.Parameters {
		pname := asParameterName(each.Name)
		switch each.Type {
		case "query":
			fmt.Fprintf(buf, ".\n\t\tQuery(%q,%s)", each.Name, pname)
		case "header":
			fmt.Fprintf(buf, ".\n\t\tHeader(%q,%s)", each.Name, pname)
		case "body":
			fmt.Fprintf(buf, ".\n\t\tContent(%s,\"application/json\")", pname)
		}
	}
	return buf.String()
}

func parameterDeclarationFrom(op *spec.Operation) string {
	buf := new(bytes.Buffer)
	for _, each := range op.Parameters {
		buf.WriteString(", ")
		fmt.Fprintf(buf, "%s %s", asParameterName(each.Name), asGoDatatype(each.Type))
	}
	return buf.String()
}

func asGoDatatype(t string) string {
	mapping := map[string]string{
		"timestamp": "*time.Time",
		"integer":   "int",
		"boolean":   "bool",
		"long":      "int",
		"string":    "string",
	}
	dt, ok := mapping[strings.ToLower(t)]
	if !ok {
		return "interface{}"
	}
	return dt
}

func asParameterName(n string) string {
	return strings.Replace(n, "-", "_", -1)
}
