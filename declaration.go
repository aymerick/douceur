package douceur

import "fmt"

type Declaration struct {
	Property string
	Value    string
}

func NewDeclaration() *Declaration {
	return &Declaration{}
}

func (decl *Declaration) String() string {
	return fmt.Sprintf("%s: %s;", decl.Property, decl.Value)
}
