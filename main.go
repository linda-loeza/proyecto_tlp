package main

import (
	"fmt"
	"os"
	"proyecto_tlp/gestion"
	"proyecto_tlp/storage"
)

var listaContactos []gestion.Contacto

func main() {
	fmt.Println("Iniciando sistema...")

	// carga inicial desde storage
	contactosGuardados, err := storage.CargarContactos()
	if err != nil {
		fmt.Println("Nota: No se pudo cargar el archivo de contactos (o aún no existe).")
		listaContactos = []gestion.Contacto{}
	} else {
		listaContactos = contactosGuardados
		fmt.Printf("¡Se cargaron %d contactos exitosamente!\n", len(listaContactos))
	}

	for {
		fmt.Println("\nMenú de opciones:")
		fmt.Println("1. Agregar contacto")
		fmt.Println("2. Mostrar contactos")
		fmt.Println("3. Buscar contacto por nombre")
		fmt.Println("4. Eliminar contacto")
		fmt.Println("5. Salir")
		fmt.Print("Seleccione una opción: ")

		var opcion int
		_, err := fmt.Scanln(&opcion)
		if err != nil {
			fmt.Println("Entrada no válida. Intente nuevamente.")
			var descartar string
			fmt.Scanln(&descartar)
			continue
		}

		if opcion < 1 || opcion > 5 {
			fmt.Println("Opción no válida. Intente nuevamente.")
			var descartar string
			fmt.Scanln(&descartar)
			continue
		}

		switch opcion {
		case 1:
			listaContactos = gestion.AgregarContacto(listaContactos)

		case 2:
			fmt.Println("Mostrando contactos...")
			ordenados := gestion.OrdenarAlfabeticamente(listaContactos)
			gestion.ImprimirTablaContactos(ordenados)
			cumpleHoy := gestion.FiltrarCumpleanosHoy(listaContactos)
			if len(cumpleHoy) > 0 {
				fmt.Println("\n--- Cumpleaños hoy ---")
				gestion.ImprimirTablaContactos(cumpleHoy)
			}

		case 3:
			gestion.BuscarContacto(listaContactos)

		case 4:
			listaContactos = gestion.EliminarContacto(listaContactos)

		case 5:
			fmt.Println("Guardando cambios en el archivo y saliendo del sistema...")
			if err := storage.GuardarContactos(listaContactos); err != nil {
				fmt.Println("Error al escribir los datos en el disco duro:", err)
			} else {
				fmt.Println("contactos.txt actualizado con éxito")
			}
			os.Exit(0)
		}
	}

}
