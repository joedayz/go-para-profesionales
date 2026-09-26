// Package memoria guarda cursos en un slice, en el orden en que llegan.
//
// No importa al servicio. No nombra la interfaz que va a cumplir.
// Le basta con tener Guardar, Buscar y Listar sobre curso.Curso.
package memoria

import "go-para-profesionales/clase02/05-diseno/curso"

// Memoria es un almacén de juguete. En producción esto sería Postgres
// u otro proceso; el servicio no tendría que enterarse.
type Memoria struct {
	cursos []curso.Curso
}

// Nuevo devuelve un puntero porque Guardar modifica el slice.
// El receptor de los métodos es *Memoria: un Memoria por valor
// no tendría estos métodos en su method set y no cumpliría la interfaz.
func Nuevo() *Memoria {
	return &Memoria{}
}

// Guardar agrega una copia del curso al final.
// Devuelve error aunque aquí siempre sea nil: un disco real puede fallar,
// y la firma tiene que dejar ese hueco.
func (m *Memoria) Guardar(c curso.Curso) error {
	m.cursos = append(m.cursos, c)
	return nil
}

// Buscar recorre el slice. El primero con ese id gana.
func (m *Memoria) Buscar(id string) (curso.Curso, error) {
	for _, c := range m.cursos {
		if c.ID == id {
			return c, nil
		}
	}
	return curso.Curso{}, curso.ErrNoEncontrado
}

// Listar devuelve una copia del slice.
// Si devolviéramos m.cursos, quien recibe la lista podría
// modificar el arreglo interno del almacén.
func (m *Memoria) Listar() ([]curso.Curso, error) {
	salida := make([]curso.Curso, len(m.cursos))
	copy(salida, m.cursos)
	return salida, nil
}
