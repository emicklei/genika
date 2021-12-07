### rest api test file generator
It targets a Swagger 2.0 API JSON endpoints

	gk-forest -url "http://<yourhost>/apidocs.json" -o .
	
## Petstore example

	mkdir tmp
	gk-forest -url "https://raw.githubusercontent.com/OAI/OpenAPI-Specification/main/examples/v2.0/json/petstore.json" -o tmp

(c) 2015,2016, http://ernestmicklei.com. MIT License