package douceur

type Stylesheet struct {
	Rules []*Rule
}

func NewStylesheet() *Stylesheet {
	return &Stylesheet{}
}
