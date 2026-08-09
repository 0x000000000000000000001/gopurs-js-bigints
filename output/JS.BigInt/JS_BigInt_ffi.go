package JS_BigInt

import "gopurs/output/gopurs_runtime"

import (
	"math/big"
	"strconv"
)

func toBigInt(v any) *big.Int {
	if val, ok := v.(gopurs_runtime.Value); ok {
		return val.PtrVal().(*big.Int)
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


// --- Auto-generated FFI wrappers ---
var _Gopurs_And = // TAST: (Func [(ADT ["JS","BigInt","BigInt"] []), (ADT ["JS","BigInt","BigInt"] [])] (ADT ["JS","BigInt","BigInt"] []))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := And(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
				inner_res := go_res(arg)
				return gopurs_runtime.Box(inner_res)
			})
})
var _Gopurs_AsIntN = // TAST: (Func [Int, (ADT ["JS","BigInt","BigInt"] [])] (ADT ["JS","BigInt","BigInt"] []))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_res := AsIntN(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
				inner_res := go_res(arg)
				return gopurs_runtime.Box(inner_res)
			})
})
var _Gopurs_AsUintN = // TAST: (Func [Int, (ADT ["JS","BigInt","BigInt"] [])] (ADT ["JS","BigInt","BigInt"] []))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_res := AsUintN(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
				inner_res := go_res(arg)
				return gopurs_runtime.Box(inner_res)
			})
})
var _Gopurs_BiAdd = // TAST: (Func [(ADT ["JS","BigInt","BigInt"] []), (ADT ["JS","BigInt","BigInt"] [])] (ADT ["JS","BigInt","BigInt"] []))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := BiAdd(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
				inner_res := go_res(arg)
				return gopurs_runtime.Box(inner_res)
			})
})
var _Gopurs_BiCompare = // TAST: (Func [(ADT ["JS","BigInt","BigInt"] []), (ADT ["JS","BigInt","BigInt"] [])] Int)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := BiCompare(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
				inner_res := go_res(arg)
				return gopurs_runtime.Box(inner_res)
			})
})
var _Gopurs_BiDegree = // TAST: (Func [(ADT ["JS","BigInt","BigInt"] [])] Int)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := BiDegree(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_BiDiv = // TAST: (Func [(ADT ["JS","BigInt","BigInt"] []), (ADT ["JS","BigInt","BigInt"] [])] (ADT ["JS","BigInt","BigInt"] []))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := BiDiv(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
				inner_res := go_res(arg)
				return gopurs_runtime.Box(inner_res)
			})
})
var _Gopurs_BiEquals = // TAST: (Func [(ADT ["JS","BigInt","BigInt"] []), (ADT ["JS","BigInt","BigInt"] [])] Boolean)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := BiEquals(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
				inner_res := go_res(arg)
				return gopurs_runtime.Box(inner_res)
			})
})
var _Gopurs_BiMod = // TAST: (Func [(ADT ["JS","BigInt","BigInt"] []), (ADT ["JS","BigInt","BigInt"] [])] (ADT ["JS","BigInt","BigInt"] []))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := BiMod(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
				inner_res := go_res(arg)
				return gopurs_runtime.Box(inner_res)
			})
})
var _Gopurs_BiMul = // TAST: (Func [(ADT ["JS","BigInt","BigInt"] []), (ADT ["JS","BigInt","BigInt"] [])] (ADT ["JS","BigInt","BigInt"] []))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := BiMul(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
				inner_res := go_res(arg)
				return gopurs_runtime.Box(inner_res)
			})
})
var _Gopurs_BiOne = // TAST: (ADT ["JS","BigInt","BigInt"] [])
gopurs_runtime.Box(BiOne)
var _Gopurs_BiSub = // TAST: (Func [(ADT ["JS","BigInt","BigInt"] []), (ADT ["JS","BigInt","BigInt"] [])] (ADT ["JS","BigInt","BigInt"] []))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := BiSub(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
				inner_res := go_res(arg)
				return gopurs_runtime.Box(inner_res)
			})
})
var _Gopurs_BiZero = // TAST: (ADT ["JS","BigInt","BigInt"] [])
gopurs_runtime.Box(BiZero)
var _Gopurs_FromInt = // TAST: (Func [Int] (ADT ["JS","BigInt","BigInt"] []))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_res := FromInt(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_FromNumberImpl = // TAST: (Func [(Func [(TypeVar a)] (ADT ["Data","Maybe","Maybe"] [(TypeVar a)])), (ADT ["Data","Maybe","Maybe"] [(TypeVar a)]), Number] (ADT ["Data","Maybe","Maybe"] [(ADT ["JS","BigInt","BigInt"] [])]))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_res := FromNumberImpl(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
				inner_res := go_res(arg)
				return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
				inner_res := inner_res(gopurs_runtime.Unbox[float64](arg))
				return gopurs_runtime.Box(inner_res)
			})
			})
})
var _Gopurs_FromStringAsImpl = // TAST: (Func [(Func [(TypeVar a)] (ADT ["Data","Maybe","Maybe"] [(TypeVar a)])), (ADT ["Data","Maybe","Maybe"] [(TypeVar a)]), Int, String] (ADT ["Data","Maybe","Maybe"] [(ADT ["JS","BigInt","BigInt"] [])]))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_res := FromStringAsImpl(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
				inner_res := go_res(arg)
				return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
				inner_res := inner_res(gopurs_runtime.Unbox[int](arg))
				return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
				inner_res := inner_res(gopurs_runtime.Unbox[string](arg))
				return gopurs_runtime.Box(inner_res)
			})
			})
			})
})
var _Gopurs_FromStringImpl = // TAST: (Func [(Func [(TypeVar a)] (ADT ["Data","Maybe","Maybe"] [(TypeVar a)])), (ADT ["Data","Maybe","Maybe"] [(TypeVar a)]), String] (ADT ["Data","Maybe","Maybe"] [(ADT ["JS","BigInt","BigInt"] [])]))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_res := FromStringImpl(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
				inner_res := go_res(arg)
				return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
				inner_res := inner_res(gopurs_runtime.Unbox[string](arg))
				return gopurs_runtime.Box(inner_res)
			})
			})
})
var _Gopurs_FromTypeLevelInt = // TAST: (Func [String] (ADT ["JS","BigInt","BigInt"] []))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_res := FromTypeLevelInt(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Not = // TAST: (Func [(ADT ["JS","BigInt","BigInt"] [])] (ADT ["JS","BigInt","BigInt"] []))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Not(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Or = // TAST: (Func [(ADT ["JS","BigInt","BigInt"] []), (ADT ["JS","BigInt","BigInt"] [])] (ADT ["JS","BigInt","BigInt"] []))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Or(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
				inner_res := go_res(arg)
				return gopurs_runtime.Box(inner_res)
			})
})
var _Gopurs_Pow = // TAST: (Func [(ADT ["JS","BigInt","BigInt"] []), (ADT ["JS","BigInt","BigInt"] [])] (ADT ["JS","BigInt","BigInt"] []))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Pow(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
				inner_res := go_res(arg)
				return gopurs_runtime.Box(inner_res)
			})
})
var _Gopurs_Shl = // TAST: (Func [(ADT ["JS","BigInt","BigInt"] []), (ADT ["JS","BigInt","BigInt"] [])] (ADT ["JS","BigInt","BigInt"] []))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Shl(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
				inner_res := go_res(arg)
				return gopurs_runtime.Box(inner_res)
			})
})
var _Gopurs_Shr = // TAST: (Func [(ADT ["JS","BigInt","BigInt"] []), (ADT ["JS","BigInt","BigInt"] [])] (ADT ["JS","BigInt","BigInt"] []))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Shr(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
				inner_res := go_res(arg)
				return gopurs_runtime.Box(inner_res)
			})
})
var _Gopurs_ToNumber = // TAST: (Func [(ADT ["JS","BigInt","BigInt"] [])] Number)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := ToNumber(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_ToString = // TAST: (Func [(ADT ["JS","BigInt","BigInt"] [])] String)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := ToString(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_ToStringAs = // TAST: (Func [Int, (ADT ["JS","BigInt","BigInt"] [])] String)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_res := ToStringAs(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
				inner_res := go_res(arg)
				return gopurs_runtime.Box(inner_res)
			})
})
var _Gopurs_Xor = // TAST: (Func [(ADT ["JS","BigInt","BigInt"] []), (ADT ["JS","BigInt","BigInt"] [])] (ADT ["JS","BigInt","BigInt"] []))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Xor(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
				inner_res := go_res(arg)
				return gopurs_runtime.Box(inner_res)
			})
})