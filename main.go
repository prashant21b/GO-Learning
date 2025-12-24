package main

import "fmt"
func evenodd(num int){
	if num%2==0{
		fmt.Println("Even")
	}else{
		fmt.Println("Odd")
	}
}
func main() {
	fmt.Println("Enter a number")
	num:=0
	fmt.Scan(&num)
	evenodd(num)
}
