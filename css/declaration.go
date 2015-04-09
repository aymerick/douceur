package css

import "fmt"

type Declaration struct {
	Property string
	Value    string
}

func NewDeclaration() *Declaration {
	return &Declaration{}
}

// Returns string representation of the Declaration
func (decl *Declaration) String() string {
	return fmt.Sprintf("%s: %s;", decl.Property, decl.Value)
}

// Returns true if both Declarations are equals
func (decl *Declaration) Equal(other *Declaration) bool {
	return (decl.Property == other.Property) && (decl.Value == other.Value)
}
