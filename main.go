package main

import (
	"executerSQL/query"
	"html/template"
)

func main() {

	qb := query.LoadFile("test.sql")

	var query = qb.GetSnippet("HolaJony")
	query.Escape(template.HTMLEscapeString).Param("id", 25).Param("height", 1.78).Param("name", "JC").Get()

	query = qb.GetSnippet("HolaClaudio")
	query.Escape(template.HTMLEscapeString).Param("id", 25).Param("lastName", "JC").Get()
}
