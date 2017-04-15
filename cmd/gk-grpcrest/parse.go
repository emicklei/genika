package main

import (
	"log"
	"os"

	"github.com/emicklei/proto"
)

func main() {
	r, _ := os.Open(os.Args[1])
	defer r.Close()
	p := proto.NewParser(r)
	def, err := p.Parse()
	if err != nil {
		log.Fatal(err)
	}
	c := Collector{def}
	srv := c.Services()
	log.Println(srv[0].Name)
	log.Println(c.RPCsOf(srv[0])[0].Name)

	routes := []routeData{}
	for _, each := range c.RPCsOf(srv[0]) {
		comment := ""
		if each.Comment != nil {
			comment = each.Comment.Message()
		}
		if len(each.Options) > 0 {
			m, p := httpMethodAndPath(each.Options)
			d := routeData{
				HTTPMethod:        m,
				HTTPPath:          p,
				InputMessage:      each.RequestType,
				OutputMessage:     each.ReturnsType,
				ResourceOperation: each.Name,
				ServiceName:       srv[0].Name,
				Doc:               comment,
			}
			routes = append(routes, d)
		}
	}
	d := resourceData{
		ServiceName: srv[0].Name,
		APIPackage:  "account",
		Routes:      routes,
	}
	if err := resourceTemplate.Execute(os.Stdout, d); err != nil {
		log.Fatalf("writing resource failed:%v", err)
	}

	if err := webserviceTemplate.Execute(os.Stdout, d); err != nil {
		log.Fatalf("writing resource failed:%v", err)
	}

	if err := operationsTemplate.Execute(os.Stdout, d); err != nil {
		log.Fatalf("writing resource failed:%v", err)
	}

}

func httpMethodAndPath(options []*proto.Option) (string, string) {
	for _, each := range options {
		if each.Name == "(google.api.http)" {
			for _, other := range each.AggregatedConstants {
				switch other.Name {
				case "get":
					return "GET", other.Source
				case "put":
					return "PUT", other.Source
				case "post":
					return "POST", other.Source
				case "delete":
					return "DELETE", other.Source
				case "head":
					return "HEAD", other.Source
				case "patch":
					return "PATCH", other.Source
				}
			}
		}
	}
	log.Fatal("no valid google.api.http found")
	return "", ""
}

func optionValueAt(options []*proto.Option, key string) string {
	for _, each := range options {
		for _, other := range each.AggregatedConstants {
			if other.Name == key {
				return other.Source
			}
		}
	}
	log.Fatal("no option found with key:" + key)
	return ""
}
