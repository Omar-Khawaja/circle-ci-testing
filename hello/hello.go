package hello

import (
	"fmt"
)

func Hello(name string) int {
	return fmt.Sprintf("Hello, %s!", name)
}
