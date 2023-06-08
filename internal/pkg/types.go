package pkg

import "go/ast"

type Package struct {
	Name    string    `json:"name"`
	Fns     []Fn      `json:"fns"`
	Body    string    `json:"body"`
	Ast     *ast.File `json:"-"`
	Imports []string  `json:"imports"`
}

type Fn struct {
	Name    string   `json:"name"`
	Body    string   `json:"body"`
	LPos    int      `json:"lpos"`
	RPos    int      `json:"rpos"`
	Imports []string `json:"imports"`
}
