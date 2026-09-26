// Clase 02 - Demo 1: Structs
// Curso: Go para Profesionales | joedayz.pe
//
// Una idea: un struct agrupa datos y se copia entero al asignarlo.
// En Java o C# una clase vive detrás de una referencia.
// En Go el struct es el valor.
package main

import "fmt"

// Curso es un dato con forma. Todavía no tiene métodos ni reglas:
// eso entra en los demos siguientes.
type Curso struct {
	ID     string
	Titulo string
	Precio float64
	Activo bool
}

func main() {
	// ─────────────────────────────────────────────
	// 1. Zero value
	//    Declarar sin literal deja cada campo en su cero.
	//    string → "", números → 0, bool → false.
	//    No hay null: el struct siempre existe.
	// ─────────────────────────────────────────────
	var vacio Curso
	fmt.Printf("Zero value: %#v\n", vacio)

	// ─────────────────────────────────────────────
	// 2. Literal
	//    Los campos se nombran. El orden no importa.
	//    Lo que omites se queda en cero: este borrador
	//    nace inactivo y a precio 0.
	// ─────────────────────────────────────────────
	borrador := Curso{Titulo: "Kafka en la práctica"}
	fmt.Printf("Borrador:   %#v\n", borrador)

	goBasico := Curso{
		ID:     "go-basico",
		Titulo: "Go para profesionales",
		Precio: 149.90,
		Activo: true,
	}
	fmt.Printf("Publicado:  %s · $%.2f\n", goBasico.Titulo, goBasico.Precio)

	// ─────────────────────────────────────────────
	// 3. Comparación
	//    Si todos los campos se pueden comparar, el struct
	//    también. No hace falta escribir Equals.
	// ─────────────────────────────────────────────
	igual := Curso{
		ID:     "go-basico",
		Titulo: "Go para profesionales",
		Precio: 149.90,
		Activo: true,
	}
	fmt.Println("Mismo contenido:", goBasico == igual)

	// ─────────────────────────────────────────────
	// 4. Asignar copia
	//    copia y goBasico son dos valores distintos.
	//    Cambiar el precio de la copia no toca el original.
	// ─────────────────────────────────────────────
	copia := goBasico
	copia.Precio = 99
	fmt.Printf("Original: $%.2f · copia: $%.2f\n", goBasico.Precio, copia.Precio)

	// ─────────────────────────────────────────────
	// 5. Slice de structs
	//    Meter goBasico en el slice copia el valor de ese momento.
	// ─────────────────────────────────────────────
	catalogo := []Curso{
		goBasico,
		{ID: "docker", Titulo: "Docker de cero", Precio: 99, Activo: true},
	}
	fmt.Println("Catálogo:")
	for _, c := range catalogo {
		fmt.Printf("  %-28s $%.2f\n", c.Titulo, c.Precio)
	}

	// ─────────────────────────────────────────────
	// 6. Puntero
	//    & toma la dirección. ptr.Precio es lo mismo que
	//    (*ptr).Precio: Go desreferencia el campo solo.
	//    Cambia el original. El elemento del slice no: era una copia.
	// ─────────────────────────────────────────────
	ptr := &goBasico
	ptr.Precio = 129.90
	fmt.Printf("Vía puntero, el original queda en $%.2f\n", goBasico.Precio)
	fmt.Printf("El del catálogo sigue en $%.2f\n", catalogo[0].Precio)
}
