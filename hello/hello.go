package hello

import (
	"fmt"
)

func Hello(name string) (int, bool) {
	return fmt.Sprintf("Hello, %s!", name)
}
