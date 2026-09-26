// Clase 02 - Demo 4: Interfaces
// Curso: Go para Profesionales | joedayz.pe
//
// Una idea: una interfaz es una lista de métodos.
// Quien tiene esos métodos la cumple, aunque no lo declare.
package main

import "fmt"

type Curso struct {
	ID     string
	Titulo string
}

// Memoria guarda cursos en un map. No menciona ninguna interfaz.
type Memoria struct {
	cursos map[string]Curso
}

func (m Memoria) Buscar(id string) (Curso, bool) {
	curso, ok := m.cursos[id]
	return curso, ok
}

// CatalogoFijo tampoco. Misma firma, otro comportamiento.
// El receptor no se nombra porque este tipo no guarda estado.
type CatalogoFijo struct{}

func (CatalogoFijo) Buscar(id string) (Curso, bool) {
	if id == "go" {
		return Curso{ID: "go", Titulo: "Go para profesionales"}, true
	}
	return Curso{}, false
}

// Buscador aparece cuando hace falta una función que sirva
// para los dos. La declara quien la usa, no quien guarda los datos.
//
// Un solo método. Mientras más métodos pidas, menos tipos caben.
type Buscador interface {
	Buscar(id string) (Curso, bool)
}

// describir no sabe si detrás hay un map o un valor fijo.
// Solo sabe que puede llamar Buscar.
func describir(b Buscador, id string) {
	curso, ok := b.Buscar(id)
	if !ok {
		fmt.Printf("  %-8s → no está\n", id)
		return
	}
	fmt.Printf("  %-8s → %s\n", id, curso.Titulo)
}

func explicarOrigen(b Buscador) {
	// El type switch recupera el tipo concreto cuando de verdad
	// te hace falta. Dentro de describir no hizo falta.
	switch origen := b.(type) {
	case Memoria:
		fmt.Printf("  memoria con %d cursos\n", len(origen.cursos))
	case CatalogoFijo:
		fmt.Println("  catálogo fijo, sin estado")
	default:
		fmt.Printf("  otro tipo: %T\n", origen)
	}
}

func main() {
	memoria := Memoria{cursos: map[string]Curso{
		"go":     {ID: "go", Titulo: "Go para profesionales"},
		"docker": {ID: "docker", Titulo: "Docker de cero"},
	}}
	fijo := CatalogoFijo{}

	// ─────────────────────────────────────────────
	// 1. Los dos entran donde piden un Buscador
	//    Nadie escribió "implements Buscador".
	// ─────────────────────────────────────────────
	fmt.Println("Desde memoria:")
	describir(memoria, "go")
	describir(memoria, "docker")

	fmt.Println("Desde catálogo fijo:")
	describir(fijo, "go")
	describir(fijo, "docker")

	// ─────────────────────────────────────────────
	// 2. Una lista mixta
	//    El slice guarda la interfaz, no un tipo concreto.
	// ─────────────────────────────────────────────
	fmt.Println("\nMisma llamada, distintos orígenes:")
	buscadores := []Buscador{memoria, fijo}
	for _, b := range buscadores {
		describir(b, "docker")
		explicarOrigen(b)
	}

	// ─────────────────────────────────────────────
	// 3. Afirmación con ok
	//    b.(Memoria) sin ok entra en pánico si el tipo no coincide.
	//    Con ok, el programa sigue. Es el mismo patrón de los maps.
	// ─────────────────────────────────────────────
	var b Buscador = fijo
	if _, ok := b.(Memoria); !ok {
		fmt.Println("\nEse buscador no es la memoria")
	}

	// ─────────────────────────────────────────────
	// 4. Nil con trampa
	//    Una interfaz guarda (tipo, valor).
	//    Vacía: tipo nil y valor nil. Comparar con nil da true.
	//    Con un puntero nulo adentro: el tipo SÍ está, el valor no.
	//    Comparar con nil da false. Llamar al método entraría en pánico.
	// ─────────────────────────────────────────────
	fmt.Println("\n--- Nil ---")
	var sinTipo Buscador
	fmt.Println("1) sin tipo ni valor:", sinTipo == nil)

	var memoriaNula *Memoria
	conTipoNulo := Buscador(memoriaNula)
	fmt.Println("2) con tipo, valor nulo:", conTipoNulo == nil)
	fmt.Printf("   dentro hay: %T %v\n", conTipoNulo, conTipoNulo)

	// Buscar de este demo devuelve bool, como el ok de un map.
	// En el demo 5 el almacén devuelve error: no encontrar un curso
	// y que falle el disco no son la misma situación.
}
