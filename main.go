package main

import (
	"executerSQL/query"
	"html/template"
)

func main() {
	qb := query.NewQueryCatalog().Escape(template.HTMLEscapeString).Comment("--").LoadFile("test.sql")

	queryExecute := qb.GetSnippet("QueryTest1")
	queryExecute.Param("id", 2599999).Param("height", 24).Get()

	queryExecute = qb.GetSnippet("QueryTest2")
	queryExecute.Param("id", 25.20).Param("lastName", "JC").Get()

	qb = query.NewQueryCatalog().Escape(template.HTMLEscapeString).Comment("%%").LoadFile("test.pl")
	queryExecute = qb.GetSnippet("PrologQueryTest2")
	queryExecute.Param("feature", "Parametro").Get()
}
