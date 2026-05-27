# Documento para llevar el orden de lo que falta en el proyecto

## Cumple (parcial)

- Estructura general de archivos: existen main.go, contacto.go, gestion.go, persistencia.go.a
- struct Contacto con nombre/cumpleaños/teléfono/correo: está en contacto.go.
- “Agregar” y “Mostrar”: están en gestion.go y se usan desde el menú en main.go.

## No cumple / falta implementar

- Persistencia real en contactos.txt: persistencia.go solo imprime mensajes; no lee/escribe archivo y tampoco se llama desde main.go (el PDF pide cargar al iniciar y guardar al salir).
- Ordenamiento al listar por nombre con sort: actualmente lista en el orden de inserción (ver gestion.go).
- CRUD completo: faltan “buscar por nombre” y “eliminar contacto” (el PDF lo pide explícitamente).
- Filtro de cumpleaños “hoy”: no existe lógica de fecha actual ni filtro (el PDF lo pide).
- Menú de 5 opciones: hoy el menú tiene 3 opciones (ver main.go).
- Validaciones básicas: no hay validación de formato/valores; además fmt.Scanln corta el “nombre completo” en el primer espacio (ver gestion.go).
