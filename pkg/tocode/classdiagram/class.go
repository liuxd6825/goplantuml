package classdiagram

import (
	"context"
	"github.com/jfeliu007/goplantuml/pkg/tocode/classdiagram/utils"
	"strings"
)

type Class struct {
	NameElement
	Fields      []*Field  `json:"fields,omitempty"`
	Methods     []*Method `json:"methods,omitempty"`
	GenericType []string  `json:"genericType,omitempty"` // 是泛型类
	IsInterface bool      `json:"isInterface"`
}

func NewClass(ctx context.Context, line string, namespace string, notes []*Note) *Class {
	class := &Class{
		Fields:  make([]*Field, 0),
		Methods: make([]*Method, 0),
	}
	ns := namespace
	parseRes := ParseClassName(line, "class")
	if parseRes.NamespaceName != "" {
		ns = ns + "." + parseRes.NamespaceName
	}
	class.InitName(parseRes.Name, parseRes.Alias)
	class.InitBase(line, "class", ns, notes)
	class.GenericType = parseRes.GenericType
	return class
}

func NewInterface(ctx context.Context, line string, namespace string, notes []*Note) *Class {
	class := &Class{
		Fields:  make([]*Field, 0),
		Methods: make([]*Method, 0),
	}
	ns := namespace
	parseRes := ParseClassName(line, "interface")
	if parseRes.NamespaceName != "" {
		ns = ns + "." + parseRes.NamespaceName
	}
	class.InitName(parseRes.Name, parseRes.Alias)
	class.InitBase(line, "interface", ns, notes)
	class.GenericType = parseRes.GenericType
	class.IsInterface = true
	return class
}

func (c *Class) GetTypeName() string {
	return "class"
}

func (c *Class) Parse(ctx context.Context, reader ParseReader) error {

	if !strings.HasSuffix(c.Line, "{") {
		return nil
	}
	/*
		{field} A field (despite parentheses)
		{method} Some method
	*/
	var notes []*Note
	for reader.Scan() {
		line := reader.ReadLine()
		if line == "}" {
			return nil
		} else if strings.HasPrefix(line, "==") {
			continue
		} else if strings.HasPrefix(line, "..") {
			continue
		} else if strings.HasPrefix(line, "__") {
			continue
		} else if strings.HasPrefix(line, "--") {
			continue
		} else if strings.HasPrefix(line, "'") {
			notes = append(notes, NewNote(ctx, line))
		} else if utils.StringContains(line, "(", "{method}") {
			c.AddMethod(NewMethod(ctx, c, line, notes).Init(line))
			notes = nil
		} else if len(line) > 0 {
			c.AddField(NewField(ctx, c, line, c.NamespaceName, notes))
			notes = nil
		}
	}
	return nil
}

func (c *Class) AddField(member *Field) {
	c.Fields = append(c.Fields, member)
}

func (c *Class) AddMethod(member *Method) {
	c.Methods = append(c.Methods, member)
}

func (c *Class) LenFields() int {
	return len(c.Fields)
}
func (c *Class) LenMethods() int {
	return len(c.Methods)
}

type ResultParseClassName struct {
	Name          string   `json:"name"`
	Alias         string   `json:"alias"`
	GenericType   []string `json:"genericType"`
	NamespaceName string   `json:"namespaceName"`
}

func ParseClassName(line string, tagName string) *ResultParseClassName {
	name, alias, namespaceName, list := utils.ParseName(line, tagName)
	res := &ResultParseClassName{
		Name:          name,
		Alias:         alias,
		NamespaceName: namespaceName,
	}

	isAdd := false

	for _, item := range list {
		// 解析 <<User>>
		if strings.HasPrefix(item, "<<") && strings.HasSuffix(item, ">>") {
			str := strings.Trim(item[2:len(item)-2], " ")
			res.GenericType = strings.Split(str, " ")
			break
		}

		// 解析 << User >>
		if strings.HasPrefix(item, "<<") {
			isAdd = true
			res.AddGenericType(item[2:])
		} else if strings.HasSuffix(item, ">>") {
			isAdd = false
			res.AddGenericType(item[:len(item)-2])
		} else if isAdd {
			res.GenericType = append(res.GenericType, item)
		}
	}
	return res
}

func (c *ResultParseClassName) AddGenericType(typeName string) {
	s := strings.Trim(typeName, " ")
	if s != "" && s != "," && s != ";" && s != "extends" {
		c.GenericType = append(c.GenericType, s)
	}
}
