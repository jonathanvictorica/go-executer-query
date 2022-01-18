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
const PrefixVarValue = "var_value:"
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
			nameParam := b.getNameParam(bufferString)
			if nameParam == "" {
				continue
			}
			parameters[nameParam] = QueryParam{
				TypeParam: b.getTypeParam(bufferString),
				Value:     b.getValueInitParam(bufferString),
			}
		} else {
			return parameters
		}
	}
	return parameters

}

func (b QueryCatalog) getNameParam(bufferString string) string {
	if !strings.Contains(bufferString, PrefixVar) {
		return ""
	}
	return strings.Split(strings.Split(bufferString, PrefixVar)[1], ",")[0]
}

func (b QueryCatalog) getValueInitParam(bufferString string) string {
	if !strings.Contains(bufferString, PrefixVarValue) {
		return ""
	}
	return strings.Split(strings.Split(bufferString, PrefixVarValue)[1], ",")[0]
}

func (b QueryCatalog) getTypeParam(bufferString string) string {
	if !strings.Contains(bufferString, PrefixVarType) {
		return VarTypeDefault
	}
	var value = strings.Split(strings.Split(bufferString, PrefixVarType)[1], ",")[0]
	if value == "" {
		return VarTypeDefault
	}
	return value
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
		if strings.TrimSpace(bufferString) == "" {
			return queryValue
		}
		queryValue += " " + bufferString
	}
	return queryValue
}
