package main

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"text/template"
)

func writeInit() {
	full, _ := filepath.Abs(*oTargetDirectory)
	where := path.Join(full, "init_test.go")
	_, err := os.Stat(where)
	if err == nil {
		log.Printf("[forestgen] skipped, not overwriting %s\n", where)
		return
	}
	s, err := os.Create(where)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()
	setup.Execute(s, *oApiName)
	log.Printf("[forestgen] written %s\n", where)
}

var setup = template.Must(template.New("init").Parse(`package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/emicklei/forest"
)

// go test -v -url=http://localhost:8080

// flag variables
var url = flag.String("url", "http://localhost:8080", "the endpoint url against which the tests are run")

// the {{.}}
var {{.}} *forest.APITesting
var _ = fmt.Println

func init() {
	flag.Parse()
	{{.}} = forest.NewClient(*url, new(http.Client))
}

func DefaultConfig(pathTemplate string, pathParams ...interface{}) *forest.RequestConfig {
	cfg := forest.NewConfig(pathTemplate, pathParams...)
	// modify the config for your tests
	return cfg
}
`))
