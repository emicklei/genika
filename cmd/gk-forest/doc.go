package main

/*
Command forestgen can generate Go source files to create an testsuite for REST apis using the forest package.
It targets endpoints that can provide a Swagger v1 JSON file.

	forest -url "http://<yourhost>/apidocs.json" -o .

(c) 2015, http://ernestmicklei.com. MIT License
*/
