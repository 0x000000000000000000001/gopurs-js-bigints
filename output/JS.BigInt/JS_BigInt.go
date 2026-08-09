package JS_BigInt

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Int "gopurs/output/Data.Int"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	unsafe "unsafe"
)

var cache_showBigInt gopurs_runtime.Value
var once_showBigInt sync.Once
func Get_showBigInt() gopurs_runtime.Value {
	once_showBigInt.Do(func() {
		cache_showBigInt = gopurs_runtime.RecordDict1("show", Get_toString())
	})
	return cache_showBigInt
}

var cache_semiringBigInt gopurs_runtime.Value
var once_semiringBigInt sync.Once
func Get_semiringBigInt() gopurs_runtime.Value {
	once_semiringBigInt.Do(func() {
		cache_semiringBigInt = gopurs_runtime.RecordDict4("add", "mul", "one", "zero", Get_biAdd(), Get_biMul(), Get_biOne(), Get_biZero())
	})
	return cache_semiringBigInt
}

var cache_ringBigInt gopurs_runtime.Value
var once_ringBigInt sync.Once
func Get_ringBigInt() gopurs_runtime.Value {
	once_ringBigInt.Do(func() {
		cache_ringBigInt = gopurs_runtime.RecordDict2("Semiring0", "sub", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_semiringBigInt()
}), Get_biSub())
	})
	return cache_ringBigInt
}

var cache_eqBigInt gopurs_runtime.Value
var once_eqBigInt sync.Once
func Get_eqBigInt() gopurs_runtime.Value {
	once_eqBigInt.Do(func() {
		cache_eqBigInt = gopurs_runtime.RecordDict1("eq", Get_biEquals())
	})
	return cache_eqBigInt
}

var cache_ordBigInt gopurs_runtime.Value
var once_ordBigInt sync.Once
func Get_ordBigInt() gopurs_runtime.Value {
	once_ordBigInt.Do(func() {
		cache_ordBigInt = gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eqBigInt()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
v_2_0 := gopurs_runtime.Apply2(Get_biCompare(), x_0, y_1)
_ = v_2_0
var __t1 gopurs_runtime.Value
{
if (v_2_0.IntVal) == (1) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}
goto end_branch_1
} else {

}
}
{
if (v_2_0.IntVal) == (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}
}
end_branch_1:
return __t1
})
}))
	})
	return cache_ordBigInt
}

var cache_commutativeRingBigInt gopurs_runtime.Value
var once_commutativeRingBigInt sync.Once
func Get_commutativeRingBigInt() gopurs_runtime.Value {
	once_commutativeRingBigInt.Do(func() {
		cache_commutativeRingBigInt = gopurs_runtime.RecordDict1("Ring0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_ringBigInt()
}))
	})
	return cache_commutativeRingBigInt
}

var cache_euclideanRingBigInt gopurs_runtime.Value
var once_euclideanRingBigInt sync.Once
func Get_euclideanRingBigInt() gopurs_runtime.Value {
	once_euclideanRingBigInt.Do(func() {
		cache_euclideanRingBigInt = gopurs_runtime.RecordDict4("CommutativeRing0", "degree", "div", "mod", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_commutativeRingBigInt()
}), Get_biDegree(), Get_biDiv(), Get_biMod())
	})
	return cache_euclideanRingBigInt
}

var cache_toInt gopurs_runtime.Value
var once_toInt sync.Once
func Get_toInt() gopurs_runtime.Value {
	once_toInt.Do(func() {
		cache_toInt = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_toInt(x_0_box))}
})
	})
	return cache_toInt
}

var cache_odd gopurs_runtime.Value
var once_odd sync.Once
func Get_odd() gopurs_runtime.Value {
	once_odd.Do(func() {
		cache_odd = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_odd(x_0_box))
})
	})
	return cache_odd
}

var cache_fromTLInt gopurs_runtime.Value
var once_fromTLInt sync.Once
func Get_fromTLInt() gopurs_runtime.Value {
	once_fromTLInt.Do(func() {
		cache_fromTLInt = gopurs_runtime.Func3(func(_dollar__unused_0_box gopurs_runtime.Value, dictReflectable_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromTLInt(_dollar__unused_0_box, dictReflectable_1_box, v_2_box)
})
	})
	return cache_fromTLInt
}

var cache_fromStringAs gopurs_runtime.Value
var once_fromStringAs sync.Once
func Get_fromStringAs() gopurs_runtime.Value {
	once_fromStringAs.Do(func() {
		cache_fromStringAs = gopurs_runtime.Apply2(Get_fromStringAsImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})
	})
	return cache_fromStringAs
}

var cache_fromString gopurs_runtime.Value
var once_fromString sync.Once
func Get_fromString() gopurs_runtime.Value {
	once_fromString.Do(func() {
		cache_fromString = gopurs_runtime.Apply2(Get_fromStringImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})
	})
	return cache_fromString
}

var cache_fromNumber gopurs_runtime.Value
var once_fromNumber sync.Once
func Get_fromNumber() gopurs_runtime.Value {
	once_fromNumber.Do(func() {
		cache_fromNumber = gopurs_runtime.Apply2(Get_fromNumberImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})
	})
	return cache_fromNumber
}

var cache_even gopurs_runtime.Value
var once_even sync.Once
func Get_even() gopurs_runtime.Value {
	once_even.Do(func() {
		cache_even = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_even(x_0_box))
})
	})
	return cache_even
}

var cache_parity gopurs_runtime.Value
var once_parity sync.Once
func Get_parity() gopurs_runtime.Value {
	once_parity.Do(func() {
		cache_parity = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_parity(n_0_box)
})
	})
	return cache_parity
}

func Call_toInt(x_0_loop gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[int64] {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return (*pkg_Data_Maybe.Constructor_Just[int64])(gopurs_runtime.Apply(pkg_Data_Int.Get_fromNumber(), gopurs_runtime.Apply(Get_toNumber(), x_0)).UnsafePtr)
}

func Call_odd(x_0_loop gopurs_runtime.Value) bool {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Eq.Get_eqBoolean(), "eq"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_eqBigInt(), "eq"), gopurs_runtime.Apply2(Get_and(), x_0, gopurs_runtime.RecordGet(Get_semiringBigInt(), "one")), gopurs_runtime.RecordGet(Get_semiringBigInt(), "zero")), gopurs_runtime.Bool(false)).IntVal) != (0)
}

func Call_fromTLInt(_dollar__unused_0_loop gopurs_runtime.Value, dictReflectable_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var dictReflectable_1 gopurs_runtime.Value = dictReflectable_1_loop
_ = dictReflectable_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply(Get_fromTypeLevelInt(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictReflectable_1, "reflectType"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}))
}

func Call_even(x_0_loop gopurs_runtime.Value) bool {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_eqBigInt(), "eq"), gopurs_runtime.Apply2(Get_and(), x_0, gopurs_runtime.RecordGet(Get_semiringBigInt(), "one")), gopurs_runtime.RecordGet(Get_semiringBigInt(), "zero")).IntVal) != (0)
}

func Call_parity(n_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var n_0 gopurs_runtime.Value = n_0_loop
_ = n_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_eqBigInt(), "eq"), gopurs_runtime.Apply2(Get_and(), n_0, gopurs_runtime.RecordGet(Get_semiringBigInt(), "one")), gopurs_runtime.RecordGet(Get_semiringBigInt(), "zero")).IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2591059121, UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 658452902, UnsafePtr: nil}
}
end_branch_0:
return __t0
}

func Get_and() gopurs_runtime.Value {
	return _Gopurs_And
}

func Get_asIntN() gopurs_runtime.Value {
	return _Gopurs_AsIntN
}

func Get_asUintN() gopurs_runtime.Value {
	return _Gopurs_AsUintN
}

func Get_biAdd() gopurs_runtime.Value {
	return _Gopurs_BiAdd
}

func Get_biCompare() gopurs_runtime.Value {
	return _Gopurs_BiCompare
}

func Get_biDegree() gopurs_runtime.Value {
	return _Gopurs_BiDegree
}

func Get_biDiv() gopurs_runtime.Value {
	return _Gopurs_BiDiv
}

func Get_biEquals() gopurs_runtime.Value {
	return _Gopurs_BiEquals
}

func Get_biMod() gopurs_runtime.Value {
	return _Gopurs_BiMod
}

func Get_biMul() gopurs_runtime.Value {
	return _Gopurs_BiMul
}

func Get_biOne() gopurs_runtime.Value {
	return _Gopurs_BiOne
}

func Get_biSub() gopurs_runtime.Value {
	return _Gopurs_BiSub
}

func Get_biZero() gopurs_runtime.Value {
	return _Gopurs_BiZero
}

func Get_fromInt() gopurs_runtime.Value {
	return _Gopurs_FromInt
}

func Get_fromNumberImpl() gopurs_runtime.Value {
	return _Gopurs_FromNumberImpl
}

func Get_fromStringAsImpl() gopurs_runtime.Value {
	return _Gopurs_FromStringAsImpl
}

func Get_fromStringImpl() gopurs_runtime.Value {
	return _Gopurs_FromStringImpl
}

func Get_fromTypeLevelInt() gopurs_runtime.Value {
	return _Gopurs_FromTypeLevelInt
}

func Get_not() gopurs_runtime.Value {
	return _Gopurs_Not
}

func Get_or() gopurs_runtime.Value {
	return _Gopurs_Or
}

func Get_pow() gopurs_runtime.Value {
	return _Gopurs_Pow
}

func Get_shl() gopurs_runtime.Value {
	return _Gopurs_Shl
}

func Get_shr() gopurs_runtime.Value {
	return _Gopurs_Shr
}

func Get_toNumber() gopurs_runtime.Value {
	return _Gopurs_ToNumber
}

func Get_toString() gopurs_runtime.Value {
	return _Gopurs_ToString
}

func Get_toStringAs() gopurs_runtime.Value {
	return _Gopurs_ToStringAs
}

func Get_xor() gopurs_runtime.Value {
	return _Gopurs_Xor
}
