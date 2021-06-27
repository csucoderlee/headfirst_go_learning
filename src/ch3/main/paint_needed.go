package main

import "fmt"

func main() {
	paintNeeded(12.8, 16.5)
	println(calculate())
}

//打印结果
func paintNeeded(width float64, height float64) {
	area := width * height
	fmt.Println(" the area is ", area)
	fmt.Printf(" the area is %.2f \n", area)
}

//返回结果
func paintNeed1(width float64, height float64) float64 {
	return width * height
}

func paintNeed2(width float64, height float64) (total float64, err error) {
	if width < 0 {
		return 0, fmt.Errorf("the width %0.2f value is invalid", width)
	}

	if height < 0 {
		return 0, fmt.Errorf("the height %0.2f value is invalid", height)
	}

	return width * height, nil
}

func calculate() float64 {
	var total, area float64

	area = paintNeed1(4, 3)

	total += area
	area = paintNeed1(4, 3)
	total += area

	return total
}
