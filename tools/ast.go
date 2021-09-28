package tools

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

func Search(folder, outputFile, pkgName, structName, ifName, outputTemplate string) {
	fset := token.NewFileSet()

	pkgs, err := parser.ParseDir(fset, folder, nil, parser.AllErrors)
	if err != nil {
		log.Fatalf("Unable to parse %s folder", folder)
	}
	var appPkg *ast.Package
	for _, pkg := range pkgs {
		if pkg.Name == pkgName {
			appPkg = pkg
			break
		}
	}
	if appPkg == nil {
		log.Fatalf("Unable to find package %s", pkgName)
	}

	funcs := make([]string, 0)
	for _, file := range appPkg.Files {
		log.Printf("parsing %s\n", fset.File(file.Pos()).Name())
		if fset.File(file.Pos()).Name() == outputFile {
			continue
		}
		ast.Inspect(file, func(n ast.Node) bool {
			if fun, ok := n.(*ast.FuncDecl); ok {
				if fun.Recv != nil {
					if fun.Name.IsExported() {
						if fun.Recv != nil && len(fun.Recv.List) == 1 {
							if r, rok := fun.Recv.List[0].Type.(*ast.StarExpr); rok && r.X.(*ast.Ident).Name == structName {
								funcs = append(funcs, functionDef(fun, fset))
							}
						}
					}
				}

			}
			return true
		})
	}
}
