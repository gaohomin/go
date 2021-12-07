package main

import (
	"fmt"
)

func syaHi() {
	fmt.Println("hi")
}

func repeatLine(line string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(line)
	}
}

func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a height of %.2f is invalid", height)
	}
	area := width * height
	return area / 10.0, nil
}

func double(number int) int {
	root := number
	number = number * number
	fmt.Printf("The square of %d is %d", root, number)
	return number
}

func createPointer() *float64 {
	myFloat := 98.5
	return &myFloat
}

func createWithPointer(myBoolPointer *bool) {
	fmt.Println(*myBoolPointer)
}

func doubleWithPointerValue(number *int) {
	*number *= 2
}

func main() {
	//var myFloatPointer *float64 = createPointer()
	//fmt.Println(*myFloatPointer)
	//
	//var myBool bool = true
	//createWithPointer(&myBool)
	//
	//amount, err := paintNeeded(4.2, 3.0)
	//if err != nil {
	//	log.Fatal(err)
	//} else {
	//	fmt.Printf("%.2f liters needed \n", amount)
	//}
	//
	//var myInt int
	//myInt = 10
	//var myIntPointer *int
	//myIntPointer = &myInt
	//fmt.Println(reflect.TypeOf(&myInt))
	//fmt.Println(myIntPointer)
	//fmt.Println(*myIntPointer)
	//*myIntPointer = 90
	//fmt.Println(myInt)
	//
	//amount2 := 6
	//fmt.Println(amount2)
	//double(amount2)

	amountP := 6
	doubleWithPointerValue(&amountP)
	fmt.Println(amountP)
}
