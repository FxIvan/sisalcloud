package main

import (
	"fmt"
	"errors"
)

// Structura que lo utilizaremos en las funciones
type Circle struct {
	Radius int
}

//Cada vez que llamemos a este funcion lo aremos como c.Area(), por que utilizamos el structura Circle
func (c Circle) Area() float64 {
	//Decimos que pi es un float64, por que si no lo declaramos, el compilador lo toma como un float32
	var pi float64 = 3.14
	//tenemos que convertir la variable Radius a float64, por que si no lo hacemos, el compilador nos tira un error
	return pi * float64(c.Radius) * float64(c.Radius)
}


//funcion que recibe un int y devuelve un float64
func area(radius int) float64 {
	var pi float64 = 3.14
	return pi*float64(radius)*float64(radius)
}

//func GetX() (x X, err error)

func main() {
	
	err := errors.New("Some Error")
	//Capturamos el error, este debe ser diferente de nil, osino no se ejecuta el if
	if err != nil {
   		fmt.Print(err)
	}

	//Arrays
	var a[10]int 	//	[0 0 0 0 0 0 0 0 0 0]
	a[0] = 1		//	[1 0 0 0 0 0 0 0 0 0]
	a[1] = 2		//	[1 2 0 0 0 0 0 0 0 0]
	//var aSlice []int // nil, es declarada pero no es usada
	fmt.Println(a)

	//MAP 
	//make es una funcion que crea un map, el primer parametro es el tipo de dato de la key, y el segundo es el tipo de dato del value
	var ranks map[string]string = make(map[string]string)
	ranks["Joe"] = "hola"  // set
	ranks["Jane"] = "chau"
	rankOfJoe := ranks["Jane"] // capturamos el valor de la key Jane
	fmt.Println("rankOfJoe",rankOfJoe)
	fmt.Println(ranks)//map[Jane:chau Joe:hola]


	fmt.Println(area(56))


//	myX, err := GetX()
//	if err != nil {
//     	fmt.Printf("error: %s", err)
//	}

	//Este es un ejemplo de como llamamos a la funcion Area() que esta dentro del structura Circle
	var c Circle
	c.Radius = 2
	fmt.Println(c.Area())


	//Funciones anonimas, son funciones que no tienen nombre pero que se pueden ejecutar
	areaSquared := func(radius int) float64 {  
		return float64(radius)*float64(radius)
	}
	//Asi se llama a la funcion anonima
	fmt.Println(areaSquared(2))

}