# mocking

## requirement

```go
package mocking

import (
  "testing"
  "bytes"
)
func TestCountdown(t *testing.T) {
  buffer := &bytes.Buffer{}
  Countdown(buffer)

  got := buffer.String()
  want := "3"
  if got != want {
    t.Errorf("got %q want %q", got, want)
  }
}
```