// Clase 01 - Demo 2: Tipos de datos en Go
// Curso: Go para Profesionales | joedayz.pe
package main

import "fmt"

func main() {
	// ─────────────────────────────────────────────
	// 1. Tipos numéricos
	//    Go diferencia tamaños: int8/16/32/64, float32/64
	//    En la práctica: int y float64 son los más usados
	// ─────────────────────────────────────────────
	var entero int = 42
	var grande int64 = 9_000_000_000 // _ como separador visual (Go 1.13+)
	var decimal float64 = 3.14159
	var complejo complex128 = 2 + 3i

	fmt.Printf("int: %d\n", entero)
	fmt.Printf("int64: %d\n", grande)
	fmt.Printf("float64: %.5f\n", decimal)
	fmt.Printf("complex128: %v\n", complejo)

	// ─────────────────────────────────────────────
	// 2. string y rune
	//    string = secuencia de bytes UTF-8
	//    rune   = alias de int32, representa un código Unicode
	// ─────────────────────────────────────────────
	saludo := "¡Hola mundo!"
	fmt.Printf("String: %q\n", saludo)
	fmt.Printf("Longitud en bytes: %d\n", len(saludo))

	// Iterar sobre runas (caracteres Unicode correctamente)
	fmt.Print("Runas: ")
	for _, r := range saludo {
		fmt.Printf("%c", r)
	}
	fmt.Println()

	// ─────────────────────────────────────────────
	// 3. bool
	// ─────────────────────────────────────────────
	esProduccion := true
	fmt.Printf("¿Producción? %t\n", esProduccion)

	// ─────────────────────────────────────────────
	// 4. Arrays: tamaño fijo, parte del tipo
	//    Raramente se usan directo; se prefieren slices
	// ─────────────────────────────────────────────
	var dias [7]string
	dias[0] = "Lunes"
	dias[6] = "Domingo"
	fmt.Printf("Array de días: %v\n", dias)

	// ─────────────────────────────────────────────
	// 5. Slices: la estructura más usada en Go
	//    Abstracción dinámica sobre un array
	// ─────────────────────────────────────────────
	lenguajes := []string{"Go", "Java", "Python"}
	lenguajes = append(lenguajes, "Rust")
	fmt.Printf("Slice: %v | len=%d cap=%d\n", lenguajes, len(lenguajes), cap(lenguajes))

	// Slice de slice (sub-slice)
	primerosDos := lenguajes[0:2]
	fmt.Printf("Sub-slice: %v\n", primerosDos)

	// ─────────────────────────────────────────────
	// 6. Maps: diccionario clave→valor
	// ─────────────────────────────────────────────
	precios := map[string]float64{
		"Go":     149.99,
		"Docker": 99.99,
		"Kafka":  199.99,
	}
	precios["Kubernetes"] = 249.99

	fmt.Printf("Precios: %v\n", precios)

	// Acceso seguro: ok pattern (evita panics)
	precio, ok := precios["Go"]
	if ok {
		fmt.Printf("Precio de Go: $%.2f\n", precio)
	}

	// Eliminar una clave
	delete(precios, "Docker")
	fmt.Printf("Tras delete: %v\n", precios)

	// ─────────────────────────────────────────────
	// 7. Zero values (concepto clave de Go)
	//    Toda variable tiene un valor inicial definido
	// ─────────────────────────────────────────────
	var i int
	var f float64
	var s string
	var b bool
	fmt.Printf("Zero values → int:%d float64:%f string:%q bool:%t\n", i, f, s, b)
}
