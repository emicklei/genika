# gtg - go template based go source generator

### Install

	go install

### Using orderedmap

	//go:generate gtg -pkg=model -type=SecurityScheme -tmp=orderedmap -out=.

### Using optional

	//go:generate gtg -pkg=model -type=Account -tmp=optional -out=.


(c)2016, http://ernestmicklei.com MIT License