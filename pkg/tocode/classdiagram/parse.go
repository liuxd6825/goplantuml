package classdiagram

import (
	"bufio"
	"context"
	"github.com/jfeliu007/goplantuml/pkg/tocode/classdiagram/utils"
	"os"
	"path/filepath"
	"strings"
)

type ParseReader interface {
	Scan() bool
	ReadLine() string
}

type parseReader struct {
	scanner *bufio.Scanner
}

func (p *parseReader) Scan() bool {
	return p.scanner.Scan()
}

func (p *parseReader) ReadLine() string {
	txt := strings.Trim(p.scanner.Text(), " ")
	if utils.IsHtml(txt) {
		txt, _ = utils.HtmlToText(txt)
	}
	return utils.StringClear(txt, "\n", "\r")
}

func NewParseReader(filename string) (ParseReader, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	return &parseReader{
		scanner: scanner,
	}, nil
}

func ParseFile(ctx context.Context, fileName string) (*UMLBlock, error) {
	filename, err := filepath.Abs(fileName)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func() { _ = file.Close() }()

	reader, err := NewParseReader(filename)
	if err != nil {
		return nil, err
	}

	block := NewUMLBlock(filename)
	err = block.Parse(ctx, reader)
	return block, err
}
