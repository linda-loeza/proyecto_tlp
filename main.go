package main

import (
	"fmt"
	"os"
	"proyecto_tlp/gestion"
	"proyecto_tlp/storage"
)

var listaContactos []gestion.Contacto //aqui se van a guardar los objetos de tipo contacto

func main() {
	fmt.Println("Iniciando sistema...")

	//carga inicial
	contactosGuardados, err := storage.CargarContactos()
	if err != nil {
		fmt.Println("Nota: No se pudo cargar el archivo de contactos (o aún no existe).")
		listaContactos = []gestion.Contacto{} // Inicializamos un slice vacío
	} else {
		listaContactos = contactosGuardados
		fmt.Printf("¡Se cargaron %d contactos exitosamente!\n", len(listaContactos))
	}

	for {
		fmt.Println("\nMenú de opciones:")
		fmt.Println("1. Agregar contacto")
		fmt.Println("2. Borrar un contacto")
		fmt.Println("3. Mostrar contacto especificado por nombre")
		fmt.Println("4. Mostrar contactos y su fecha de cumpleaños ordenados a partir del dia actual")
		fmt.Println("5. Mostrar contactos ordenados alfabeticamente por nombre")
		fmt.Println("6. Salir")
		fmt.Print("Seleccione una opción: ")

		var opcion int
		_, err := fmt.Scanln(&opcion)
		if err != nil {
			fmt.Println("Entrada no válida. Intente nuevamente.")
			var descartar string
			fmt.Scanln(&descartar) // "borra" la entrada no válida
			continue
		}

		//validacion
		if opcion < 1 || opcion > 6 {
			fmt.Println("Opción no válida. Intente nuevamente.")
			var descartar string
			fmt.Scanln(&descartar) //"borra" la entrada no válida
			continue
		}
		switch opcion {
		case 1:
			fmt.Println("\n[Opción 1: Agregar contacto]")
			gestion.AgregarContacto()
		case 2:
			fmt.Println("\n[Opción 2: Borrar un contacto]")
			// agregar

		case 3:
			fmt.Println("\n[Opción 3: Mostrar contacto especificado por nombre]")
			// ...

		case 4:
			fmt.Println("\n--- Contactos por Fecha de Cumpleaños (Desde hoy) ---")
			//	gestion.ImprimirTablaContactos(...)

		case 5:
			fmt.Println("\n--- Contactos Ordenados Alfabéticamente ---")
			//gestion.ImprimirTablaContactos(...)

		case 6:
			fmt.Println("Guardando cambios en el archivo y saliendo del sistema...")
			err := storage.GuardarContactos(listaContactos)
			if err != nil {
				fmt.Println("Error al escribir los datos en el disco duro:", err)
			} else {
				fmt.Println("archivo JSON actualizado con éxito")
			}
			os.Exit(0)
		}
	}

}
