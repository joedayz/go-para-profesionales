// Package catalogo publica cursos.
// Conoce el dato (curso.Curso) y el comportamiento que necesita guardar.
// No conoce a memoria ni a ningún otro almacén.
package catalogo

import (
	"errors"
	"fmt"
	"strings"

	"go-para-profesionales/clase02/05-diseno/curso"
)

var (
	// ErrTituloVacio rechaza un título en blanco o hecho solo de espacios.
	ErrTituloVacio = errors.New("el título no puede estar vacío")

	// ErrPrecioInvalido rechaza precios negativos.
	// Cero sí pasa: es un curso gratuito.
	ErrPrecioInvalido = errors.New("el precio no puede ser negativo")
)

// Repositorio es lo que el servicio necesita para recordar cursos.
// La declara este paquete, que es quien la usa.
// Cualquier tipo con estos tres métodos sirve, esté en el paquete que esté.
type Repositorio interface {
	Guardar(curso.Curso) error
	Buscar(id string) (curso.Curso, error)
	Listar() ([]curso.Curso, error)
}

// Servicio aplica las reglas y habla con un Repositorio.
// El campo repo es la interfaz, no un *memoria.Memoria.
type Servicio struct {
	repo      Repositorio
	siguiente int
}

// Nuevo se niega a arrancar sin almacén.
// repo == nil solo es cierto si la interfaz está vacía
// (sin tipo y sin valor). Un puntero nulo ya metido dentro
// de la interfaz no se detecta aquí: eso se vio en el demo 4.
func Nuevo(repo Repositorio) (*Servicio, error) {
	if repo == nil {
		return nil, errors.New("el catálogo necesita un repositorio")
	}
	return &Servicio{repo: repo}, nil
}

// Publicar valida, asigna el id y pide al almacén que guarde.
// El id sale de aquí, no del repositorio: identificar un curso
// es decisión del servicio.
func (s *Servicio) Publicar(titulo string, precio float64) (curso.Curso, error) {
	titulo = strings.TrimSpace(titulo)
	if titulo == "" {
		return curso.Curso{}, ErrTituloVacio
	}
	if precio < 0 {
		return curso.Curso{}, ErrPrecioInvalido
	}

	nuevo := curso.Curso{
		ID:     fmt.Sprintf("CUR-%03d", s.siguiente+1),
		Titulo: titulo,
		Precio: precio,
	}
	if err := s.repo.Guardar(nuevo); err != nil {
		// %w conserva el error original. errors.Is sigue pudiendo verlo.
		return curso.Curso{}, fmt.Errorf("guardar curso: %w", err)
	}
	s.siguiente++
	return nuevo, nil
}

// Detalle delega la búsqueda. Hoy no agrega reglas.
// Mañana puede sumar auditoría sin que main abra el almacén.
func (s *Servicio) Detalle(id string) (curso.Curso, error) {
	return s.repo.Buscar(id)
}

// Listar devuelve lo que el almacén tenga, en el orden en que lo guardó.
func (s *Servicio) Listar() ([]curso.Curso, error) {
	return s.repo.Listar()
}
