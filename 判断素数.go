package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Print("请输入一个整数")
	fmt.Scanf("%d", &a)
	if a%2 == 1 {
		fmt.Println("这是一个素数")
	} else {
		fmt.Println("这不是一个素数")
	}

}
