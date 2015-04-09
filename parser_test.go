package douceur

import (
	"fmt"
	"strings"
	"testing"
)

func MustParse(t *testing.T, css string, nbRules int) *Stylesheet {
	stylesheet, err := Parse(css)
	if err != nil {
		t.Fatal("Failed to parse css", err, css)
	}

	if len(stylesheet.Rules) != nbRules {
		t.Fatal("Failed to parse Qualified Rules", css)
	}

	return stylesheet
}

func MustEqualRule(t *testing.T, parsedRule *Rule, expectedRule *Rule) {
	if !parsedRule.Equal(expectedRule) {
		diff := parsedRule.Diff(expectedRule)

		t.Fatal(fmt.Sprintf("Rule parsing error\nExpected:\n\"%s\"\nGot:\n\"%s\"\nDiff:\n%s", expectedRule, parsedRule, strings.Join(diff, "\n")))
	}
}

func MustEqualCSS(t *testing.T, ruleString string, expected string) {
	if ruleString != expected {
		t.Fatal(fmt.Sprintf("CSS generation error\n   Expected:\n\"%s\"\n   Got:\n\"%s\"", expected, ruleString))
	}
}

func TestQualifiedRule(t *testing.T) {
	css := `/* This is a comment */
p > a {
    color: blue;
    text-decoration: underline; /* This is a comment */
}`

	expectedRule := &Rule{
		Kind:    QUALIFIED_RULE,
		Prelude: "p > a",
		Declarations: []*Declaration{
			&Declaration{
				Property: "color",
				Value:    "blue",
			},
			&Declaration{
				Property: "text-decoration",
				Value:    "underline",
			},
		},
	}

	expectedCSS := `p > a {
  color: blue;
  text-decoration: underline;
}`

	stylesheet := MustParse(t, css, 1)
	rule := stylesheet.Rules[0]

	MustEqualRule(t, rule, expectedRule)

	MustEqualCSS(t, stylesheet.String(), expectedCSS)
}

func TestAtRuleCharset(t *testing.T) {
	css := `@charset "UTF-8";`

	expectedRule := &Rule{
		Kind:    AT_RULE,
		Name:    "@charset",
		Prelude: "\"UTF-8\"",
	}

	expectedCSS := `@charset "UTF-8";`

	stylesheet := MustParse(t, css, 1)
	rule := stylesheet.Rules[0]

	MustEqualRule(t, rule, expectedRule)

	MustEqualCSS(t, stylesheet.String(), expectedCSS)
}

func TestAtRuleCounterStyle(t *testing.T) {
	css := `@counter-style footnote {
  system: symbolic;
  symbols: '*' ⁑ † ‡;
  suffix: '';
}`

	expectedRule := &Rule{
		Kind:    AT_RULE,
		Name:    "@counter-style",
		Prelude: "footnote",
		Declarations: []*Declaration{
			&Declaration{
				Property: "system",
				Value:    "symbolic",
			},
			&Declaration{
				Property: "symbols",
				Value:    "'*' ⁑ † ‡",
			},
			&Declaration{
				Property: "suffix",
				Value:    "''",
			},
		},
	}

	stylesheet := MustParse(t, css, 1)
	rule := stylesheet.Rules[0]

	MustEqualRule(t, rule, expectedRule)

	MustEqualCSS(t, stylesheet.String(), css)
}

func TestAtRuleDocument(t *testing.T) {
	css := `@document url(http://www.w3.org/),
               url-prefix(http://www.w3.org/Style/),
               domain(mozilla.org),
               regexp("https:.*")
{
  /* CSS rules here apply to:
     + The page "http://www.w3.org/".
     + Any page whose URL begins with "http://www.w3.org/Style/"
     + Any page whose URL's host is "mozilla.org" or ends with
       ".mozilla.org"
     + Any page whose URL starts with "https:" */

  /* make the above-mentioned pages really ugly */
  body { color: purple; background: yellow; }
}`

	expectedRule := &Rule{
		Kind: AT_RULE,
		Name: "@document",
		Prelude: `url(http://www.w3.org/),
               url-prefix(http://www.w3.org/Style/),
               domain(mozilla.org),
               regexp("https:.*")`,
		Rules: []*Rule{
			&Rule{
				Kind:    QUALIFIED_RULE,
				Prelude: "body",
				Declarations: []*Declaration{
					&Declaration{
						Property: "color",
						Value:    "purple",
					},
					&Declaration{
						Property: "background",
						Value:    "yellow",
					},
				},
			},
		},
	}

	expectCSS := `@document url(http://www.w3.org/),
               url-prefix(http://www.w3.org/Style/),
               domain(mozilla.org),
               regexp("https:.*") {
  body {
    color: purple;
    background: yellow;
  }
}`

	stylesheet := MustParse(t, css, 1)
	rule := stylesheet.Rules[0]

	MustEqualRule(t, rule, expectedRule)

	MustEqualCSS(t, stylesheet.String(), expectCSS)
}

func TestAtRuleFontFace(t *testing.T) {
	css := `@font-face {
  font-family: MyHelvetica;
  src: local("Helvetica Neue Bold"),
       local("HelveticaNeue-Bold"),
       url(MgOpenModernaBold.ttf);
  font-weight: bold;
}`

	expectedRule := &Rule{
		Kind: AT_RULE,
		Name: "@font-face",
		Declarations: []*Declaration{
			&Declaration{
				Property: "font-family",
				Value:    "MyHelvetica",
			},
			&Declaration{
				Property: "src",
				Value: `local("Helvetica Neue Bold"),
       local("HelveticaNeue-Bold"),
       url(MgOpenModernaBold.ttf)`,
			},
			&Declaration{
				Property: "font-weight",
				Value:    "bold",
			},
		},
	}

	stylesheet := MustParse(t, css, 1)
	rule := stylesheet.Rules[0]

	MustEqualRule(t, rule, expectedRule)

	MustEqualCSS(t, stylesheet.String(), css)
}

func TestAtRuleFontFeatureValues(t *testing.T) {
	css := `@font-feature-values Font Two { /* How to activate nice-style in Font Two */
  @styleset {
    nice-style: 4;
  }
}`
	expectedRule := &Rule{
		Kind:    AT_RULE,
		Name:    "@font-feature-values",
		Prelude: "Font Two",
		Rules: []*Rule{
			&Rule{
				Kind: AT_RULE,
				Name: "@styleset",
				Declarations: []*Declaration{
					&Declaration{
						Property: "nice-style",
						Value:    "4",
					},
				},
			},
		},
	}

	expectedCSS := `@font-feature-values Font Two {
  @styleset {
    nice-style: 4;
  }
}`

	stylesheet := MustParse(t, css, 1)
	rule := stylesheet.Rules[0]

	MustEqualRule(t, rule, expectedRule)

	MustEqualCSS(t, stylesheet.String(), expectedCSS)
}

func TestAtRuleImport(t *testing.T) {
	css := `@import "my-styles.css";
@import url('landscape.css') screen and (orientation:landscape);`

	expectedRule1 := &Rule{
		Kind:    AT_RULE,
		Name:    "@import",
		Prelude: "\"my-styles.css\"",
	}

	expectedRule2 := &Rule{
		Kind:    AT_RULE,
		Name:    "@import",
		Prelude: "url('landscape.css') screen and (orientation:landscape)",
	}

	stylesheet := MustParse(t, css, 2)

	MustEqualRule(t, stylesheet.Rules[0], expectedRule1)
	MustEqualRule(t, stylesheet.Rules[1], expectedRule2)

	MustEqualCSS(t, stylesheet.String(), css)
}

func TestAtRuleKeyframes(t *testing.T) {
	css := `@keyframes identifier {
  0% { top: 0; left: 0; }
  100% { top: 100px; left: 100%; }
}`
	expectedRule := &Rule{
		Kind:    AT_RULE,
		Name:    "@keyframes",
		Prelude: "identifier",
		Rules: []*Rule{
			&Rule{
				Kind:    QUALIFIED_RULE,
				Prelude: "0%",
				Declarations: []*Declaration{
					&Declaration{
						Property: "top",
						Value:    "0",
					},
					&Declaration{
						Property: "left",
						Value:    "0",
					},
				},
			},
			&Rule{
				Kind:    QUALIFIED_RULE,
				Prelude: "100%",
				Declarations: []*Declaration{
					&Declaration{
						Property: "top",
						Value:    "100px",
					},
					&Declaration{
						Property: "left",
						Value:    "100%",
					},
				},
			},
		},
	}

	expectedCSS := `@keyframes identifier {
  0% {
    top: 0;
    left: 0;
  }
  100% {
    top: 100px;
    left: 100%;
  }
}`

	stylesheet := MustParse(t, css, 1)
	rule := stylesheet.Rules[0]

	MustEqualRule(t, rule, expectedRule)

	MustEqualCSS(t, stylesheet.String(), expectedCSS)
}

func TestAtRuleMedia(t *testing.T) {
	css := `@media screen, print {
  body { line-height: 1.2 }
}`
	expectedRule := &Rule{
		Kind:    AT_RULE,
		Name:    "@media",
		Prelude: "screen, print",
		Rules: []*Rule{
			&Rule{
				Kind:    QUALIFIED_RULE,
				Prelude: "body",
				Declarations: []*Declaration{
					&Declaration{
						Property: "line-height",
						Value:    "1.2",
					},
				},
			},
		},
	}

	expectedCSS := `@media screen, print {
  body {
    line-height: 1.2;
  }
}`

	stylesheet := MustParse(t, css, 1)
	rule := stylesheet.Rules[0]

	MustEqualRule(t, rule, expectedRule)

	MustEqualCSS(t, stylesheet.String(), expectedCSS)
}

func TestAtRuleNamespace(t *testing.T) {
	css := `@namespace svg url(http://www.w3.org/2000/svg);`
	expectedRule := &Rule{
		Kind:    AT_RULE,
		Name:    "@namespace",
		Prelude: "svg url(http://www.w3.org/2000/svg)",
	}

	stylesheet := MustParse(t, css, 1)
	rule := stylesheet.Rules[0]

	MustEqualRule(t, rule, expectedRule)

	MustEqualCSS(t, stylesheet.String(), css)
}

func TestAtRulePage(t *testing.T) {
	css := `@page :left {
  margin-left: 4cm;
  margin-right: 3cm;
}`
	expectedRule := &Rule{
		Kind:    AT_RULE,
		Name:    "@page",
		Prelude: ":left",
		Declarations: []*Declaration{
			&Declaration{
				Property: "margin-left",
				Value:    "4cm",
			},
			&Declaration{
				Property: "margin-right",
				Value:    "3cm",
			},
		},
	}

	stylesheet := MustParse(t, css, 1)
	rule := stylesheet.Rules[0]

	MustEqualRule(t, rule, expectedRule)

	MustEqualCSS(t, stylesheet.String(), css)
}

func TestAtRuleSupports(t *testing.T) {
	css := `@supports (animation-name: test) {
    /* specific CSS applied when animations are supported unprefixed */
    @keyframes { /* @supports being a CSS conditional group at-rule, it can includes other relevent at-rules */
      0% { top: 0; left: 0; }
      100% { top: 100px; left: 100%; }
    }
}`
	expectedRule := &Rule{
		Kind:    AT_RULE,
		Name:    "@supports",
		Prelude: "(animation-name: test)",
		Rules: []*Rule{
			&Rule{
				Kind: AT_RULE,
				Name: "@keyframes",
				Rules: []*Rule{
					&Rule{
						Kind:    QUALIFIED_RULE,
						Prelude: "0%",
						Declarations: []*Declaration{
							&Declaration{
								Property: "top",
								Value:    "0",
							},
							&Declaration{
								Property: "left",
								Value:    "0",
							},
						},
					},
					&Rule{
						Kind:    QUALIFIED_RULE,
						Prelude: "100%",
						Declarations: []*Declaration{
							&Declaration{
								Property: "top",
								Value:    "100px",
							},
							&Declaration{
								Property: "left",
								Value:    "100%",
							},
						},
					},
				},
			},
		},
	}

	expectedCSS := `@supports (animation-name: test) {
  @keyframes {
    0% {
      top: 0;
      left: 0;
    }
    100% {
      top: 100px;
      left: 100%;
    }
  }
}`

	stylesheet := MustParse(t, css, 1)
	rule := stylesheet.Rules[0]

	MustEqualRule(t, rule, expectedRule)

	MustEqualCSS(t, stylesheet.String(), expectedCSS)
}
