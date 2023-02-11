package internaluser

import (
	"github.com/csjoy/package-demo/formatter/internal"
)

func InternalUse(s string) string {
	return internal.Echos(s)
}
