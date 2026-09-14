// Clase 01 - Demo 3: Funciones en Go
// Curso: Go para Profesionales | joedayz.pe
package main

import (
	"errors"
	"fmt"
)

// ─────────────────────────────────────────────────────────────
// 1. Función básica: un parámetro, un retorno
// ─────────────────────────────────────────────────────────────
func saludar(nombre string) string {
	return fmt.Sprintf("¡Hola, %s!", nombre)
}

// ─────────────────────────────────────────────────────────────
// 2. Múltiples retornos (idioma Go más característico)
//    Convención: el último retorno suele ser un error
// ─────────────────────────────────────────────────────────────
func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("división por cero no permitida")
	}
	return a / b, nil
}

// ─────────────────────────────────────────────────────────────
// 3. Retornos nombrados
//    Útil en funciones cortas para mayor claridad
// ─────────────────────────────────────────────────────────────
func minMax(nums []int) (min, max int) {
	min, max = nums[0], nums[0]
	for _, n := range nums[1:] {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return // "naked return" devuelve los valores nombrados
}

// ─────────────────────────────────────────────────────────────
// 4. Variadic: cantidad variable de argumentos
// ─────────────────────────────────────────────────────────────
func sumar(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// ─────────────────────────────────────────────────────────────
// 5. Funciones como valores (first-class functions)
//    Go trata las funciones como cualquier otro tipo
// ─────────────────────────────────────────────────────────────
func aplicar(valor int, operacion func(int) int) int {
	return operacion(valor)
}

// ─────────────────────────────────────────────────────────────
// 6. Closure: función que captura su entorno
//    Patrón muy usado para counters, builders, etc.
// ─────────────────────────────────────────────────────────────
func nuevoContador() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// ─────────────────────────────────────────────────────────────
// 7. defer: pospone ejecución hasta que la función retorna
//    Orden LIFO — muy usado para liberar recursos
// ─────────────────────────────────────────────────────────────
func procesarArchivo(nombre string) {
	fmt.Printf("Abriendo archivo: %s\n", nombre)
	defer fmt.Printf("Cerrando archivo: %s\n", nombre) // se ejecuta al final

	fmt.Println("  Procesando contenido...")
	fmt.Println("  Validando datos...")
}

func main() {
	// 1. Función básica
	fmt.Println(saludar("Joe"))

	// 2. Múltiples retornos con manejo de error
	resultado, err := dividir(10, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 3 = %.4f\n", resultado)
	}

	_, err = dividir(5, 0)
	if err != nil {
		fmt.Println("Error capturado:", err)
	}

	// 3. Retornos nombrados
	numeros := []int{4, 2, 9, 1, 7, 3}
	min, max := minMax(numeros)
	fmt.Printf("Números: %v → min=%d, max=%d\n", numeros, min, max)

	// 4. Variadic
	fmt.Println("Suma variadic:", sumar(1, 2, 3, 4, 5))

	// Expandir un slice con ... (spread operator)
	vals := []int{10, 20, 30}
	fmt.Println("Suma desde slice:", sumar(vals...))

	// 5. Funciones como valores
	doble := func(n int) int { return n * 2 }
	triple := func(n int) int { return n * 3 }
	fmt.Println("Doble de 7:", aplicar(7, doble))
	fmt.Println("Triple de 7:", aplicar(7, triple))

	// Función anónima inline
	fmt.Println("Cuadrado de 5:", aplicar(5, func(n int) int { return n * n }))

	// 6. Closure
	contador := nuevoContador()
	fmt.Println("Contador:", contador(), contador(), contador()) // 1 2 3

	// Cada closure tiene su propio estado
	otroContador := nuevoContador()
	fmt.Println("Otro contador:", otroContador()) // 1 (independiente)

	// 7. defer - orden de ejecución
	procesarArchivo("datos.csv")

	// Defer múltiple → se ejecutan en orden LIFO
	fmt.Println("\n--- Orden LIFO de defer ---")
	for i := 1; i <= 3; i++ {
		defer fmt.Printf("defer #%d ejecutado\n", i)
	}
	fmt.Println("Fin de main()")
}
