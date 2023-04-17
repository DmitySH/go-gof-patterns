package main

import (
	"fmt"
	"github.com/DmitySH/go-gof-patterns/Singleton/singletonwarehouse"
)

func main() {
	w1 := singletonwarehouse.GetWarehouse("New Riga", "Moscow Oblast'")
	fmt.Println(w1)

	w2 := singletonwarehouse.GetWarehouse("OOO packers", "Khimki")
	fmt.Println(w2)

	fmt.Println(w1 == w2)
}
