# douceur

A simple CSS parser in Go.

Vaguely inspired by [CSS Syntax Module Level 3](http://www.w3.org/TR/css3-syntax) and [corresponding JS parser](https://github.com/tabatkins/parse-css).

Uses [Gorilla CSS3 tokenizer](https://github.com/gorilla/css).


## Install

Fetch package:

    $ go get github.com/aymerick/douceur

Parse CSS:

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


## Tool usage

Install tool:

    $ go install github.com/aymerick/douceur

Parse a CSS file:

    $ douceur parse inputfile.css


## Test

    go test ./... -v
