package cmd

import (
	"context"
	"github.com/jfeliu007/goplantuml/pkg/tocode/classdiagram"
	"github.com/jfeliu007/goplantuml/pkg/tocode/classdiagram/tags"
	"text/template"
)

func Builder(ctx context.Context, outputPath string, tplPath string, block *classdiagram.UMLBlock) error {
	return build(ctx, outputPath, tplPath, block.Root.Namespaces)
}

func build(ctx context.Context, outputPath string, tplPath string, namespaces map[string]*classdiagram.Namespace) error {
	for _, namespace := range namespaces {
		for _, e := range namespace.Elements {
			if e.GetTypeName() == "class" {
				class := e.(*classdiagram.Class)
				if tags := class.FindTags(tags.TagTypeAggregate); len(tags) > 0 {
					aggregate := NewAggregate(namespace, class)
					buildCode(ctx, aggregate, outputPath, tplPath, "aggregate.tpl", "aggregate_cmd.tpl", "aggregate_event.tpl")
				}
			} else if e.GetTypeName() == "enum" {

			}
		}
		if err := build(ctx, namespace.GetPathName(outputPath), tplPath, namespace.Namespaces); err != nil {
			return err
		}
	}
	return nil
}

func buildCode(ctx context.Context, element any, outputFile string, tplPath string, templateFiles ...string) {
	for _, tplFile := range templateFiles {
		tFile := tplPath + "/" + tplFile
		tpl := template.New("")
		tpl.ParseFiles(tFile)
	}
}
