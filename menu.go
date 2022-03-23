package main

import "fmt"

func main(){
	var code string
	fmt.Println("输入指令1、2、3")
	for {
		fmt.Scanf("%s",&code)
		switch code {
		case "1":
			fmt.Println("输入指令1")
	    case "2":
		    fmt.Println("输入指令2")
		case "3":
		    fmt.Println("输入指令3")
		default:
			fmt.Println("无效指令")
		}
	}
}