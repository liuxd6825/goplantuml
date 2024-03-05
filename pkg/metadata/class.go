package metadata

import "github.com/jfeliu007/goplantuml/pkg/tocode/classdiagram"

type Class struct {
	Titles Titles                `json:"titles"`
	Name   string                `json:"name"`
	Fields []*classdiagram.Field `json:"fields"`
	Grids  []*Grid               `json:"grids"`
	Forms  []*Form               `json:"forms"`
}

type Field struct {
	DataType DataType `json:"dataType"`
	Size     int      `json:"size"`
	Name     string   `json:"name"`
	Titles   Titles   `json:"titles"`
}

type Titles struct {
	CN string `json:"cn"`
	EN string `json:"en"`
	DE string `json:"de"`
}

type DataType string

const (
	DataString   DataType = "string"
	DataInt      DataType = "int"
	DataMoney    DataType = "money"
	DataFloat    DataType = "float"
	DataDate     DataType = "date"
	DataTime     DataType = "time"
	DataDatetime DataType = "datetime"
	DataBool     DataType = "bool"
	DataEnum     DataType = "enum"
	DataObject   DataType = "object"
	DataYear     DataType = "year"
	DataMonth    DataType = "month"
	DataDay      DataType = "day"
)
