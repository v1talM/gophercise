package main

import (
	"fmt"
	"gophercise/linkParser"
	"strings"
)

var exampleHtml = `
<html>
	<head>
		<title>example</title>
	</head>
	<body>
		<a href="/page-1">
			text is 
<span>Hello World!</span>
			wow!!
		</a>
		
	</body>
</html>
`

func main() {
	doc := strings.NewReader(exampleHtml)

	links, err := linkParser.Parse(doc)
	if err != nil {
		panic(err)
	}
	for _, link := range links {
		fmt.Printf("%+v\n", link)
	}
}