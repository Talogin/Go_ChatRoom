package main

import (
	"fmt"
)

func main() {

	var key int

	var loop = true
	//初始化Menu
	//选择登录或者注册
	for loop {
		fmt.Println("------------Please Select the Option You Want------------")
		fmt.Println("\t\t\t 1. Login")
		fmt.Println("\t\t\t 2. Register")
		fmt.Println("\t\t\t 3. Exit")

		fmt.Scanf("%d", &key)
		switch key {
		case 1:
			fmt.Println("Login")
			loop = false
		case 2:
			fmt.Println("Register")
			loop = false
		case 3:
			fmt.Println("Exit")
			loop = false
		default:
			fmt.Println("The operation you choose does not exist, please try again.")
		}
	}

}
