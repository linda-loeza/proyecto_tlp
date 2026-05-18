package main

import (
	"fmt"
	"os"
)

var listaContactos []Contacto //aqui se van a guardar los objetos de tipo contacto

func main() {
	fmt.Println("Iniciando sistema...")

	for {
		fmt.Println("\nMenú de opciones:")
		fmt.Println("1. Agregar contacto")
		fmt.Println("2. Mostrar contactos")
		fmt.Println("3. Salir")
		fmt.Print("Seleccione una opción: ")
		var opcion int
		fmt.Scanln(&opcion)

		//validacion
		if opcion < 1 || opcion > 3 {
			fmt.Println("Opción no válida. Intente nuevamente.")
			var descartar string
			fmt.Scanln(&descartar) //"borra" la entrada no válida
			continue
		}
		switch opcion {
		case 1:
			fmt.Println("Agregando nuevo contacto...")
			AgregarContacto()
		case 2:
			fmt.Println("Mostrando contactos...")
			MostrarContactos()
		case 3:
			fmt.Println("Saliendo del sistema...")
			os.Exit(0)
		default:
			fmt.Println("Opción no válida. Intente nuevamente.")
		}
	}

}
