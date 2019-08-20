package basic

func DemoSwitch() {
	println("begin", getFuncName(1))
	flag := "09"
	switch flag {
	case "01":
		println("01")
	case "02":
		println(02)
	case "09":
		println(true)
	default:
		println(false)
	}
	println("end", getFuncName(1))
	println(getFuncName(2))
}
