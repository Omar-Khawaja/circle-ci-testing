package hello

import (
    "testing"
)

func TestHello(t *testing.T) {
    got := Hello("Omar")
    want := "Hello, Omar!"

    if got != want {
        t.Errorf("got '%s', wanted '%s'", got, want)
    }
}
