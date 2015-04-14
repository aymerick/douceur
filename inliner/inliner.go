package inliner

import (
	"errors"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// CSS Inliner
type Inliner struct {
	// Raw HTML
	html string

	// Parsed HTML document
	doc *goquery.Document
}

// Instanciate a new Inliner
func NewInliner(html string) *Inliner {
	return &Inliner{
		html: html,
	}
}

// Inlines css into html document
func Inline(html string) (string, error) {
	result, err := NewInliner(html).Inline()
	if err != nil {
		return "", err
	} else {
		return result, nil
	}
}

// Inlines CSS and returns HTML
func (inliner *Inliner) Inline() (string, error) {
	// parse HTML document
	if err := inliner.parseHTML(); err != nil {
		return "", err
	}

	// @todo Parse styles

	// @todo Collect elements

	// generate HTML document
	result, _ := inliner.genHTML()

	// @todo Finnish that !
	return result, errors.New("NOT IMPLEMENTED")
}

// Parses raw html
func (inliner *Inliner) parseHTML() error {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(inliner.html))
	if err != nil {
		return err
	}

	inliner.doc = doc

	return nil
}

// Generates HTML
func (inliner *Inliner) genHTML() (string, error) {
	return inliner.doc.Html()
}
