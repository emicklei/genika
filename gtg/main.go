package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Gotype struct {
	Package      string
	TypeName     string
	TemplateName string
}

var toolHome = filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "emicklei", "go-templates")

var (
	oBaseDir  = flag.String("base", toolHome, "location of template folders")
	oPackage  = flag.String("pkg", "model", "name of the package")
	oTypename = flag.String("type", "Object", "name of the type")
	oTemplate = flag.String("tmp", "associationlist", "name of the go-templates folder that contains the templates")
	oTarget   = flag.String("out", "/tmp", "target folder to generate Go source files")
)

func main() {
	flag.Parse()
	g := Gotype{
		Package:      *oPackage,
		TypeName:     *oTypename,
		TemplateName: *oTemplate,
	}
	input := filepath.Join(*oBaseDir, *oTemplate, *oTemplate+".go")
	output := filepath.Join(*oTarget, strings.ToLower(g.TypeName)+"_"+*oTemplate+".go")
	if err := process(input, g, output); err != nil {
		log.Fatalf("template generation failed:%s\n%s\n%v\n", input, output, err)
	}
	input = filepath.Join(*oBaseDir, *oTemplate, *oTemplate+"_test.go")
	output = filepath.Join(*oTarget, strings.ToLower(g.TypeName)+"_"+*oTemplate+"_test.go")
	if err := process(input, g, output); err != nil {
		log.Fatalf("template generation failed:%s\n%s\n%v\n", input, output, err)
	}
}

func process(input string, g Gotype, output string) error {
	data, err := ioutil.ReadFile(input)
	if err != nil {
		return err
	}
	t := template.Must(template.New("gen").Parse(string(data)))
	out, err := os.Create(output)
	if err != nil {
		return err
	}
	defer out.Close()
	return t.Execute(out, g)
}
