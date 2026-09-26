// Clase 02 - Demo 5: Diseño de un servicio
// Curso: Go para Profesionales | joedayz.pe
//
// Una idea: el servicio pide un comportamiento.
// main decide qué tipo concreto lo cumple.
package main

import (
	"errors"
	"fmt"

	"go-para-profesionales/clase02/05-diseno/catalogo"
	"go-para-profesionales/clase02/05-diseno/curso"
	"go-para-profesionales/clase02/05-diseno/memoria"
)

// Estas líneas no se ejecutan. Si Memoria o repositorioRoto pierden
// un método, o cambian una firma, este archivo deja de compilar.
var (
	_ catalogo.Repositorio = (*memoria.Memoria)(nil)
	_ catalogo.Repositorio = repositorioRoto{}
)

// repositorioRoto es el segundo almacén. Vive en main a propósito:
// el paquete catalogo no lo conoce y no hace falta que lo conozca.
type repositorioRoto struct{}

func (repositorioRoto) Guardar(curso.Curso) error {
	return errors.New("disco lleno")
}

func (repositorioRoto) Buscar(string) (curso.Curso, error) {
	return curso.Curso{}, curso.ErrNoEncontrado
}

func (repositorioRoto) Listar() ([]curso.Curso, error) {
	return nil, errors.New("disco lleno")
}

func main() {
	// ─────────────────────────────────────────────
	// 1. Armar el servicio
	//    memoria.Nuevo() devuelve *memoria.Memoria.
	//    catalogo.Nuevo pide un catalogo.Repositorio.
	//    El puntero entra porque tiene los tres métodos.
	// ─────────────────────────────────────────────
	svc, err := catalogo.Nuevo(memoria.Nuevo())
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("--- Publicar ---")
	for _, item := range []struct {
		titulo string
		precio float64
	}{
		{"Go para profesionales", 149.90},
		{"Docker de cero", 99},
	} {
		publicado, err := svc.Publicar(item.titulo, item.precio)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println(publicado.Etiqueta())
	}

	// ─────────────────────────────────────────────
	// 2. Reglas del servicio
	//    Estas fallas no llegan al almacén.
	//    errors.Is compara con el valor centinela, no con el texto.
	// ─────────────────────────────────────────────
	fmt.Println("\n--- Reglas ---")
	_, err = svc.Publicar("   ", 50)
	if errors.Is(err, catalogo.ErrTituloVacio) {
		fmt.Println("Título vacío →", err)
	}

	_, err = svc.Publicar("Kafka", -10)
	if errors.Is(err, catalogo.ErrPrecioInvalido) {
		fmt.Println("Precio negativo →", err)
	}

	// Cero es válido: curso gratuito. No es el caso negativo.
	gratis, err := svc.Publicar("Charla de bienvenida", 0)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Gratuito →", gratis.Etiqueta())

	// ─────────────────────────────────────────────
	// 3. Lo que responde el almacén
	// ─────────────────────────────────────────────
	fmt.Println("\n--- Consultas ---")
	encontrado, err := svc.Detalle("CUR-001")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Detalle:", encontrado.Etiqueta())

	_, err = svc.Detalle("CUR-999")
	if errors.Is(err, curso.ErrNoEncontrado) {
		fmt.Println("Detalle CUR-999 →", err)
	}

	fmt.Println("\nListado:")
	cursos, err := svc.Listar()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for _, c := range cursos {
		fmt.Println(" ", c.Etiqueta())
	}

	// ─────────────────────────────────────────────
	// 4. Cambiar el almacén, no el servicio
	//    Mismo Publicar. Otro Repositorio.
	//    El error del disco llega envuelto y el texto lo dice.
	// ─────────────────────────────────────────────
	fmt.Println("\n--- Otro almacén, el mismo servicio ---")
	roto, err := catalogo.Nuevo(repositorioRoto{})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	_, err = roto.Publicar("Kubernetes", 189)
	if err != nil {
		fmt.Println("Publicar →", err)
	}
}
