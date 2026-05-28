# Sistema de Gestión de Contactos y Cumpleaños

Este es un sistema modular de consola desarrollado en **Go (Golang)** como proyecto para la asignatura de **Teoría de Lenguajes de Programación (TLP)** en la **Facultad de Matemáticas de la Universidad Autónoma de Yucatán (UADY)**.

El sistema permite administrar una agenda de contactos local, aplicando validaciones estrictas de datos, algoritmos de ordenamiento personalizados y persistencia nativa en formato JSON.

## Características Principales

* **Persistencia en JSON:** Los contactos se cargan automáticamente al iniciar el sistema desde un archivo `contactos.json` y se guardan de forma segura al salir.
* **Validaciones:** El menú y las capturas están blindados contra entradas inválidas (letras en campos numéricos, formatos de fecha incorrectos, etc.).
* **Algoritmos de Ordenamiento:**
    * **Alfabético:** Ordena los contactos de la A a la Z (insensible a mayúsculas/minúsculas).
    * **Próximos Cumpleaños:** Un algoritmo cronológico dinámico que calcula los días restantes a partir de la **fecha actual del sistema**, dejando al principio los cumpleaños más cercanos y enviando los que ya pasaron al ciclo del siguiente año.

---

## Estructura del Proyecto

El proyecto sigue una arquitectura modular separada por responsabilidades de paquetes:

```text
proyecto_tlp/
├── main.go            # Orquestador del ciclo principal, menú interactivo y flujos.
├── go.mod             # Definición del módulo de Go.
├── contactos.json     # Base de datos en texto plano (generada automáticamente).
├── gestion/
│   └── gestion.go     # Estructura de datos (Contacto), validaciones y algoritmos de ordenamiento.
└── storage/
    └── persistencia.go # Lógica de lectura y escritura del archivo JSON usando el paquete `os`.
