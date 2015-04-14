package css

import "fmt"

// A parsed style property
type Declaration struct {
	Property  string
	Value     string
	Important bool
}

// Instanciate a new Declaration
func NewDeclaration() *Declaration {
	return &Declaration{}
}

// Returns string representation of the Declaration
func (decl *Declaration) String() string {
	result := fmt.Sprintf("%s: %s", decl.Property, decl.Value)

	if decl.Important {
		result += " !important"
	}

	result += ";"

	return result
}

// Returns true if both Declarations are equals
func (decl *Declaration) Equal(other *Declaration) bool {
	return (decl.Property == other.Property) && (decl.Value == other.Value) && (decl.Important == other.Important)
}
