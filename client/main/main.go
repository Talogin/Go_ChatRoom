package main

import (
	"fmt"
	"Go_ChatRoom/client/process"
)

func main() {

	var key int

	var userId int
	var userPwd string
	// var userName string
	
	var loop = true
	
	for loop {
		fmt.Println("------------Please Select the Option You Want------------")
		fmt.Println("\t\t\t 1. Login")
		fmt.Println("\t\t\t 2. Register")
		fmt.Println("\t\t\t 3. Exit")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("............Login............")
			fmt.Println("Please enter your user id:")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("Please enter your password:")
			fmt.Scanf("%s\n", &userPwd)
			up := &process.UserProcess {}
			err := up.Login(userId, userPwd)
			if err != nil {
				fmt.Println("error")
			}
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
