package process

import "fmt"

func ShowMenu() {
	fmt.Println("-------恭喜xxx登录成功---------")
	fmt.Println("-------1. 显示在线用户列表---------")
	fmt.Println("-------2. 发送消息---------")
	fmt.Println("-------3. 信息列表---------")
	fmt.Println("-------4. 退出系统---------")
	fmt.Println("请选择(1-4):")
	var key int
	var content string


	fmt.Scanf("%d\n", &key) 
	switch key {
	case 1:
		
	}
}
