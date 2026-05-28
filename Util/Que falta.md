# Documento para llevar el orden de lo que falta en el proyecto

NOTA: El análisis esta hecho por IA para comprobar que falta en el proyecto, que se puede mejorar o que esta implementado parcialmente de acuerdo al PDF

## Cumple

- Estructura general de archivos: existen main.go, contacto.go, gestion.go, persistencia.go.
- struct Contacto con nombre/cumpleaños/teléfono/correo: está en contacto.go.
- Menú de usuario (5 opciones): agregar, mostrar, buscar, eliminar, salir (main.go).
- CRUD en memoria: agregar, buscar por nombre y eliminar contacto (gestion.go).
- Filtro de cumpleaños “hoy”: existe lógica con fecha actual y se muestra al listar (gestion.go / main.go).

## No cumple / falta implementar

- Persistencia real en contactos.txt: persistencia.go solo imprime mensajes; no lee/escribe archivo y tampoco se llama desde main.go (el PDF pide cargar al iniciar y guardar al salir).
- Ordenamiento al listar por nombre con sort: actualmente lista en el orden de inserción (ver gestion.go).
- Validaciones básicas (parcial): se valida formato de fecha dd/mm (ej: dd/mm) y email (contiene '@' y '.'); falta validar teléfono y otros formatos.
- Entrada de nombre: se captura con fmt.Scanln (1 palabra); actualmente el prompt pide “Nombre (Sin apellidos)”. Si se requiere “nombre completo”, habría que leer línea completa.
