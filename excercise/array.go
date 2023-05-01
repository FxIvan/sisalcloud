package main

import "fmt"

func main(){
	//[5]int -> 5 elementos dentro del array
	var array = new([5]int);
	fmt.Println(array);

	i := 1

	for i < 5{
		fmt.Println(i)
		i = i + 1

	}

	for j := 1;  j < 10; j++{
		if j == 2 {
			fmt.Println("Es 2")
		}
	}

var a  [5]int
a[2] = 8
a[3] = 1
a[4] = 10

var e string

for a,b := range a{
	if b == 1{
		e = "Es 1"
	}else if b == 8{
		e = "Es 8"
	}else if b == 10{

		e = "Es 10"
	}
	fmt.Println(a,b,e)
}
}

