package tools

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"log"
)

func SearchStructsInPackage(name string) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "./msgpack/v5.go", nil, parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}
	conf := types.Config{
		Importer: importer.Default(),
		Error:    func(err error) {},
	}
	pkg, err := conf.Check(".", fset, []*ast.File{file}, nil)
	names := pkg.Scope().Names()
	for i := 0; i < len(names); i++ {
		o := pkg.Scope().Lookup(names[i])
		t := o.Type().Underlying() //.(*types.Struct)
		fmt.Printf("name: %v , type: %v \n", names[i], t)
		if casted, ok := t.(*types.Struct); ok {
			fmt.Printf("casted type: %v \n", casted)

			j, _ := json.Marshal(casted)
			fmt.Printf("json: %v\n", j)
		}
	}
	//s := pkg.Scope().Names()
	//internal := s.Type().Underlying().(*types.Struct)
}
