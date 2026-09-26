// Clase 02 - Demo 2: Métodos
// Curso: Go para Profesionales | joedayz.pe
//
// Una idea: el receptor (lo que va antes del nombre) decide
// si el método ve una copia o el original.
package main

import (
	"errors"
	"fmt"
)

// Curso sigue siendo un struct. Los métodos se declaran fuera,
// no dentro del tipo. No hay clase que los envuelva.
type Curso struct {
	ID     string
	Titulo string
	Precio float64
	Activo bool
}

// Etiqueta, ConDescuento, AplicarDescuento y NuevoCurso van en mayúscula:
// son la API del tipo. Dentro de main la mayúscula no oculta nada;
// en el demo 5 es lo que otro paquete puede llamar.

// Etiqueta lee el curso. Receptor por valor: recibe una copia y no puede
// cambiar el original. Para leer, la copia alcanza.
func (c Curso) Etiqueta() string {
	return fmt.Sprintf("%s — $%.2f", c.Titulo, c.Precio)
}

// demostrarCopia no es un patrón: es el experimento.
// El descuento cae sobre la copia que recibe el método y se pierde
// al retornar, porque nadie devuelve esa copia.
func (c Curso) demostrarCopia(porcentaje float64) {
	c.Precio = c.Precio * (1 - porcentaje)
}

// ConDescuento sí devuelve la copia ya rebajada.
// El curso sobre el que lo llamas sigue igual.
func (c Curso) ConDescuento(porcentaje float64) (Curso, error) {
	precio, err := precioRebajado(c.Precio, porcentaje)
	if err != nil {
		return Curso{}, err
	}
	c.Precio = precio
	return c, nil
}

// AplicarDescuento usa receptor puntero: c es *Curso, así que
// c.Precio escribe en el original.
func (c *Curso) AplicarDescuento(porcentaje float64) error {
	precio, err := precioRebajado(c.Precio, porcentaje)
	if err != nil {
		return err
	}
	c.Precio = precio
	return nil
}

func precioRebajado(precio, porcentaje float64) (float64, error) {
	if porcentaje < 0 || porcentaje > 1 {
		return 0, errors.New("el porcentaje debe estar entre 0 y 1")
	}
	return precio * (1 - porcentaje), nil
}

// NuevoCurso concentra la validación. Un literal Curso{...} sigue
// pudiendo armar un precio negativo: dentro del mismo paquete los
// campos exportados se ven. Quien respete la regla usa el constructor.
func NuevoCurso(id, titulo string, precio float64) (Curso, error) {
	if titulo == "" {
		return Curso{}, errors.New("el título no puede estar vacío")
	}
	if precio < 0 {
		return Curso{}, errors.New("el precio no puede ser negativo")
	}
	return Curso{
		ID:     id,
		Titulo: titulo,
		Precio: precio,
		Activo: true,
	}, nil
}

func main() {
	// ─────────────────────────────────────────────
	// 1. Constructor
	//    Si el título falta o el precio no tiene sentido,
	//    no llega a existir el curso.
	// ─────────────────────────────────────────────
	base, err := NuevoCurso("go", "Go para profesionales", 200)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Creado:", base.Etiqueta())

	if _, err = NuevoCurso("x", "", 10); err != nil {
		fmt.Println("Título vacío →", err)
	}
	if _, err = NuevoCurso("x", "Kafka", -5); err != nil {
		fmt.Println("Precio negativo →", err)
	}

	// ─────────────────────────────────────────────
	// 2. Receptor por valor, sin devolver nada
	//    El 25% se aplica y se tira. El original sigue en 200.
	// ─────────────────────────────────────────────
	fmt.Printf("\nAntes: $%.2f\n", base.Precio)
	base.demostrarCopia(0.25)
	fmt.Printf("Tras demostrarCopia: $%.2f (la copia se descartó)\n", base.Precio)

	// ─────────────────────────────────────────────
	// 3. Receptor por valor que devuelve otro curso
	//    Patrón útil cuando no quieres mutar.
	// ─────────────────────────────────────────────
	rebajado, err := base.ConDescuento(0.25)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Original: $%.2f · devuelto: %s\n", base.Precio, rebajado.Etiqueta())

	if _, err = base.ConDescuento(1.5); err != nil {
		fmt.Println("Porcentaje inválido →", err)
	}

	// ─────────────────────────────────────────────
	// 4. Receptor puntero
	//    base es una variable, así que Go toma su dirección
	//    aunque aquí no hayas escrito &base.
	//    Un temporal no direccionable (Curso{}.AplicarDescuento)
	//    no compilaría: no hay dónde guardar el cambio.
	// ─────────────────────────────────────────────
	if err = base.AplicarDescuento(0.25); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Tras AplicarDescuento:", base.Etiqueta())
}
