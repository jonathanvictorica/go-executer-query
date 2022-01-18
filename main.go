package main

import (
	"executerSQL/query"
	"html/template"
)

func main() {

	qb := query.LoadFile("test.sql")

	var query = qb.GetSnippet("QueryTest1")
	query.Escape(template.HTMLEscapeString).Param("id", 25).Param("height", 1.78).Param("name", "JC").Get()

	query = qb.GetSnippet("QueryTest2")
	query.Escape(template.HTMLEscapeString).Param("id", 25).Param("lastName", "JC").Get()

	query = qb.GetSnippet("PrologQueryTest2")
	query.Escape(template.HTMLEscapeString).Param("feature", "CreationUser").Get()
}
