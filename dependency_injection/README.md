# dependency_injection

## requirement

```go
package dependency_jection

import (
  "testing"
  "bytes"
)

func TestGreet(t *testing.T) {
  buffer := bytes.Buffer{}
  Greet(&buffer, "Chris")

  got := buffer.String()
  want := "Hello, Chris"

  if got != want {
    t.Errorf("got %q want %q", got, want)
  }
}
```