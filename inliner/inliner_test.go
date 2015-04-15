package inliner

import (
	"fmt"
	"testing"
)

// Simple rule inlining with two declarations
func TestInliner(t *testing.T) {
	input := `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
  <head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
<style type="text/css">
  p {
    font-family: 'Helvetica Neue', Verdana, sans-serif;
    color: #eee;
  }
</style>
  </head>
  <body>
    <p>
      Inline me please!
    </p>
</body>
</html>`

	expectedOutput := `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd"><html xmlns="http://www.w3.org/1999/xhtml"><head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>

  </head>
  <body>
    <p style="color: #eee; font-family: &#39;Helvetica Neue&#39;, Verdana, sans-serif;">
      Inline me please!
    </p>

</body></html>`

	output, err := Inline(input)
	if err != nil {
		t.Fatal("Failed to inline html:", err)
	}

	if output != expectedOutput {
		t.Fatal(fmt.Sprintf("CSS inliner error\nExpected:\n\"%s\"\nGot:\n\"%s\"", expectedOutput, output))
	}
}

// Alreadu inlined style has more priority than <style>
func TestInlineStylePriority(t *testing.T) {
	input := `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
  <head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
<style type="text/css">
  p {
    font-family: 'Helvetica Neue', Verdana, sans-serif;
    color: #eee;
  }
</style>
  </head>
  <body>
    <p style="color: #222;">
      Inline me please!
    </p>
</body>
</html>`

	expectedOutput := `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd"><html xmlns="http://www.w3.org/1999/xhtml"><head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>

  </head>
  <body>
    <p style="color: #222; font-family: &#39;Helvetica Neue&#39;, Verdana, sans-serif;">
      Inline me please!
    </p>

</body></html>`

	output, err := Inline(input)
	if err != nil {
		t.Fatal("Failed to inline html:", err)
	}

	if output != expectedOutput {
		t.Fatal(fmt.Sprintf("CSS inliner error\nExpected:\n\"%s\"\nGot:\n\"%s\"", expectedOutput, output))
	}
}

// Pseudo-class selectors can't be inlined
func TestNotInlinable(t *testing.T) {
	input := `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
  <head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
<style type="text/css">
    a:hover {
      color: #2795b6 !important;
    }

    a:active {
      color: #2795b6 !important;
    }
</style>
  </head>
  <body>
    <p>
      <a href="http://aymerick.com">Superbe website</a>
    </p>
</body>
</html>`

	expectedOutput := `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd"><html xmlns="http://www.w3.org/1999/xhtml"><head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>

  <style type="text/css">
a:hover {
  color: #2795b6 !important;
}
a:active {
  color: #2795b6 !important;
}
</style></head>
  <body>
    <p>
      <a href="http://aymerick.com">Superbe website</a>
    </p>

</body></html>`

	output, err := Inline(input)
	if err != nil {
		t.Fatal("Failed to inline html:", err)
	}

	if output != expectedOutput {
		t.Fatal(fmt.Sprintf("CSS inliner error\nExpected:\n\"%s\"\nGot:\n\"%s\"", expectedOutput, output))
	}
}

// Some styles causes insertion of additional element attributes
func TestStyleToAttr(t *testing.T) {
	input := `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
  <head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
<style type="text/css">
  body {
    background-color: #f2f2f2;
  }
</style>
  </head>
  <body>
    <p>
      Inline me please!
    </p>
</body>
</html>`

	expectedOutput := `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd"><html xmlns="http://www.w3.org/1999/xhtml"><head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>

  </head>
  <body style="background-color: #f2f2f2;" bgcolor="#f2f2f2">
    <p>
      Inline me please!
    </p>

</body></html>`

	output, err := Inline(input)
	if err != nil {
		t.Fatal("Failed to inline html:", err)
	}

	if output != expectedOutput {
		t.Fatal(fmt.Sprintf("CSS inliner error\nExpected:\n\"%s\"\nGot:\n\"%s\"", expectedOutput, output))
	}
}
