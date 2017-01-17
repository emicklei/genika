# gtg - go template based go source generator

Inspired by 
- https://github.com/ncw/gotemplate
- http://bouk.co/blog/idiomatic-generics-in-go/

### Install

	go install

### Using orderedmap

	//go:generate gtg -pkg=model -type=SecurityScheme -tmp=orderedmap -out=.

### Using optional

	//go:generate gtg -pkg=model -type=Account -tmp=optional -out=.


(c)2016, http://ernestmicklei.com MIT License