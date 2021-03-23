package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Persona struct {
	nombre string
	edad   int
	nss    string
	sexo   Sexo
	peso   float64
	altura float64
}

func (p *Persona) ShowStatistics() {
	p.showWeight()
	p.showIsLegalAge()
	fmt.Println(p.ToString())
}

func (p *Persona) showIsLegalAge() {
	if p.EsMayorDeEdad() {
		fmt.Println("Es mayor de edad")
	} else {
		fmt.Println("Es menor de edad")
	}
}

func (p *Persona) showWeight() {
	imc := p.CalcularIMC()
	var str string

	if imc == FALTA_DE_PESO {
		str = "Tiene falta de peso"
	} else if imc == PESO_NORMAL {
		str = "Su peso es normal"
	} else {
		str = "Tiene sobrepeso"
	}

	fmt.Println(str)
}

func (p *Persona) CalcularIMC() IMC {
	imc := p.peso / math.Pow(p.altura, 2)

	if p.sexo == HOMBRE {
		if imc < 20 {
			return FALTA_DE_PESO
		} else if imc >= 20 && imc <= 25 {
			return PESO_NORMAL
		} else {
			return SOBREPESO
		}
	} else {
		if imc < 19 {
			return FALTA_DE_PESO
		} else if imc >= 19 && imc <= 24 {
			return PESO_NORMAL
		} else {
			return SOBREPESO
		}
	}
}

func (p *Persona) EsMayorDeEdad() bool {
	if p.edad >= 18 {
		return true
	}

	return false
}

func (p *Persona) comprobarSexo(sexo byte) bool {
	if byte(p.sexo) == sexo {
		return true
	}

	return false
}

func (p *Persona) ToString() string {
	return fmt.Sprintf(`
	PERSONA(
		nombre: %s,
		edad: %d,
		nss: %s,
		sexo: %s,
		peso: %f,
		altura: %f
	)
	`, p.nombre, p.edad, p.nss, string(p.sexo), p.peso, p.altura)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func (p *Persona) generaNSS() string {
	rand.Seed(time.Now().UnixNano())

	var nss = make([]byte, 8)
	for i := 0; i < 8; i++ {
		nss[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(nss)
}

func (p *Persona) SetNombre(nombre string) {
	p.nombre = nombre
}

func (p *Persona) SetEdad(edad int) {
	p.edad = edad
}

func (p *Persona) SetSexo(sexo Sexo) {
	p.sexo = sexo
}

func (p *Persona) SetPeso(peso float64) {
	p.peso = peso
}

func (p *Persona) SetAltura(altura float64) {
	p.altura = altura
}

func New(nombre string, edad int, nss string, sexo Sexo, peso, altura float64) *Persona {
	return &Persona{
		nombre,
		edad,
		nss,
		sexo,
		peso,
		altura,
	}
}

func NewByDefault() *Persona {
	p := &Persona{
		nombre: "",
		edad:   0,
		sexo:   HOMBRE,
		peso:   0.0,
		altura: 0.0,
	}

	p.nss = p.generaNSS()
	return p
}
