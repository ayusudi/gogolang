package main

import (
	"01-task-week-1/add"
	"01-task-week-1/divide"
	"01-task-week-1/variadic_reduction"
	"fmt"
)

func main() {
    addResult := add.Add(5, 5)
    fmt.Println("Add(5, 5) =", addResult)
    
    bagiResult := divide.Bagi(10, 5)
    fmt.Println("Bagi(10, 5) =", bagiResult)
    
    kurangResult := variadic_reduction.KurangVariadic(40, 20, 10, 5)
    fmt.Println("KurangVariadic(40, 20, 10, 5) =", kurangResult)
}