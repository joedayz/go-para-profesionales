# Clase 01 — Fundamentos de Go

> **Curso:** Go para Profesionales · [joedayz.pe](https://www.joedayz.pe/curso/online/go-para-profesionales)  
> **Dirigido a:** Profesionales de Java, C#, PHP que quieren dominar Go para backend

---

## ¿Qué aprenderás en esta clase?

- Sintaxis básica y convenciones del lenguaje
- Tipos de datos: primitivos, slices, maps y zero values
- Funciones: múltiples retornos, variadic, closures y defer
- Estructuras de control: `if`, `for` (único bucle), `switch`
- Organización del código con paquetes

---

## Estructura del código

```
clase01-fundamentos/
├── go.mod                        ← módulo del proyecto
├── 01-sintaxis/
│   └── main.go                   ← variables, constantes, fmt
├── 02-tipos/
│   └── main.go                   ← int, string, bool, slice, map, zero values
├── 03-funciones/
│   └── main.go                   ← funciones, closures, defer
├── 04-control-flujo/
│   └── main.go                   ← if, for (5 formas), switch, labels
└── 05-paquetes/
    ├── main.go                   ← consume el paquete local
    └── mathutil/
        └── mathutil.go           ← paquete reutilizable (exportado vs privado)
```

---

## Cómo ejecutar cada demo

Asegúrate de tener Go ≥ 1.23 instalado: `go version`

```bash
# Demo 1 — Sintaxis básica
go run ./01-sintaxis/

# Demo 2 — Tipos de datos
go run ./02-tipos/

# Demo 3 — Funciones
go run ./03-funciones/

# Demo 4 — Control de flujo
go run ./04-control-flujo/

# Demo 5 — Paquetes
go run ./05-paquetes/
```

---

## Conceptos clave por demo

### Demo 1 · Sintaxis básica
| Concepto | Ejemplo |
|---|---|
| Declaración con tipo | `var nombre string = "Joe"` |
| Inferencia de tipo | `pais := "Perú"` |
| Zero value | `var saldo float64` → `0.0` |
| Constante | `const version = "1.0.0"` |
| Swap idiomático | `x, y = y, x` |

### Demo 2 · Tipos de datos
| Tipo | Notas |
|---|---|
| `int`, `float64` | Los más usados en producción |
| `string` | Secuencia de bytes UTF-8; iterar con `range` para runas |
| `[]T` (slice) | Dinámico; `append`, `len`, `cap` |
| `map[K]V` | Acceso seguro con ok-pattern: `v, ok := m[k]` |
| Zero values | `int→0`, `float64→0.0`, `string→""`, `bool→false` |

### Demo 3 · Funciones
| Patrón | Descripción |
|---|---|
| Múltiples retornos | `func dividir(a, b float64) (float64, error)` |
| Retornos nombrados | `func minMax(nums []int) (min, max int)` |
| Variadic | `func sumar(nums ...int) int` |
| Closure | Función que captura variables de su entorno |
| `defer` | Ejecuta al final de la función (LIFO si son varios) |

### Demo 4 · Control de flujo
| Estructura | Diferencia vs otros lenguajes |
|---|---|
| `if` con init | `if v, ok := m[k]; ok { ... }` — scope limitado |
| `for` clásico | `for i := 0; i < n; i++` |
| `for` como while | `for condicion { ... }` |
| `for range` | Itera slices, maps, strings, channels |
| `for {}` | Loop infinito; salir con `break` |
| `switch` | Sin `break` explícito; soporta múltiples casos |
| Labels | `break etiqueta` para salir de loops anidados |

### Demo 5 · Paquetes
| Concepto | Descripción |
|---|---|
| Exportado | Identificador en **M**ayúscula → visible fuera del paquete |
| No exportado | Identificador en **m**inúscula → privado al paquete |
| Import path | Se construye desde el `module` en `go.mod` |
| Alias de import | `import mu "ruta/mathutil"` |
| Blank import | `import _ "paquete"` → solo registra `init()` |

---

## Diferencias clave vs Java / C#

| Go | Java / C# |
|---|---|
| `:=` infiere el tipo | `var` o tipo explícito siempre |
| Un solo bucle `for` | `for`, `while`, `do-while`, `foreach` |
| Múltiples retornos | Necesitas objetos/tuplas/out params |
| `error` como valor | Excepciones con `try/catch` |
| Visibilidad por mayúscula | `public`, `private`, `protected` |
| Sin clases, solo structs | Todo gira alrededor de clases |
| `defer` nativo | `finally` o `using`/`try-with-resources` |

---

## Próxima clase

**Clase 02 — Programación orientada a servicios**  
Structs, interfaces, composición y diseño limpio en Go.
