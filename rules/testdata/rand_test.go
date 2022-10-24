package testdata

import (
	"fmt"
	"testing"
)

func TestOverflow(t *testing.T) {
	bigInt := 1 << 64

	fmt.Println(bigInt)
	fmt.Println(bigInt + 1)
}
