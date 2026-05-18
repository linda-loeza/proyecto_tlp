package main

import "fmt"

// AgregarContacto pide los datos al usuario y los mete al slice global
func AgregarContacto() {
	var nuevo Contacto

	fmt.Println("\n--- Registrar Nuevo Contacto ---")

	fmt.Print("Nombre completo: ")
	fmt.Scanln(&nuevo.Nombre)

	fmt.Print("Fecha de cumpleaños (dd/mm): ")
	fmt.Scanln(&nuevo.Cumpleanos)

	fmt.Print("Teléfono: ")
	fmt.Scanln(&nuevo.Telefono)

	fmt.Print("Correo electrónico: ")
	fmt.Scanln(&nuevo.Email)

	// append es la función de Go para agregar elementos a un slice dinámico
	listaContactos = append(listaContactos, nuevo)

	fmt.Println("Contacto guardado con éxito en la memoria")
}

func MostrarContactos() {
	if len(listaContactos) == 0 {
		fmt.Println("No hay contactos registrados.")
		return
	} else {
		fmt.Println("\n--- Lista de Contactos ---")
		for _, contacto := range listaContactos {
			fmt.Println(contacto.StringContacto())
		}
	}
}
