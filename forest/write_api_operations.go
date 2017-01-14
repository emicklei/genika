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

	listing := getListing()
	for _, each := range listing.Apis {
		decl := getApiDeclaration(each.Path)
		for _, api := range decl.Apis {
			for _, op := range api.Operations {
				writeApiOperation(api, op, out)
			}
		}
	}
	log.Printf("[forestgen] written %s\n", where)
}

func writeApiOperation(api Api, op Operation, out *os.File) {
	data := apiopData{
		Nickname:              op.Nickname,
		Description:           html.UnescapeString(op.Summary),
		ShortDescription:      strings.Replace(strings.Split(html.UnescapeString(op.Summary), "\n")[0], "\"", "'", -1),
		HttpMethod:            op.GetMethod(),
		Path:                  api.Path,
		PathParameters:        pathParametersFrom(op),
		QueryAndHeaderCalls:   configBuildCallsFrom(op),
		ParametersDeclaration: parameterDeclarationFrom(op),
		ApiName:               *oApiName,
	}
	apiopTemplate.Execute(out, data)
}

func pathParametersFrom(op Operation) string {
	buf := new(bytes.Buffer)
	for _, each := range op.Parameters {
		pname := asParameterName(each.Name)
		switch each.ParamType {
		case "path":
			buf.WriteString(", ")
			buf.WriteString(pname)
		}
	}
	return buf.String()
}

func configBuildCallsFrom(op Operation) string {
	buf := new(bytes.Buffer)
	for _, each := range op.Parameters {
		pname := asParameterName(each.Name)
		switch each.ParamType {
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

func parameterDeclarationFrom(op Operation) string {
	buf := new(bytes.Buffer)
	for _, each := range op.Parameters {
		buf.WriteString(", ")
		fmt.Fprintf(buf, "%s %s", asParameterName(each.Name), asGoDatatype(*each.Type))
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
