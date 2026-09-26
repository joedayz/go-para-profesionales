# Clase 02 — Programación orientada a servicios

> **Curso:** Go para Profesionales · [joedayz.pe](https://www.joedayz.pe/curso/online/go-para-profesionales)  
> **Dirigido a:** Profesionales de Java, C#, PHP que quieren dominar Go para backend  
> **Cómo verla:** un demo por video, en orden. Cada carpeta se ejecuta sola.

Go no organiza el backend con jerarquías de clases. Organiza datos (`struct`), comportamiento (métodos), piezas armadas (composición) y bordes (interfaces pequeñas). El servicio coordina. El `main` elige las piezas concretas.

---

## De menos a más

| Demo | Quédate con esto |
|---|---|
| 01 Structs | Un struct es un valor. Asignarlo copia. |
| 02 Métodos | El receptor decide si cambias la copia o el original. |
| 03 Composición | Embeber no es heredar. Si el curso no es un instructor, nómbralo. |
| 04 Interfaces | La interfaz la cumple quien tiene los métodos, aunque no lo declare. |
| 05 Diseño | El servicio pide un comportamiento. `main` decide quién lo cumple. |

La historia es la misma en los cinco: publicar cursos. El struct crece de significado, no hace falta arrastrar el código del demo anterior.

---

## ¿Qué aprenderás en esta clase?

- Declarar un `struct`, su zero value y la diferencia entre copia y puntero
- Colgar métodos de un tipo, con receptor por valor o por puntero
- Validar en un constructor (`NuevoCurso`, `Nuevo`)
- Armar tipos por composición, y cuándo embeber presta métodos que no quieres
- Definir una interfaz pequeña del lado de quien la usa
- Inyectar esa interfaz en un servicio y cambiar la implementación sin tocarlo

---

## Estructura del código

```
clase02-servicios/
├── go.mod
├── README.md
├── 01-structs/
│   └── main.go                      ← valor, copia, puntero
├── 02-metodos/
│   └── main.go                      ← receptor, constructor
├── 03-composicion/
│   └── main.go                      ← embedding vs campo nombrado
├── 04-interfaces/
│   └── main.go                      ← contrato implícito, nil
└── 05-diseno/
    ├── main.go                      ← arma el servicio y lo prueba
    ├── curso/curso.go               ← el dato
    ├── memoria/memoria.go           ← un almacén; no conoce al servicio
    └── catalogo/catalogo.go         ← reglas + interfaz Repositorio
```

En el demo 5 abre los archivos en este orden: `curso`, `memoria`, `catalogo`, `main`.

`memoria` no importa a `catalogo`. `catalogo` no importa a `memoria`. Los dos hablan de `curso.Curso`. `main` es el único que conoce a los dos y enchufa uno dentro del otro.

---

## Cómo ejecutar cada demo

Asegúrate de tener Go ≥ 1.23 instalado: `go version`

Desde esta carpeta (`clase02-servicios/`):

```bash
# Demo 1 — El dato
go run ./01-structs/

# Demo 2 — El comportamiento
go run ./02-metodos/

# Demo 3 — Armar tipos
go run ./03-composicion/

# Demo 4 — El contrato
go run ./04-interfaces/

# Demo 5 — El servicio
go run ./05-diseno/
```

---

## Conceptos clave por demo

### Demo 1 · Structs

| Concepto | Qué mostrar |
|---|---|
| Zero value | `var vacio Curso` existe: strings vacíos, precio 0, `Activo` en false |
| Literal | Los campos se nombran; lo que omites queda en cero |
| Copia | `copia := goBasico` y luego cambiar `copia.Precio` no toca el original |
| Puntero | `ptr := &goBasico` sí cambia el original. El slice ya guardó una copia y no se entera |
| Igualdad | `==` compara campo por campo si todos los campos son comparables |

### Demo 2 · Métodos

| Concepto | Qué mostrar |
|---|---|
| Receptor por valor | `demostrarCopia` modifica una copia y el original sigue en 200 |
| Devolver la copia | `ConDescuento` entrega otro `Curso`; el original no se mueve |
| Receptor puntero | `AplicarDescuento` escribe en el original |
| Dirección automática | Sobre una variable, Go permite `base.AplicarDescuento` sin escribir `&base` |
| Constructor | `NuevoCurso` rechaza título vacío y precio negativo |

La mayúscula de `NuevoCurso` o `Etiqueta` no esconde nada mientras todo vive en `package main`. Pasa a importar en el demo 5: ahí la mayúscula es la frontera.

### Demo 3 · Composición

| Concepto | Qué mostrar |
|---|---|
| Embedding | `Instructor` mete `Persona` sin nombre de campo. `joe.Nombre` es el de la persona |
| Promoción | `SaludoCorto` se llama sobre el instructor, pero el receptor sigue siendo `Persona` |
| Los dos métodos conviven | `joe.Presentacion()` ve la especialidad. `joe.Persona.Presentacion()` no |
| Embeber de más | `CursoEmbebido` se presenta como si el curso fuera Joe |
| Campo nombrado | `Curso` tiene un `Instructor`. La ficha la arma el curso |

Embebe cuando decir `instructor.Nombre` se lee mejor que `instructor.Persona.Nombre` y el rol no se confunde. Nombra el campo cuando las dos cosas son distintas.

### Demo 4 · Interfaces

| Concepto | Qué mostrar |
|---|---|
| Sin `implements` | `Memoria` y `CatalogoFijo` no mencionan a `Buscador` |
| Uso | `describir` acepta la interfaz y llama `Buscar` |
| Lista mixta | `[]Buscador{memoria, fijo}` |
| Type switch | Recupera el concreto solo cuando el origen importa |
| Ok | `b.(Memoria)` sin `ok` entra en pánico si el tipo no es ese |
| Nil | Interfaz vacía `== nil`. Interfaz con `*Memoria` nulo adentro `!= nil` |

`Buscar` devuelve `bool` a propósito, como el ok de un map. En el demo 5 esa señal pasa a ser un `error`: no encontrar un curso y que falle el disco no son lo mismo.

### Demo 5 · Diseño

| Pieza | Responsabilidad |
|---|---|
| `curso.Curso` | El dato que viaja. `ErrNoEncontrado` vive aquí para que los dos lados compartan el mismo valor |
| `catalogo.Repositorio` | Los tres métodos que el servicio necesita. La interfaz está del lado de quien la usa |
| `catalogo.Servicio` | Título, precio e id. No importa `memoria` |
| `memoria.Memoria` | Slice en orden de llegada. Devuelve copias. No importa `catalogo` |
| `main` | Construye, inyecta y, al final, cambia el almacén por `repositorioRoto` |

Dos detalles de método que cierran el demo 2:

- `Guardar` tiene receptor `*Memoria` porque modifica el slice. Quien entra en la interfaz es `*Memoria`, no `Memoria`.
- `var _ catalogo.Repositorio = (*memoria.Memoria)(nil)` no corre: si la firma deja de coincidir, no compila.

`Publicar` envuelve el fallo del disco con `fmt.Errorf("guardar curso: %w", err)`. El texto gana contexto y `errors.Is` todavía puede ver el error de adentro.

---

## Diferencias clave vs Java / C#

| Go | Java / C# |
|---|---|
| `struct` + métodos fuera del tipo | Clase que junta campos y métodos |
| Composición y embedding | Herencia (`extends` / `:`) |
| Interfaz implícita, chica, del lado de quien la usa | `implements`, a menudo declarada junto a la clase |
| Aceptar interfaces, devolver tipos concretos | Inyectar la interfaz que definió el framework o la impl |
| `NuevoX(...)` valida y construye | Constructor de clase, a veces más setters |
| `error` como valor de retorno | Excepción que sube sola por el stack |
| Mayúscula / minúscula | `public`, `private`, `protected` |

---

## Reglas de esta clase

1. **Acepta interfaces, devuelve structs.** `Nuevo` del catálogo recibe un `Repositorio` y devuelve un `*Servicio`. `memoria.Nuevo` devuelve `*Memoria`, no la interfaz.
2. **Interfaz chica.** `Buscador` tiene un método. `Repositorio` tiene tres, porque el servicio usa tres. Si un método no se llama, no va en el contrato.
3. **La declara quien la usa.** `catalogo` pide `Repositorio`. `memoria` solo tiene métodos.

---

## Clase anterior

**Clase 01 — Fundamentos de Go**  
Sintaxis, tipos, funciones, control de flujo y paquetes. El demo 5 de esta clase apoya la visibilidad (mayúscula / minúscula) y el `error` como valor que ya vimos ahí.

Código: [`clase01-fundamentos`](../clase01-fundamentos/)
