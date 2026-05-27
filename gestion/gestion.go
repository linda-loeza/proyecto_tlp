package gestion

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"text/tabwriter"
)

var listaContactos []Contacto

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

func ImprimirTablaContactos() {
	// Inicializamos el tabwriter en la salida estándar
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 3, ' ', 0)

	// el encabezado
	fmt.Fprintln(w, "NOMBRE\tCUMPLEAÑOS\tTELÉFONO\tEMAIL")
	fmt.Fprintln(w, "--\t------\t--------\t-----")

	// cada fila
	for _, c := range listaContactos {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", c.Nombre, c.Cumpleanos, c.Telefono, c.Email)
	}

	w.Flush() // Muestra el resultado final en pantalla
}

// LimpiarPantalla borra el texto viejo de la terminal
func LimpiarPantalla() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
