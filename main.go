package main

import "fmt"
func evenodd(num int){
	if num%2==0{
		fmt.Println("Even")
	}else{
		fmt.Println("Odd")
	}
}
func sumN(num int){
	var sum int=0
     for i:=1;i<=num;i++{
	    sum+=i
	}
	fmt.Println(sum)
}

func mulTable(num int){
	 for i:=1;i<=10;i++{
		fmt.Println(num*i)
	}
}
// func largest(num1 int,num2 int ,num3 int) int {
// 	if num1>num2{
// 		if num1>num3{
//           return num1
// 		}
// 	}
// 	else{
// 		if num2>num3{
// 			return num2
// 		}			
// 		return num3	
// 	}
// }
type USER struct{
	Name string
	Age int

}
type PRODUCT struct{
	Name string
	Price int
	Quantity int
}

func updateStock(product *PRODUCT){
	product.Quantity=product.Quantity-1
}

func main() {
	// fmt.Println("Enter a number")
	// num:=0
	// fmt.Scan(&num)
	// evenodd(num)
    sumN(10)
    mulTable(5)
   // fmt.Println(largest(10,20,30))
	var user USER
	user.Name="Prashant"
	user.Age=21
	fmt.Println(user)
	product:=PRODUCT{
		Name:"Laptop",
		Price:100000,
		Quantity:10,
	}
	fmt.Println(product)
	updateStock(&product)
	fmt.Println(product)
}
