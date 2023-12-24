package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/nonsocode/xpress/pkg/parser"
)

func Benchmark_xpress(b *testing.B) {
	params := createParams()

	var err error
	var out interface{}
	ast := parser.NewParser(fmt.Sprintf("@{{%s}}", example)).Parse()
	evaluator := parser.NewInterpreter()
	evaluator.SetMembers(params)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, err = evaluator.Evaluate(context.Background(), ast)
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	if !out.(bool) {
		b.Fail()
	}
}

func Benchmark_xpress_func(b *testing.B) {
	params := createParams()
	params["join"] = func(params ...interface{}) (interface{}, error) {
		return params[0].(string) + params[1].(string), nil
	}
	var err error
	var out interface{}
	ast := parser.NewParser(fmt.Sprintf("@{{%s}}", `join("hello", ", world")`)).Parse()
	evaluator := parser.NewInterpreter()
	evaluator.SetMembers(params)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, err = evaluator.Evaluate(context.TODO(), ast)
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	if out.(string) != "hello, world" {
		b.Fail()
	}
}
