package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Definimos una constante para el sexo por defecto
const sexoPorDefecto = "H"

// Estructura de la clase Persona
type Persona struct {
	nombre string
	edad   int
	DNI    string
	sexo   string
	peso   float64
	altura float64
}

// Constructor por defecto
func NewPersonaDefecto() *Persona {
	return &Persona{
		DNI:  generarDNI(),
		sexo: sexoPorDefecto,
	}
}

// Constructor con nombre, edad y sexo, el resto por defecto
func NewPersona(nombre string, edad int, sexo string) *Persona {
	return &Persona{
		nombre: nombre,
		edad:   edad,
		DNI:    generarDNI(),
		sexo:   comprobarSexo(sexo),
	}
}

// Constructor con todos los atributos como parámetro
func NewPersonaCompleto(nombre string, edad int, DNI string, sexo string, peso, altura float64) *Persona {
	return &Persona{
		nombre: nombre,
		edad:   edad,
		DNI:    DNI,
		sexo:   comprobarSexo(sexo),
		peso:   peso,
		altura: altura,
	}
}

// Método para calcular el IMC (Índice de Masa Corporal)
func (p *Persona) calcularIMC() int {
	imc := p.peso / (math.Pow(p.altura, 2))
	switch {
	case imc < 20:
		return -1 // Por debajo del peso ideal
	case imc >= 20 && imc <= 25:
		return 0 // Peso ideal
	default:
		return 1 // Sobrepeso
	}
}

// Método para verificar si es mayor de edad
func (p *Persona) esMayorDeEdad() bool {
	return p.edad >= 18
}

// Método toString para mostrar la información del objeto
func (p *Persona) toString() string {
	return fmt.Sprintf("Nombre: %s\nEdad: %d\nDNI: %s\nSexo: %s\nPeso: %.2f\nAltura: %.2f", p.nombre, p.edad, p.DNI, p.sexo, p.peso, p.altura)
}

// Método privado para generar un DNI aleatorio
func generarDNI() string {
	rand.Seed(time.Now().UnixNano())
	dni := ""
	for i := 0; i < 8; i++ {
		dni += fmt.Sprintf("%d", rand.Intn(10))
	}
	return dni
}

// Método privado para comprobar el sexo
func comprobarSexo(sexo string) string {
	if sexo == "H" || sexo == "M" {
		return sexo
	}
	return sexoPorDefecto
}

func main() {
	// Solicitar datos al usuario
	var nombre, sexo string
	var edad int
	var peso, altura float64

	fmt.Print("Ingrese el nombre: ")
	fmt.Scan(&nombre)

	fmt.Print("Ingrese la edad: ")
	fmt.Scan(&edad)

	fmt.Print("Ingrese el sexo (H/M): ")
	fmt.Scan(&sexo)

	fmt.Print("Ingrese el peso (en kg): ")
	fmt.Scan(&peso)

	fmt.Print("Ingrese la altura (en metros): ")
	fmt.Scan(&altura)

	// Crear objetos Persona
	persona1 := NewPersona(nombre, edad, sexo)
	persona2 := NewPersonaCompleto(nombre, edad, generarDNI(), sexo, peso, altura)
	persona3 := NewPersonaDefecto()

	// Calcular IMC y verificar si es mayor de edad para cada persona
	resultado1 := persona1.calcularIMC()
	resultado2 := persona2.calcularIMC()
	resultado3 := persona3.calcularIMC()

	mayorDeEdad1 := persona1.esMayorDeEdad()
	mayorDeEdad2 := persona2.esMayorDeEdad()
	mayorDeEdad3 := persona3.esMayorDeEdad()

	// Mostrar información de cada persona
	fmt.Println("\nInformación de la Persona 1:")
	fmt.Println(persona1.toString())
	fmt.Printf("Resultado IMC: %d\n", resultado1)
	fmt.Printf("Mayor de edad: %v\n", mayorDeEdad1)

	fmt.Println("\nInformación de la Persona 2:")
	fmt.Println(persona2.toString())
	fmt.Printf("Resultado IMC: %d\n", resultado2)
	fmt.Printf("Mayor de edad: %v\n", mayorDeEdad2)

	fmt.Println("\nInformación de la Persona 3:")
	fmt.Println(persona3.toString())
	fmt.Printf("Resultado IMC: %d\n", resultado3)
	fmt.Printf("Mayor de edad: %v\n", mayorDeEdad3)
}
