package query

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type Query struct {
	Name           string
	Value          string
	FunctionEscape func(string) string
	Parameters     map[string]QueryParam
	err            error
}

type QueryParam struct {
	TypeParam string
	Value     string
}

func (q *Query) Escape(FunctionEscape func(string) string) *Query {
	q.FunctionEscape = FunctionEscape
	return q
}

func (q *Query) Param(nameParam string, valueParam interface{}) *Query {
	if q.err != nil {
		return q
	}

	if valueParam == nil {
		q.err = errors.New("valueParam must not be null")
		return q
	}
	param := q.Parameters[nameParam]
	var typeParam = reflect.TypeOf(valueParam).Name()
	if typeParam != param.TypeParam {
		q.err = errors.New("Type invalid for parameter " + nameParam)
		return q
	}
	param.Value = fmt.Sprint(valueParam)
	if q.FunctionEscape != nil {
		param.Value = q.FunctionEscape(param.Value)
	}
	q.Parameters[nameParam] = param
	return q
}

func (q *Query) formatParameter(value string) string {
	return "${{" + value + "}}"
}

func (q *Query) Get() error {
	if q.err != nil {
		fmt.Println(q.err)
		return q.err
	}
	var queryValue = q.Value
	for key, param := range q.Parameters {
		queryValue = strings.ReplaceAll(queryValue, q.formatParameter(key), param.Value)
	}
	fmt.Println(queryValue)
	return nil
}
