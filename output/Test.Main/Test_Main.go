package Test_Main

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Effect "gopurs/output/Effect"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	pkg_JS_BigInt "gopurs/output/JS.BigInt"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
)

var cache_main gopurs_runtime.Value
var once_main sync.Once
func Get_main() gopurs_runtime.Value {
	once_main.Do(func() {
		cache_main = gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Control_Bind.Get_discardUnit(), "discard"), pkg_Effect.Get_bindEffect(), gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("Running BigInt tests...")), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
c_1_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_JS_BigInt.Get_semiringBigInt(), "add"), gopurs_runtime.Apply(pkg_JS_BigInt.Get_fromInt(), gopurs_runtime.Int(10)), gopurs_runtime.Apply(pkg_JS_BigInt.Get_fromInt(), gopurs_runtime.Int(20)))
_ = c_1_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_JS_BigInt.Get_eqBigInt(), "eq"), c_1_0, gopurs_runtime.Apply(pkg_JS_BigInt.Get_fromInt(), gopurs_runtime.Int(30))).IntVal) != (0) {
__t1 = gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("10 + 20 = 30 (OK)"))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("Failed: expected 30, got "), gopurs_runtime.Apply(pkg_JS_BigInt.Get_toString(), c_1_0)))
}
end_branch_1:
return __t1
}))
	})
	return cache_main
}




