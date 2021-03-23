package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	for {
		if !getData() {
			break
		}
	}
}

func getData() bool {
	persona := NewByDefault()
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Introduzca la siguiente informacion: ")
	fmt.Print("Nombre: ")
	name, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("\nInserte correctamente el nombre")
	}

	name = strings.Trim(name, "\n")
	persona.SetNombre(name)

	fmt.Print("Edad: ")
	edad, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("\nInserte correctamente la edad")
	}

	edadInt, err := strconv.Atoi(strings.Trim(edad, "\n"))

	if err != nil {
		fmt.Println("\nSu edad debe ser un numero entero")
	}

	persona.SetEdad(edadInt)

	fmt.Print("Sexo (H/M): ")
	sexo, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("\nInserte correctamente el sexo")
	}

	sexo = strings.Trim(sexo, "\n")
	if sexo == "H" {
		persona.SetSexo(HOMBRE)
	} else if sexo == "M" {
		persona.SetSexo(MUJER)
	} else {
		fmt.Println("\nInserte correctamente el sexo")
	}

	fmt.Print("Peso (kg): ")
	peso, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("\nInserte correctamente su peso")
	}

	peso = strings.Trim(peso, "\n")
	pesoFloat, err := strconv.ParseFloat(peso, 64)

	if err != nil {
		fmt.Println("\nSu peso debe ser un flotante")
	}

	persona.SetPeso(pesoFloat)

	fmt.Print("Altura (m): ")
	altura, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("\nInserte correctamente su altura")
	}

	altura = strings.Trim(altura, "\n")
	alturaFloat, err := strconv.ParseFloat(altura, 64)

	if err != nil {
		fmt.Println("\nSu peso debe ser un flotante")
	}

	persona.SetAltura(alturaFloat)

	persona.ShowStatistics()

	fmt.Println("\nDesea ingresar otra persona? (S/N): ")
	res, _ := reader.ReadByte()

	if res == 'S' || res == 's' {
		return true
	}

	return false
}
