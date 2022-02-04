package main

import (
	"fmt"
	"strings"
	"time"
)

func main()  {
	go caracterPorSegundo("Hola mundo")
	fmt.Println("Estoy aqui")
	var x string
	fmt.Scanln(&x) //Pide por pantalla que introduzca un valor
}

func caracterPorSegundo(nombre string)  {
	letras := strings.Split(nombre, "") //Separa cada caracter de nombre con ""
	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond) //Espera un segundo
		fmt.Println(letra)
	}
}
