/*
switch 语句,
不用在每个 case 后面加 break;(go 自动会加)
switch 可以加赋值语句, case 可以写变量
可以省略 switch
*/

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Go runs on ")
	b := "darwin"
	switch os := runtime.GOOS; os {
	case b:
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println(os)
	}
}