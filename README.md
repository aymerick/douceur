# douceur

A simple CSS parser and inliner in Go.

Parser uses [Gorilla CSS3 tokenizer](https://github.com/gorilla/css). It is vaguely inspired by [CSS Syntax Module Level 3](http://www.w3.org/TR/css3-syntax) and [corresponding JS parser](https://github.com/tabatkins/parse-css).

Inliner uses [goquery](github.com/PuerkitoBio/goquery) to parse HTML.


## Tool usage

Install tool:

    $ go install github.com/aymerick/douceur

Parse a CSS file and display result:

    $ douceur parse inputfile.css

Inline CSS in an HTML document and display result:

    $ douceur inline inputfile.html


## Library usage

Fetch package:

    $ go get github.com/aymerick/douceur


### Parse CSS

```go
package main

import (
    "fmt"

    "github.com/aymerick/douceur/parser"
)

func main() {
    input := `body {
    /* D4rK s1T3 */
    background-color: black;
        }

  p     {
    /* Try to read that ! HAHA! */
    color: red; /* L O L */
 }
`

    stylesheet, err := parser.Parse(input)
    if err != nil {
        panic("OMG ! SO BUGGY !")
    }

    fmt.Print(stylesheet.String())
}
```

Displays:

```css
body {
  background-color: black;
}
p {
  color: red;
}
```


### Inline HTML

    @todo !!!


## Test

    go test ./... -v
