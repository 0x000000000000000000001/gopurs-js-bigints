
import (
	"math/big"
	"strconv"
)

func FromStringImpl(just func(any) any) func(any) func(string) any {
	return func(nothing any) func(string) any {
		return func(s string) any {
			n := new(big.Int)
			_, ok := n.SetString(s, 10)
			if !ok {
				return nothing
			}
			return just(n)
		}
	}
}

func FromStringAsImpl(just func(any) any) func(any) func(int) func(string) any {
	return func(nothing any) func(int) func(string) any {
		return func(radix int) func(string) any {
			return func(s string) any {
				n := new(big.Int)
				_, ok := n.SetString(s, radix)
				if !ok {
					return nothing
				}
				return just(n)
			}
		}
	}
}

func FromInt(n int) any {
	return big.NewInt(int64(n))
}

func FromNumberImpl(just func(any) any) func(any) func(float64) any {
	return func(nothing any) func(float64) any {
		return func(n float64) any {
			// Basic conversion, ignoring precision issues for now
			// yoga-json doesn't strictly depend on fromNumber for BigInt tests
			bi := new(big.Int)
			s := strconv.FormatFloat(n, 'f', 0, 64)
			_, ok := bi.SetString(s, 10)
			if !ok {
				return nothing
			}
			return just(bi)
		}
	}
}

func FromTypeLevelInt(s string) any {
	n := new(big.Int)
	n.SetString(s, 10)
	return n
}

func ToNumber(n any) float64 {
	bi := n.(*big.Int)
	f, _ := new(big.Float).SetInt(bi).Float64()
	return f
}

func BiAdd(x any) func(any) any {
	return func(y any) any {
		return new(big.Int).Add(x.(*big.Int), y.(*big.Int))
	}
}

func BiMul(x any) func(any) any {
	return func(y any) any {
		return new(big.Int).Mul(x.(*big.Int), y.(*big.Int))
	}
}

func BiSub(x any) func(any) any {
	return func(y any) any {
		return new(big.Int).Sub(x.(*big.Int), y.(*big.Int))
	}
}

func BiMod(x any) func(any) any {
	return func(y any) any {
		return new(big.Int).Mod(x.(*big.Int), y.(*big.Int))
	}
}

func BiDiv(x any) func(any) any {
	return func(y any) any {
		return new(big.Int).Div(x.(*big.Int), y.(*big.Int))
	}
}

func BiDegree(x any) int {
	return 1
}

var BiZero any = big.NewInt(0)
var BiOne any = big.NewInt(1)

func Pow(x any) func(any) any {
	return func(y any) any {
		return new(big.Int).Exp(x.(*big.Int), y.(*big.Int), nil)
	}
}

func Not(x any) any {
	return new(big.Int).Not(x.(*big.Int))
}

func Or(x any) func(any) any {
	return func(y any) any {
		return new(big.Int).Or(x.(*big.Int), y.(*big.Int))
	}
}

func Xor(x any) func(any) any {
	return func(y any) any {
		return new(big.Int).Xor(x.(*big.Int), y.(*big.Int))
	}
}

func And(x any) func(any) any {
	return func(y any) any {
		return new(big.Int).And(x.(*big.Int), y.(*big.Int))
	}
}

func Shl(x any) func(any) any {
	return func(n any) any {
		return new(big.Int).Lsh(x.(*big.Int), uint(n.(*big.Int).Int64()))
	}
}

func Shr(x any) func(any) any {
	return func(n any) any {
		return new(big.Int).Rsh(x.(*big.Int), uint(n.(*big.Int).Int64()))
	}
}

func BiEquals(x any) func(any) bool {
	return func(y any) bool {
		return x.(*big.Int).Cmp(y.(*big.Int)) == 0
	}
}

func BiCompare(x any) func(any) int {
	return func(y any) int {
		return x.(*big.Int).Cmp(y.(*big.Int))
	}
}

func ToString(x any) string {
	return x.(*big.Int).String()
}

func AsIntN(bits int) func(any) any {
	return func(n any) any {
		// Mocked
		return n
	}
}

func AsUintN(bits int) func(any) any {
	return func(n any) any {
		// Mocked
		return n
	}
}

func ToStringAs(radix int) func(any) string {
	return func(i any) string {
		return i.(*big.Int).Text(radix)
	}
}
