// Package curso es el dato que cruzan el servicio y el almacén.
// No tiene reglas de negocio: esas viven en el servicio.
package curso

import (
	"errors"
	"fmt"
)

// ErrNoEncontrado es el mismo valor para quien busca y para quien guarda.
// Así main puede reconocerlo con errors.Is, viva donde viva la búsqueda.
var ErrNoEncontrado = errors.New("curso no encontrado")

// Curso es un valor. ID, título y precio viajan juntos.
type Curso struct {
	ID     string
	Titulo string
	Precio float64
}

// Etiqueta formatea el curso para mostrarlo. No cambia nada.
func (c Curso) Etiqueta() string {
	return fmt.Sprintf("%s (%s) — $%.2f", c.Titulo, c.ID, c.Precio)
}
