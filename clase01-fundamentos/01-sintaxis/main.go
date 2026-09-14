// Clase 01 - Demo 1: Sintaxis básica de Go
// Curso: Go para Profesionales | joedayz.pe
package main

import "fmt"

func main() {
	// ─────────────────────────────────────────────
	// 1. Hola mundo - punto de entrada siempre es main()
	// ─────────────────────────────────────────────
	fmt.Println("¡Hola, Go para Profesionales!")

	// ─────────────────────────────────────────────
	// 2. Declaración de variables
	//    Go es tipado estático, pero infiere el tipo
	// ─────────────────────────────────────────────

	// Forma 1: declaración explícita
	var nombre string = "Joe"
	var edad int = 35

	// Forma 2: inferencia de tipo (más común)
	pais := "Perú"

	// Forma 3: var sin valor → usa el zero value del tipo
	var saldo float64 // zero value = 0.0
	var activo bool   // zero value = false

	fmt.Println("Nombre:", nombre)
	fmt.Println("Edad:", edad)
	fmt.Println("País:", pais)
	fmt.Println("Saldo inicial:", saldo)
	fmt.Println("Activo:", activo)

	// ─────────────────────────────────────────────
	// 3. Constantes
	//    Se evalúan en tiempo de compilación
	// ─────────────────────────────────────────────
	const version = "1.0.0"
	const maxConexiones = 100

	fmt.Printf("Versión: %s | Max conexiones: %d\n", version, maxConexiones)

	// ─────────────────────────────────────────────
	// 4. Múltiple asignación y swap idiomático
	// ─────────────────────────────────────────────
	x, y := 10, 20
	fmt.Printf("Antes del swap: x=%d, y=%d\n", x, y)
	x, y = y, x
	fmt.Printf("Después del swap: x=%d, y=%d\n", x, y)

	// ─────────────────────────────────────────────
	// 5. fmt: los verbos más usados en producción
	// ─────────────────────────────────────────────
	precio := 99.95
	fmt.Printf("Producto: %-15s Precio: $%.2f\n", "Curso Go", precio)
	fmt.Printf("Tipo de precio: %T\n", precio) // %T → imprime el tipo
}
