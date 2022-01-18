package query

import (
	"bufio"
	"log"
	"os"
	"strings"
)

const PrefixComment = "--"
const PrefixNameQuery = "name:"
const PrefixVar = "var:"
const PrefixVarType = "var_type:"
const VarTypeDefault = "string"

type QueryCatalog struct {
	fileNameQuerys string
	querys         map[string]Query
}

func LoadFile(nameFile string) *QueryCatalog {
	var queryBase = &QueryCatalog{fileNameQuerys: nameFile}
	var queryCatalog, err = queryBase.createQuerys(nameFile)
	if err != nil {
		panic("Error loading file initialization")
	}
	queryBase.querys = queryCatalog
	return queryBase
}

func (b QueryCatalog) GetSnippet(queryName string) Query {
	return b.querys[queryName]
}

func (b QueryCatalog) createParametersQuery(fileScanner *bufio.Scanner) map[string]QueryParam {
	var parameters = make(map[string]QueryParam)
	for fileScanner.Scan() {
		bufferString := fileScanner.Text()
		if strings.HasPrefix(bufferString, PrefixComment) {
			parameterFile := strings.Split(bufferString, ",")
			nameParam := strings.TrimSpace(strings.Split(parameterFile[0], PrefixVar)[1])
			parameters[nameParam] = QueryParam{
				TypeParam: b.getTypeParam(parameterFile),
			}
		} else {
			return parameters
		}
	}
	return parameters

}

func (b QueryCatalog) getTypeParam(parameterFile []string) string {
	typeParam := VarTypeDefault
	if len(parameterFile) > 1 {
		typeParam = strings.TrimSpace(strings.Split(parameterFile[1], PrefixVarType)[1])
	}
	return typeParam
}

func (b QueryCatalog) createQuerys(file string) (map[string]Query, error) {
	var querys = make(map[string]Query)
	readFile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		bufferString := fileScanner.Text()
		if strings.HasPrefix(bufferString, PrefixComment) && strings.Contains(bufferString, PrefixNameQuery) {
			var nameQuery = strings.TrimSpace(strings.Split(bufferString, PrefixNameQuery)[1])
			var parameters map[string]QueryParam
			parameters = b.createParametersQuery(fileScanner)
			queryValue := b.getQueryValue(fileScanner)
			querys[nameQuery] = Query{
				Name:       nameQuery,
				Value:      queryValue,
				Parameters: parameters,
			}
		}
	}
	readFile.Close()
	return querys, nil

}

func (b QueryCatalog) getQueryValue(fileScanner *bufio.Scanner) string {
	queryValue := fileScanner.Text()
	for fileScanner.Scan() {
		bufferString := fileScanner.Text()
		if strings.TrimSpace(bufferString) != "" {
			queryValue = queryValue + " " + bufferString
		} else {
			return queryValue
		}
	}
	return queryValue
}
