package classdiagram

import (
	"context"
	"github.com/jfeliu007/goplantuml/pkg/tocode/classdiagram/tags"
	"github.com/jfeliu007/goplantuml/pkg/tocode/classdiagram/utils"
	"strings"
)

type Element interface {
	GetLine() string
	GetTypeName() string
	GetNotes() []*Note
	GetNamespaceName() string
	GetName() string
	GetPackage() string
	GetTags() []tags.Tag
}

type BaseElement struct {
	Line          string     `json:"-"`
	TypeName      string     `json:"typeName,omitempty"`
	Notes         []*Note    `json:"notes,omitempty"`
	Tags          []tags.Tag `json:"tags,omitempty"`
	NamespaceName string     `json:"namespaceName,omitempty"`
	Package       string     `json:"package,omitempty"`
}

type NameElement struct {
	BaseElement
	Name  string `json:"name,omitempty"`
	Alias string `json:"alias,omitempty"`
}

func (n *BaseElement) InitBase(line string, typeName string, namespaceName string, notes []*Note) (err error) {
	list := strings.Split(namespaceName, ".")
	n.Package = list[len(list)-1]
	n.Line = line
	n.TypeName = typeName
	n.Notes = notes
	n.NamespaceName = namespaceName
	n.Tags, err = tags.ParseNotes(notesToStrList(notes))
	return
}

func (n *BaseElement) FindTags(tagType string) []tags.Tag {
	var tags []tags.Tag
	for _, tag := range n.Tags {
		if tagType == string(tag.GetTagType()) {
			tags = append(tags, tag)
		}
	}
	return tags
}

func (n *BaseElement) GetPackage() string {
	return n.Package
}

func (n *BaseElement) GetLine() string {
	return n.Line
}

func (n *BaseElement) SetLine(val string) {
	n.Line = val
}

func (n *BaseElement) GetTypeName() string {
	return n.TypeName
}

func (n *BaseElement) SetTypeName(val string) {
	n.TypeName = val
}

func (n *BaseElement) GetNamespaceName() string {
	return n.NamespaceName
}

func (n *BaseElement) SetNamespaceName(val string) {
	n.NamespaceName = val
}

func (n *BaseElement) GetNotes() []*Note {
	return n.Notes
}

func (n *BaseElement) SetNotes(val []*Note) {
	n.Notes = val
}

func (n *BaseElement) GetTags() []tags.Tag {
	return n.Tags
}

func (n *BaseElement) SetTags(val []tags.Tag) {
	n.Tags = val
}

func (n *BaseElement) NotesText() string {
	return GetNotesText(n.Notes)
}

func (n *BaseElement) Register(ctx context.Context, node Element) {
	pctx := GetParseContext(ctx)
	if pctx != nil {
		pctx.AddNode(node)
	}
}

func (n *NameElement) GetName() string {
	return n.Name
}

func (n *NameElement) FUpperName() string {
	return utils.FirstUpperName(n.Name)
}

func (n *NameElement) FLowerName() string {
	return utils.FirstLowerName(n.Name)
}

func (n *NameElement) SetName(val string) {
	n.Name = val
}

func (n *NameElement) GetAlias() string {
	return n.Alias
}

func (n *NameElement) SetAlias(val string) {
	n.Alias = val
}

func (n *NameElement) SN() string {
	if len(n.Name) > 0 {
		return strings.ToLower(n.Name[0:1])
	}
	return "n"
}

func (n *NameElement) InitName(name string, alias string) {
	n.Name = name
	n.Alias = alias
}

func (n *NameElement) Init(line string, typeName string, namespace string, notes []*Note, name string, alias string) {
	n.BaseElement.InitBase(line, typeName, namespace, notes)
	n.InitName(name, alias)
}

func notesToStrList(notes []*Note) []string {
	var list []string
	for _, note := range notes {
		list = append(list, note.Text)
	}
	return list
}
