package main

import "fmt"

func CargarContactos() {
	fmt.Println("Buscando datos guardados...")

}

func GuardarContactos() {
	fmt.Println("Guardando los contactos en 'contactos.txt'...")
	fmt.Printf(" %d contactos guardados con éxito.\n", len(listaContactos))
}
