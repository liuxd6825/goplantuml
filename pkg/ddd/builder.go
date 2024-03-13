package ddd

import (
	"context"
	"github.com/jfeliu007/goplantuml/pkg/ddd/domain/cmd"
	"github.com/jfeliu007/goplantuml/pkg/tocode/classdiagram"
)

func Builder(outputPath string, tplPath string, umlFile string) error {
	ctx := context.Background()
	block, err := classdiagram.ParseFile(ctx, umlFile)
	if err != nil {
		return err
	}
	if err := cmd.Builder(ctx, outputPath, tplPath, block); err != nil {
		return err
	}
	return nil
}
