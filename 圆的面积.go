package main

import "fmt"

func main() {
	const PI = 3.14
	var radius float64
	fmt.Print("圆的半径：")
	fmt.Scanf("%f", &radius)
	area := PI * radius * radius
	fmt.Println("圆的面积为：%.2f", area)

}
