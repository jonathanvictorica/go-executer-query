package main

import (
	"executerSQL/bloom"
	"fmt"
)

func main() {
	//qb := query.NewQueryCatalog().Escape(template.HTMLEscapeString).Comment("--").LoadFile("test.sql")
	//
	//queryExecute := qb.GetSnippet("QueryTest1")
	//queryExecute.Param("id", 2599999).Param("height", 24).Get()
	//
	//queryExecute = qb.GetSnippet("QueryTest2")
	//queryExecute.Param("id", 25.20).Param("lastName", "JC").Get()
	//
	//qb = query.NewQueryCatalog().Escape(template.HTMLEscapeString).Comment("%%").LoadFile("test.pl")
	//queryExecute = qb.GetSnippet("PrologQueryTest2")
	//queryExecute.Param("feature", "Parametro").Get()

	consist := bloom.NewConcise()

	consist.Add(64)
	consist.Add(0)
	consist.Add(64)
	consist.Add(12)
	consist.Add(13)

	isContains := consist.Contains(64)
	fmt.Println(isContains)

	isContains = consist.Contains(0)
	fmt.Println(isContains)

	isContains = consist.Contains(128)
	fmt.Println(isContains)
}
