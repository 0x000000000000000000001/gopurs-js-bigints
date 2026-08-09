import (
	"math/big"
	"strconv"
)

func toBigInt(v any) *big.Int {
	if val, ok := v.(gopurs_runtime.Value); ok {
		return val.Item.(*big.Int)
	}
	return v.(*big.Int)
}

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
	bi := toBigInt(n)
	f, _ := new(big.Float).SetInt(bi).Float64()
	return f
}

func BiAdd(x any) func(any) any {
	return func(y any) any {
		return new(big.Int).Add(toBigInt(x), toBigInt(y))
	}
}

func BiMul(x any) func(any) any {
	return func(y any) any {
		return new(big.Int).Mul(toBigInt(x), toBigInt(y))
	}
}

func BiSub(x any) func(any) any {
	return func(y any) any {
		return new(big.Int).Sub(toBigInt(x), toBigInt(y))
	}
}

func BiMod(x any) func(any) any {
	return func(y any) any {
		return new(big.Int).Mod(toBigInt(x), toBigInt(y))
	}
}

func BiDiv(x any) func(any) any {
	return func(y any) any {
		return new(big.Int).Div(toBigInt(x), toBigInt(y))
	}
}

func BiDegree(x any) int {
	return 1
}

var BiZero any = big.NewInt(0)
var BiOne any = big.NewInt(1)

func Pow(x any) func(any) any {
	return func(y any) any {
		return new(big.Int).Exp(toBigInt(x), toBigInt(y), nil)
	}
}

func Not(x any) any {
	return new(big.Int).Not(toBigInt(x))
}

func Or(x any) func(any) any {
	return func(y any) any {
		return new(big.Int).Or(toBigInt(x), toBigInt(y))
	}
}

func Xor(x any) func(any) any {
	return func(y any) any {
		return new(big.Int).Xor(toBigInt(x), toBigInt(y))
	}
}

func And(x any) func(any) any {
	return func(y any) any {
		return new(big.Int).And(toBigInt(x), toBigInt(y))
	}
}

func Shl(x any) func(any) any {
	return func(n any) any {
		return new(big.Int).Lsh(toBigInt(x), uint(toBigInt(n).Int64()))
	}
}

func Shr(x any) func(any) any {
	return func(n any) any {
		return new(big.Int).Rsh(toBigInt(x), uint(toBigInt(n).Int64()))
	}
}

func BiEquals(x any) func(any) bool {
	return func(y any) bool {
		return toBigInt(x).Cmp(toBigInt(y)) == 0
	}
}

func BiCompare(x any) func(any) int {
	return func(y any) int {
		return toBigInt(x).Cmp(toBigInt(y))
	}
}

func ToString(x any) string {
	return toBigInt(x).String()
}

func AsIntN(bits int) func(any) any {
	return func(n any) any {
		return n
	}
}

func AsUintN(bits int) func(any) any {
	return func(n any) any {
		return n
	}
}

func ToStringAs(radix int) func(any) string {
	return func(i any) string {
		return toBigInt(i).Text(radix)
	}
}
