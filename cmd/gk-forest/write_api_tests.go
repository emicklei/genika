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

func writeSampleTests() {
	full, _ := filepath.Abs(*oTargetDirectory)
	where := path.Join(full, "api_test.go")
	_, err := os.Stat(where)
	if err == nil {
		log.Printf("[forestgen] skipped, not overwriting %s\n", where)
		return
	}
	out, err := os.Create(where)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	header := HeaderData{Today: time.Now(), Source: *oSwaggerJsonUrl}
	testHeader.Execute(out, header)

	listing := getListing()
	for _, each := range listing.Apis {
		decl := getApiDeclaration(each.Path)
		for _, api := range decl.Apis {
			for _, op := range api.Operations {
				writeApiTest(api, op, out)
			}
		}
	}
	log.Printf("[forestgen] written %s\n", where)
}

func writeApiTest(api Api, op Operation, out *os.File) {
	data := testData{
		Nickname:         op.Nickname,
		Description:      html.UnescapeString(op.Summary),
		ParametersValues: parameterValuesFrom(op),
	}
	testTemplate.Execute(out, data)
}

func parameterValuesFrom(op Operation) string {
	buf := new(bytes.Buffer)
	for _, each := range op.Parameters {
		buf.WriteString(", ")
		switch strings.ToLower(*each.Type) {
		case "string":
			fmt.Fprintf(buf, "%q", each.Name)
		case "boolean":
			fmt.Fprint(buf, "false")
		case "integer", "int", "long", "byte":
			fmt.Fprintf(buf, "%d", 42)
		case "timestamp", "time":
			buf.WriteString("time.Now()")
		default:
			log.Printf("[WARN] no default value for type:%s", *each.Type)
			fmt.Fprintf(buf, "%s{}", *each.Type)
		}
	}
	return buf.String()
}
