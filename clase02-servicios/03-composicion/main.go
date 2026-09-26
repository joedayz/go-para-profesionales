// Clase 02 - Demo 3: Composición
// Curso: Go para Profesionales | joedayz.pe
//
// Una idea: un struct se arma metiendo otros structs dentro.
// Embeber presta campos y métodos. No es herencia.
package main

import "fmt"

// Persona es alguien con nombre. Nada más.
type Persona struct {
	Nombre string
	Pais   string
}

func (p Persona) Presentacion() string {
	return fmt.Sprintf("%s, %s", p.Nombre, p.Pais)
}

// SaludoCorto solo conoce a la persona. Quien lo herede por embedding
// no le agrega datos: el receptor sigue siendo Persona.
func (p Persona) SaludoCorto() string {
	return "Hola, soy " + p.Nombre
}

// Instructor embebe Persona: no escribe el nombre del campo.
// El campo se llama igual que el tipo, Persona.
// Nombre, Pais, Presentacion y SaludoCorto quedan promovidos:
// se usan como si fueran de Instructor.
type Instructor struct {
	Persona
	Especialidad string
}

// Presentacion de Instructor reemplaza al selector promovido.
// No borra el de Persona: sigue en instructor.Persona.Presentacion.
// En Java esto parecería un override. Aquí conviven los dos.
func (i Instructor) Presentacion() string {
	return fmt.Sprintf("%s · especialista en %s", i.Persona.Presentacion(), i.Especialidad)
}

// CursoEmbebido mete al instructor sin nombrarlo.
// El curso pasa a "tener" Nombre, SaludoCorto y Presentacion.
// El modelo miente: un curso no es una persona.
type CursoEmbebido struct {
	Instructor
	Titulo string
	Precio float64
}

// Curso nombra el campo. Un curso tiene un instructor.
// Nada se promueve: el nombre se lee en curso.Instructor.Nombre.
type Curso struct {
	Titulo     string
	Precio     float64
	Instructor Instructor
}

func (c Curso) Ficha() string {
	return fmt.Sprintf("%s ($%.2f) · dictado por %s",
		c.Titulo, c.Precio, c.Instructor.Presentacion())
}

func main() {
	joe := Instructor{
		Persona:      Persona{Nombre: "Joe", Pais: "Perú"},
		Especialidad: "Go",
	}

	// ─────────────────────────────────────────────
	// 1. Promoción
	//    joe.Nombre es joe.Persona.Nombre.
	//    No se copió el campo: es el mismo, un nivel adentro.
	// ─────────────────────────────────────────────
	fmt.Println("--- Instructor embebe Persona ---")
	fmt.Println("Nombre promovido:", joe.Nombre)
	fmt.Println("Presentación del instructor:", joe.Presentacion())
	fmt.Println("Presentación de la persona:", joe.Persona.Presentacion())

	// ─────────────────────────────────────────────
	// 2. El método promovido no ve al tipo de afuera
	//    SaludoCorto está definido en Persona.
	//    Aunque lo llames sobre el instructor, el receptor
	//    es la persona: no puede mencionar la especialidad.
	// ─────────────────────────────────────────────
	fmt.Println("Saludo promovido:", joe.SaludoCorto())

	// ─────────────────────────────────────────────
	// 3. Embeber de más
	//    El curso saluda y se presenta como Joe.
	//    El título del curso no entra en esa frase.
	// ─────────────────────────────────────────────
	acoplado := CursoEmbebido{
		Instructor: joe,
		Titulo:     "Go para profesionales",
		Precio:     149.90,
	}
	fmt.Println("\n--- Curso con el instructor embebido ---")
	fmt.Println("El curso se presenta así:", acoplado.Presentacion())
	fmt.Println("Y su 'nombre' es:", acoplado.Nombre)

	// ─────────────────────────────────────────────
	// 4. Composición nombrada
	//    curso.Nombre no existe. El nombre es del instructor.
	//    Ficha arma la frase del curso con los datos de cada uno.
	// ─────────────────────────────────────────────
	curso := Curso{
		Titulo:     "Go para profesionales",
		Precio:     149.90,
		Instructor: joe,
	}
	fmt.Println("\n--- Curso con instructor nombrado ---")
	fmt.Println(curso.Ficha())
	fmt.Println("Instructor:", curso.Instructor.Nombre)
}
