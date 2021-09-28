package tools

import (
	"fmt"
	"go/ast"
	"go/importer"
	"unicode"

	"golang.org/x/tools/go/packages"
)

func PrintStructsInPackage(name string) {
	pkg, err := importer.Default().Import(name)
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return
	}
	for _, declName := range pkg.Scope().Names() {
		fmt.Println(declName)
	}
}

func PrintTypesInPackage(name string) {
	config := &packages.Config{
		Mode: packages.NeedSyntax,
	}
	pkgs, _ := packages.Load(config, name)
	pkg := pkgs[0]

	for _, s := range pkg.Syntax {
		for n, o := range s.Scope.Objects {
			if o.Kind == ast.Typ {
				// check if type is exported(only need for non-local types)
				if unicode.IsUpper([]rune(n)[0]) {
					// note that reflect.ValueOf(*new(%s)) won't work with interfaces
					fmt.Printf("ProcessType(new(%v.%s)),\n", name, n)
				}
			}
		}
	}
}
