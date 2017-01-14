package main

import (
	"strings"
	"time"

	"github.com/emicklei/go-restful/swagger"
)

type HeaderData struct {
	Today  time.Time
	Source string
}

func basePathFrom(url string) string {
	// guess
	j := strings.Index(url, "/apidocs.json")
	return url[:j]
}

func sanitize(resourcepath string) string {
	withoutSlashes := strings.Replace(resourcepath, "/", "_", -1)
	curly := strings.Index(withoutSlashes, "{")
	if curly > -1 {
		return withoutSlashes[:curly]
	}
	return withoutSlashes
}

func hasPathParams(resourcepath string) bool {
	return strings.Index(resourcepath, "{") != -1
}

func status(op swagger.Operation) int {
	if len(op.ResponseMessages) > 0 {
		return op.ResponseMessages[0].Code

	}
	m := op.Method
	if m == "PUT" || m == "DELETE" {
		return 204
	}
	return 200
}
