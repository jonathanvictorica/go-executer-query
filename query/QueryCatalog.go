package query

import (
	"bufio"
	"log"
	"os"
	"strings"
)

const PrefixNameQuery = "name:"
const VarTypeDefault = "string"

type QueryCatalog struct {
	fileNameQuerys string
	querys         map[string]Query
	functionEscape func(string) string
	comment        string
}

func NewQueryCatalog() *QueryCatalog {
	return &QueryCatalog{}
}

func (q *QueryCatalog) Escape(FunctionEscape func(string) string) *QueryCatalog {
	q.functionEscape = FunctionEscape
	return q
}

func (q *QueryCatalog) Comment(comment string) *QueryCatalog {
	q.comment = comment
	return q
}

func (q *QueryCatalog) LoadFile(nameFile string) *QueryCatalog {
	q.fileNameQuerys = nameFile
	var queryCatalog, err = q.createQuerys(nameFile)
	if err != nil {
		panic("Error loading file initialization")
	}
	q.querys = queryCatalog
	return q
}

func (q *QueryCatalog) GetSnippet(queryName string) Query {
	return q.querys[queryName]
}

func (q *QueryCatalog) createParametersQuery(fileScanner *bufio.Scanner) map[string]QueryParam {
	var parameters = make(map[string]QueryParam)
	for fileScanner.Scan() {
		bufferString := fileScanner.Text()
		if strings.HasPrefix(bufferString, q.comment) {
			param := bufferString[strings.Index(bufferString, ":")+1 : len(bufferString)]
			nameParam := q.getNameParam(param)
			if nameParam == "" {
				continue
			}
			parameters[nameParam] = QueryParam{
				TypeParam: q.getTypeParam(param),
				Value:     q.getValueInitParam(param),
			}
		} else {
			return parameters
		}
	}
	return parameters

}
func (q *QueryCatalog) getNameParam(bufferString string) string {
	indexTypeParam := strings.Index(bufferString, ":")
	if indexTypeParam > 0 {
		return bufferString[0:indexTypeParam]
	}
	indexEqual := strings.Index(bufferString, "=")
	if indexEqual > 0 {
		return bufferString[0:indexEqual]
	}
	return bufferString

}

func (q *QueryCatalog) getValueInitParam(bufferString string) string {
	if !strings.Contains(bufferString, "=") {
		return ""
	}
	return bufferString[strings.Index(bufferString, "=")+1 : len(bufferString)]
}

func (q *QueryCatalog) getTypeParam(bufferString string) string {
	if !strings.Contains(bufferString, ":") {
		return VarTypeDefault
	}
	if !strings.Contains(bufferString, "=") {
		return q.convertionType(bufferString[strings.Index(bufferString, ":")+1 : len(bufferString)])
	}
	return q.convertionType(bufferString[strings.Index(bufferString, ":")+1 : strings.Index(bufferString, "=")])

}

func (q *QueryCatalog) createQuerys(file string) (map[string]Query, error) {
	querys := make(map[string]Query)
	readFile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		bufferString := fileScanner.Text()
		if strings.HasPrefix(bufferString, q.comment) && strings.Contains(bufferString, PrefixNameQuery) {
			nameQuery := strings.TrimSpace(strings.Split(bufferString, PrefixNameQuery)[1])
			parameters := q.createParametersQuery(fileScanner)
			queryValue := q.getQueryValue(fileScanner)
			querys[nameQuery] = Query{
				Name:           nameQuery,
				Value:          queryValue,
				Parameters:     parameters,
				FunctionEscape: q.functionEscape,
			}
		}
	}
	readFile.Close()
	return querys, nil

}

func (q *QueryCatalog) getQueryValue(fileScanner *bufio.Scanner) string {
	queryValue := fileScanner.Text()
	for fileScanner.Scan() {
		bufferString := fileScanner.Text()
		if strings.TrimSpace(bufferString) == "" {
			return queryValue
		}
		queryValue += " " + bufferString
	}
	return queryValue
}

func (q *QueryCatalog) convertionType(typeParam string) string {
	if strings.HasPrefix(typeParam, "float") {
		return "float64"
	}
	return typeParam
}
