package douceur

import "testing"

func TestQualifiedRule(t *testing.T) {
	css := `/* This is a comment */
p > a {
    color: blue;
    text-decoration: underline; /* This is a comment */
}`

	stylesheet, err := Parse(css)
	if err != nil {
		t.Fatal("Failed to parse css", err, css)
	}

	if len(stylesheet.Rules) != 1 {
		t.Fatal("Failed to parse Qualified Rules", css)
	}
}

func TestAtRule(t *testing.T) {
	css := `@charset "UTF-8";

@counter-style circled-alpha {
  system: fixed;
  symbols: Ⓐ Ⓑ Ⓒ Ⓓ Ⓔ Ⓕ Ⓖ Ⓗ Ⓘ Ⓙ Ⓚ Ⓛ Ⓜ Ⓝ Ⓞ Ⓟ Ⓠ Ⓡ Ⓢ Ⓣ Ⓤ Ⓥ Ⓦ Ⓧ Ⓨ Ⓩ;
  suffix: " ";
}

@document url(http://www.w3.org/),
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
}

@font-face {
  font-family: MyHelvetica;
  src: local("Helvetica Neue Bold"),
       local("HelveticaNeue-Bold"),
       url(MgOpenModernaBold.ttf);
  font-weight: bold;
}

@font-feature-values Font Two { /* How to activate nice-style in Font Two */
  @styleset {
    nice-style: 4;
  }
}

@import "my-styles.css";
@import url('landscape.css') screen and (orientation:landscape);

@keyframes identifier {
  0% { top: 0; left: 0; }
  30% { top: 50px; }
  68%, 72% { left: 50px; }
  100% { top: 100px; left: 100%; }
}

@media screen, print {
  body { line-height: 1.2 }
}

@namespace svg url(http://www.w3.org/2000/svg);

@page :left {
    margin-left: 4cm;
    margin-right: 3cm;
}

@supports (animation-name: test) {
    /* specific CSS applied when animations are supported unprefixed */
    @keyframes { /* @supports being a CSS conditional group at-rule, it can includes other relevent at-rules */
      0% { top: 0; left: 0; }
      30% { top: 50px; }
      68%, 72% { left: 50px; }
      100% { top: 100px; left: 100%; }
    }
}
`

	stylesheet, err := Parse(css)
	if err != nil {
		t.Fatal("Failed to parse css", err, css)
	}

	if len(stylesheet.Rules) != 12 {
		t.Fatal("Failed to parse At Rules", css)
	}
}
