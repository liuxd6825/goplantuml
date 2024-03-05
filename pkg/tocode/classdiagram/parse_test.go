package classdiagram

import (
	"context"
	"encoding/json"
	"github.com/jfeliu007/goplantuml/pkg/tocode/classdiagram/utils"
	"testing"
)

func TestParseFile(t *testing.T) {
	ctx := context.Background()
	if black, err := ParseFile(ctx, "./test/example/tag.puml"); err != nil {
		t.Error(err)
	} else {
		s, err := json.Marshal(black)
		if err != nil {
			t.Error(err)
			return
		}
		println(string(s))
	}
}

func TestMakeCode(t *testing.T) {
	ctx := context.Background()
	err := MakeCode(ctx, "./test/example/go.puml", "./template/golang", "./source", "go", option)
	if err != nil {
		t.Error(err)
	}
}

func TestMakeCode_MethodHtml(t *testing.T) {
	ctx := context.Background()
	err := MakeCode(ctx, "./test/example/method_html.puml", "./template/golang", "./source", "go", option)
	if err != nil {
		t.Error(err)
	}
}

func option(opt *MackCodeOption) {
	opt.SetFileName = func(name string) string {
		return utils.SnakeString(name)
	}
}
