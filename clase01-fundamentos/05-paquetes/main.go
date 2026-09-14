// Clase 01 - Demo 5: Paquetes y estructura del código en Go
// Curso: Go para Profesionales | joedayz.pe
//
// En Go la organización del código se hace con paquetes.
// Cada directorio es un paquete. El import path se construye
// desde la raíz del módulo definido en go.mod.
package main

import (
	"fmt"

	// Import de nuestro paquete local usando el module path del go.mod
	"go-para-profesionales/clase01/05-paquetes/mathutil"
)

func main() {
	// ─────────────────────────────────────────────
	// 1. Usar constante exportada del paquete
	// ─────────────────────────────────────────────
	fmt.Printf("Pi = %.10f\n", mathutil.Pi)

	// ─────────────────────────────────────────────
	// 2. Funciones del paquete mathutil
	// ─────────────────────────────────────────────
	fmt.Println("\n--- Valor absoluto ---")
	fmt.Printf("Abs(-7.5)  = %.1f\n", mathutil.Abs(-7.5))
	fmt.Printf("Abs(3.2)   = %.1f\n", mathutil.Abs(3.2))

	fmt.Println("\n--- Potencia ---")
	resultado, err := mathutil.Potencia(2, 10)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("2^10 = %.0f\n", resultado)
	}

	_, err = mathutil.Potencia(3, -1)
	if err != nil {
		fmt.Println("Error capturado:", err)
	}

	fmt.Println("\n--- Área del círculo ---")
	area, err := mathutil.AreaCirculo(5)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Área con radio=5: %.4f\n", area)
	}

	// ─────────────────────────────────────────────
	// 3. Primos
	// ─────────────────────────────────────────────
	fmt.Println("\n--- Números primos hasta 50 ---")
	primos := mathutil.PrimosHasta(50)
	fmt.Println(primos)

	// ─────────────────────────────────────────────
	// 4. Visibilidad: sólo podemos acceder a lo exportado
	//    Si intentaras: mathutil.esEnteroPositivo(5)
	//    el compilador daría error → encapsulamiento garantizado
	// ─────────────────────────────────────────────
	fmt.Println("\n--- Verificación de primos individuales ---")
	numeros := []int{1, 2, 7, 13, 15, 17, 20, 23}
	for _, n := range numeros {
		if mathutil.EsPrimo(n) {
			fmt.Printf("  %2d → ✅ primo\n", n)
		} else {
			fmt.Printf("  %2d → ❌ no es primo\n", n)
		}
	}

	// ─────────────────────────────────────────────
	// 5. Alias de import (cuando hay conflicto de nombres)
	// ─────────────────────────────────────────────
	// import mu "go-para-profesionales/clase01/05-paquetes/mathutil"
	// Podríamos usar: mu.Pi, mu.Abs(...), etc.
	//
	// import _ "paquete/solo/por/efectos/secundarios"
	// El blank identifier _ en imports registra init() sin usar el paquete.
	fmt.Println("\n✅ Demo de paquetes completado")
}
