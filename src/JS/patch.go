package JS

import (
	"math/big"
	"strconv"
)

// We define a helper that won't compile locally if it lacks gopurs_runtime
// but we'll remove the package decl so it gets prepended by gopurs.
// Wait, I'll just write a script to rewrite BigInt.go
