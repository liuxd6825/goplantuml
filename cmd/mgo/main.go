package mgo

import (
	"flag"
	"github.com/jfeliu007/goplantuml/pkg/ddd"
)

func main() {
	outputPath := flag.String("output", "", "output file path. If omitted, then this will default to standard output")
	tplPath := flag.String("tpl", "", "output file path. If omitted, then this will default to standard output")
	umlFile := flag.String("uml", "", "output file path. If omitted, then this will default to standard output")

	ddd.Builder(*outputPath, *tplPath, *umlFile)
}
