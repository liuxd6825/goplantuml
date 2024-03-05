package classdiagram

import "github.com/jfeliu007/goplantuml/pkg/tocode/classdiagram/utils"

type MethodResult struct {
	method   *Method
	Name     string `json:"name,omitempty"`
	DataType string `json:"dataType,omitempty"`
}

func NewMethodResult(method *Method) *MethodResult {
	return &MethodResult{method: method}
}

func newMethodResult(f *Method, str string) *MethodResult {
	list := utils.StringSplit(str, " ")
	count := len(list)
	rVal := NewMethodResult(f)
	if count == 1 {
		rVal.Name = ""
		rVal.DataType = utils.StringTrim(list[0])
	} else if count > 1 {
		rVal.Name = utils.StringTrim(list[0])
		rVal.DataType = utils.StringTrim(list[1])
	}
	return rVal
}
