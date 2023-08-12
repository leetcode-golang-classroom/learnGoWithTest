# Hello, World

## requirement

Write a hello function with input string as name, and will output hello prefix with input name

## test function

```go
package hello

import "testing"
func TestHello(t *testing.T) {
  got := Hello("Chris")
  want := "Hello, Chris"
  if got != want {
    t.Errorf("got %q want %q", got, want)
  }
}
```