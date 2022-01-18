package query

import (
	"fmt"
	"reflect"
	"strings"
)

type Query struct {
	Name           string
	Value          string
	FunctionEscape func(string) string
	Parameters     map[string]QueryParam
}

type QueryParam struct {
	TypeParam string
	Value     string
}

const NumberType = "number"
const PrefixFloat = "float"
const PrefixInt = "int"
const PrefixUint = "uint"

func (q *Query) Escape(FunctionEscape func(string) string) *Query {
	q.FunctionEscape = FunctionEscape
	return q
}

func (q *Query) Param(nameParam string, valueParam interface{}) *Query {
	if valueParam == nil {
		panic("valueParam must not be null")
	}
	param := q.Parameters[nameParam]
	typeParam := reflect.TypeOf(valueParam).Name()

	if !q.validateTypeValue(param, typeParam) {
		panic("Type invalid for parameter " + nameParam)
	}
	param.Value = fmt.Sprint(valueParam)
	if q.FunctionEscape != nil {
		param.Value = q.FunctionEscape(param.Value)
	}
	q.Parameters[nameParam] = param
	return q
}

func (q *Query) validateTypeValue(param QueryParam, typeParam string) bool {
	if q.validateNumber(param, typeParam) {
		return true
	}
	return typeParam == param.TypeParam
}

func (q *Query) validateNumber(param QueryParam, typeParam string) bool {
	return param.TypeParam == NumberType && (strings.HasPrefix(typeParam, PrefixFloat) || strings.HasPrefix(typeParam, PrefixInt) || strings.HasPrefix(typeParam, PrefixUint))
}

func (q *Query) formatParameter(value string) string {
	return "${{" + value + "}}"
}

func (q *Query) Get() error {

	queryValue := q.Value
	for key, param := range q.Parameters {
		queryValue = strings.ReplaceAll(queryValue, q.formatParameter(key), param.Value)
	}
	fmt.Println(queryValue)
	return nil
}
