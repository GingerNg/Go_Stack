package basic

import "runtime"

func getFuncName(call int) string {
	/*
		獲取函數名稱
		call 1:自身　2:caller 調用者
	*/
	pc, _, _, _ := runtime.Caller(call)
	return runtime.FuncForPC(pc).Name()
}
