package classdiagram

import (
	"context"
	"strings"
)

type Note struct {
	Text string `json:"text"`
}

func NewNote(ctx context.Context, line string) *Note {
	text := strings.Trim(line[1:], " ")
	return &Note{
		Text: text,
	}
}
func (n *Note) GetType() string {
	return "note"
}
func (n *Note) ParseStart() bool {
	return false
}
func (n *Note) ParseEnd() bool {
	return false
}
func (n *Note) StartTag() string {
	return "'"
}
func (n *Note) EndTag() string {
	return ""
}
func (n *Note) AddLine(line string) {

}

func (n *Note) GetText() string {
	return n.Text
}
