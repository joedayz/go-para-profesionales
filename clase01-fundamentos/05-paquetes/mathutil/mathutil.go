// Package mathutil demuestra cómo crear un paquete reutilizable en Go.
// Curso: Go para Profesionales | joedayz.pe
//
// Convenciones de Go:
//   - Identificadores en MAYÚSCULA → exportados (públicos)
//   - Identificadores en minúscula → no exportados (privados al paquete)
//   - El nombre del paquete coincide con el directorio
package mathutil

import "errors"

// ─────────────────────────────────────────────────────────────
// Constantes exportadas
// ─────────────────────────────────────────────────────────────

// Pi es la constante π con precisión de float64.
const Pi = 3.141592653589793

// ─────────────────────────────────────────────────────────────
// Funciones exportadas (mayúscula → visibles desde otros paquetes)
// ─────────────────────────────────────────────────────────────

// Abs retorna el valor absoluto de n.
func Abs(n float64) float64 {
	if n < 0 {
		return -n
	}
	return n
}

// Potencia calcula base^exp para exponentes enteros no negativos.
func Potencia(base float64, exp int) (float64, error) {
	if exp < 0 {
		return 0, errors.New("el exponente no puede ser negativo")
	}
	resultado := 1.0
	for i := 0; i < exp; i++ {
		resultado *= base
	}
	return resultado, nil
}

// AreaCirculo calcula el área de un círculo dado su radio.
func AreaCirculo(radio float64) (float64, error) {
	if radio < 0 {
		return 0, errors.New("el radio no puede ser negativo")
	}
	return Pi * radio * radio, nil
}

// EsPrimo determina si n es un número primo.
func EsPrimo(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// ─────────────────────────────────────────────────────────────
// Función no exportada (minúscula → privada al paquete)
// ─────────────────────────────────────────────────────────────

// esEnteroPositivo es una helper interna, no accesible desde fuera.
func esEnteroPositivo(n int) bool {
	return n > 0
}

// PrimosHasta retorna todos los números primos ≤ limite.
func PrimosHasta(limite int) []int {
	if !esEnteroPositivo(limite) {
		return nil
	}
	primos := []int{}
	for i := 2; i <= limite; i++ {
		if EsPrimo(i) {
			primos = append(primos, i)
		}
	}
	return primos
}
