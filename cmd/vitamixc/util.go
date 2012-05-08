// Copyright 2012 Petar Maymounkov. All rights reserved.
// Use of this source code is governed by a 
// license that can be found in the LICENSE file.

package main

import (
	"go/ast"
)

func makeSimpleCallStmt(pkgAlias, funcName string) ast.Stmt {
	return &ast.ExprStmt{
		X: &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X:   &ast.Ident{ Name: pkgAlias },
				Sel: &ast.Ident{ Name: funcName },
			},
		},
	}
}

type VisitorNoReturnFunc func(ast.Node)

func (v VisitorNoReturnFunc) Visit(x ast.Node) ast.Visitor {
	v(x)
	return v
}