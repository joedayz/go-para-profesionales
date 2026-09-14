// Clase 01 - Demo 4: Estructuras de control en Go
// Curso: Go para Profesionales | joedayz.pe
//
// Go tiene menos estructuras que otros lenguajes (sin while, sin do-while),
// pero son más expresivas y consistentes.
package main

import "fmt"

// Tipo personalizado para demostrar switch con tipos propios
type EstadoPedido string

const (
	Pendiente  EstadoPedido = "pendiente"
	Procesando EstadoPedido = "procesando"
	Enviado    EstadoPedido = "enviado"
	Cancelado  EstadoPedido = "cancelado"
)

func descripcionEstado(estado EstadoPedido) string {
	switch estado {
	case Pendiente:
		return "⏳ Esperando confirmación de pago"
	case Procesando:
		return "⚙️  Preparando tu pedido"
	case Enviado:
		return "🚚 En camino a tu dirección"
	case Cancelado:
		return "❌ Pedido cancelado"
	default:
		return "🔍 Estado desconocido"
	}
}

func main() {
	// ─────────────────────────────────────────────
	// 1. if / else if / else
	//    Go no necesita paréntesis en la condición
	// ─────────────────────────────────────────────
	nota := 85

	if nota >= 90 {
		fmt.Println("Calificación: A - Excelente")
	} else if nota >= 80 {
		fmt.Println("Calificación: B - Muy bueno")
	} else if nota >= 70 {
		fmt.Println("Calificación: C - Regular")
	} else {
		fmt.Println("Calificación: F - Desaprobado")
	}

	// ─────────────────────────────────────────────
	// 2. if con inicializador (init statement)
	//    La variable declarada sólo vive dentro del if
	//    Patrón muy común con el "ok pattern" de maps/type assertion
	// ─────────────────────────────────────────────
	precios := map[string]float64{"Go": 149.99, "Docker": 99.99}

	if precio, ok := precios["Go"]; ok {
		fmt.Printf("Precio del curso Go: $%.2f\n", precio)
	} else {
		fmt.Println("Curso no encontrado")
	}

	// ─────────────────────────────────────────────
	// 3. for — el único bucle de Go
	//    Reemplaza for, while y do-while de otros lenguajes
	// ─────────────────────────────────────────────

	// Forma 1: clásico (como C)
	fmt.Print("Clásico: ")
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// Forma 2: como while
	fmt.Print("While:   ")
	n := 1
	for n < 32 {
		fmt.Printf("%d ", n)
		n *= 2
	}
	fmt.Println()

	// Forma 3: range sobre slice
	cursos := []string{"Go", "Docker", "Kubernetes", "Kafka"}
	fmt.Println("Cursos disponibles:")
	for i, curso := range cursos {
		fmt.Printf("  [%d] %s\n", i+1, curso)
	}

	// Forma 4: range sobre map (orden no garantizado)
	fmt.Println("Precios:")
	for nombre, precio := range precios {
		fmt.Printf("  %-10s $%.2f\n", nombre, precio)
	}

	// Forma 5: range con _ para ignorar el índice
	fmt.Print("Solo valores: ")
	for _, c := range cursos {
		fmt.Printf("%s ", c)
	}
	fmt.Println()

	// Forma 6: loop infinito con break
	fmt.Print("Fibonacci hasta 100: ")
	a, b := 0, 1
	for {
		if a > 100 {
			break
		}
		fmt.Printf("%d ", a)
		a, b = b, a+b
	}
	fmt.Println()

	// continue: saltar iteración
	fmt.Print("Pares del 1 al 10: ")
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// ─────────────────────────────────────────────
	// 4. switch
	//    En Go NO hace falta break (no hay fallthrough por defecto)
	//    Más expresivo que en Java/C
	// ─────────────────────────────────────────────
	fmt.Println("\n--- Estados de pedido ---")
	estados := []EstadoPedido{Pendiente, Procesando, Enviado, Cancelado}
	for _, e := range estados {
		fmt.Printf("%-12s → %s\n", e, descripcionEstado(e))
	}

	// switch sin expresión (actúa como if-else chain)
	temperatura := 38.5
	fmt.Print("\nDiagnóstico: ")
	switch {
	case temperatura < 36.0:
		fmt.Println("Hipotermia")
	case temperatura <= 37.5:
		fmt.Println("Normal")
	case temperatura <= 39.0:
		fmt.Println("Fiebre leve")
	default:
		fmt.Println("Fiebre alta — consultar médico")
	}

	// ─────────────────────────────────────────────
	// 5. Labels con break (útil en loops anidados)
	//    Equivale al break con etiqueta de otros lenguajes
	// ─────────────────────────────────────────────
	fmt.Println("\n--- Búsqueda en matriz ---")
	matriz := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	objetivo := 5
	encontrado := false

busqueda:
	for i, fila := range matriz {
		for j, val := range fila {
			if val == objetivo {
				fmt.Printf("Encontrado %d en [%d][%d]\n", objetivo, i, j)
				encontrado = true
				break busqueda // sale de ambos loops
			}
		}
	}
	if !encontrado {
		fmt.Println("No encontrado")
	}
}
