package inliner

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/aymerick/douceur/parser"
)

// An HTML element with matching CSS rules
type Element struct {
	// The goquery handler
	elt *goquery.Selection

	// The style rules to apply on that element
	styleRules []*StyleRule
}

// Instanciate a new element
func NewElement(elt *goquery.Selection) *Element {
	return &Element{
		elt: elt,
	}
}

// Add a Style Rule to Element
func (element *Element) addStyleRule(styleRule *StyleRule) {
	element.styleRules = append(element.styleRules, styleRule)
}

// Parse inline style rules
func (element *Element) parseInlineStyle() ([]*StyleRule, error) {
	result := []*StyleRule{}

	styleValue, exists := element.elt.Attr("style")
	if (styleValue == "") || !exists {
		return result, nil
	}

	declarations, err := parser.ParseDeclarations(styleValue)
	if err != nil {
		return result, err
	}

	result = append(result, NewStyleRule(INLINE_FAKE_SELECTOR, declarations))

	return result, nil
}

// Compute style attribute value
func (element *Element) computesStyle() string {
	// @todo If declaration is !important, it overwrites specificity

	// @todo !!
	return "caca: prout;"
}
