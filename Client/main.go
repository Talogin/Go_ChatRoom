package Client

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
			fmt.Println("正在登录")
			loop = false
		case 2:
			fmt.Println("正在注册")
			loop = false
		case 3:
			fmt.Println("正在退出")
			loop = false
		default:
			fmt.Println("输入有误，请重新输入")
		}
	}

}
