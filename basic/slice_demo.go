package basic

import "fmt"

func DemoSlice() {
	var numbers = make([]int, 3, 5) //make([]T, length, capacity) 按capacity的加长slice

	fmt.Sprint("numbers: %+v", numbers)

	println("emptyslice")
	emptySlice := []int{}
	emptySlice = append(emptySlice, 2, 3)
	for i := 0; i < len(emptySlice); i++ {
		//numbers[i] = i * 2
		println(emptySlice[i])
	}

	print(emptySlice)

	println("noemptyslice")
	println(numbers)
	numbers = append(numbers, 1, 2, 4)

	for i := 0; i < len(numbers); i++ {
		//numbers[i] = i * 2
		println(numbers[i])
	}

	for i, v := range numbers {
		println(i, v)
	}
	println(numbers)

}
