package print

import (
	"fmt"

	"github.com/csjoy/package-demo/formatter/internal"
)

// Format function
func Format(num int) string {
	return fmt.Sprintf("The number is %d\n", num)
}

func UseDoubler(s string) string {
	return internal.Echos(s)
}
