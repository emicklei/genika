package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
)

var (
	oSwaggerJsonUrl  = flag.String("url", "http://localhost:8080/apidocs.json", "full URL of the Swagger JSON listing")
	oTargetDirectory = flag.String("o", "/tmp", "directory to generate test files in")
	oTests           = flag.Bool("tests", true, "generate example test functions")
	oApiName         = flag.String("app", "api", "name of the client for the API")
)

func main() {
	flag.Parse()
	if full, err := filepath.Abs(*oTargetDirectory); err != nil {
		log.Fatal(err.Error())
	} else {
		os.Mkdir(full, os.ModeDir|os.ModePerm)
	}
	writeInit()
	writeApiOperations()
	if *oTests {
		writeSampleTests()
	}
}
