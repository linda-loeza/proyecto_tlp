package gestion

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"
)

var listaContactos []Contacto

func ImprimirTablaContactos(contactos []Contacto) {
	// Inicializamos el tabwriter en la salida estándar
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 3, ' ', 0)

	// el encabezado
	fmt.Fprintln(w, "NOMBRE\tCUMPLEAÑOS\tTELÉFONO\tEMAIL")
	fmt.Fprintln(w, "--\t------\t--------\t-----")

	// cada fila
	for _, c := range contactos {
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

var stdinReader = bufio.NewReader(os.Stdin)

func leerLinea(prompt string) string {
	for {
		fmt.Print(prompt)
		texto, err := stdinReader.ReadString('\n')
		if err != nil {
			return strings.TrimSpace(texto)
		}
		texto = strings.TrimSpace(texto)
		if texto == "" {
			continue
		}
		return texto
	}
}

func normalizarDDMM(fecha string) (string, bool) {
	partes := strings.Split(strings.TrimSpace(fecha), "/")
	if len(partes) != 2 {
		return "", false
	}

	dia, errDia := strconv.Atoi(strings.TrimSpace(partes[0]))
	mes, errMes := strconv.Atoi(strings.TrimSpace(partes[1]))
	if errDia != nil || errMes != nil {
		return "", false
	}
	if dia < 1 || dia > 31 || mes < 1 || mes > 12 {
		return "", false
	}

	return fmt.Sprintf("%02d/%02d", dia, mes), true
}

// FiltrarCumpleanosEnFecha devuelve los contactos cuyo cumpleaños coincide con la fecha dd/mm.
func FiltrarCumpleanosEnFecha(contactos []Contacto, ddmm string) []Contacto {
	fechaObjetivo, ok := normalizarDDMM(ddmm)
	if !ok {
		return nil
	}

	var resultados []Contacto
	for _, c := range contactos {
		fechaContacto, ok := normalizarDDMM(c.Cumpleanos)
		if !ok {
			continue
		}
		if fechaContacto == fechaObjetivo {
			resultados = append(resultados, c)
		}
	}
	return resultados
}

// FiltrarCumpleanosHoy devuelve los contactos que cumplen años hoy.
func FiltrarCumpleanosHoy(contactos []Contacto) []Contacto {
	ahora := time.Now()
	hoy := fmt.Sprintf("%02d/%02d", ahora.Day(), int(ahora.Month()))
	return FiltrarCumpleanosEnFecha(contactos, hoy)
}

func validarEmail(email string) bool {
	e := strings.TrimSpace(email)
	at := strings.Index(e, "@")
	dot := strings.LastIndex(e, ".")
	return at > 0 && dot > at+1 && dot < len(e)-1
}

func validarTelefono(telefono string) bool {
	t := strings.TrimSpace(telefono)
	if len(t) < 7 || len(t) > 15 {
		return false
	}
	for i := 0; i < len(t); i++ {
		if t[i] < '0' || t[i] > '9' {
			return false
		}
	}
	return true
}

// validarFechaDDMM valida formato dd/mm y rangos básicos.
func validarFechaDDMM(fecha string) bool {
	f := strings.TrimSpace(fecha)
	if len(f) != 5 || f[2] != '/' {
		return false
	}

	if f[0] < '0' || f[0] > '9' || f[1] < '0' || f[1] > '9' || f[3] < '0' || f[3] > '9' || f[4] < '0' || f[4] > '9' {
		return false
	}

	dia, errDia := strconv.Atoi(f[0:2])
	mes, errMes := strconv.Atoi(f[3:5])
	if errDia != nil || errMes != nil {
		return false
	}
	if dia < 1 || dia > 31 || mes < 1 || mes > 12 {
		return false
	}
	return true
}

// AgregarContacto pide los datos al usuario y los mete al slice global
func AgregarContacto(contactos []Contacto) []Contacto {
	var nuevo Contacto

	fmt.Println("\n--- Registrar Nuevo Contacto ---")

	nuevo.Nombre = leerLinea("Nombre completo: ")

	for {
		fmt.Print("Fecha de cumpleaños (dd/mm): ")
		fmt.Scanln(&nuevo.Cumpleanos)
		if validarFechaDDMM(nuevo.Cumpleanos) {
			break
		}
		fmt.Println("Fecha inválida. Use el formato dd/mm.")
	}

	for {
		fmt.Print("Teléfono: ")
		fmt.Scanln(&nuevo.Telefono)
		if validarTelefono(nuevo.Telefono) {
			break
		}
		fmt.Println("Teléfono inválido. Use solo números (7-15 dígitos).")
	}

	for {
		fmt.Print("Correo electrónico: ")
		fmt.Scanln(&nuevo.Email)
		if validarEmail(nuevo.Email) {
			break
		}
		fmt.Println("Email inválido. Debe contener '@' y '.' (ej: a@b.com).")
	}

	// append es la función de Go para agregar elementos a un slice dinámico
	contactos = append(contactos, nuevo)

	fmt.Println("Contacto guardado con éxito en la memoria")
	return contactos
}

// BuscarPorNombre retorna los contactos cuyo nombre contiene el texto buscado (ignorando mayúsculas/minúsculas).
func BuscarPorNombre(nombre string) []Contacto {
	buscado := strings.ToLower(strings.TrimSpace(nombre))
	if buscado == "" {
		return nil
	}

	var resultados []Contacto
	for _, contacto := range listaContactos {
		if strings.Contains(strings.ToLower(contacto.Nombre), buscado) {
			resultados = append(resultados, contacto)
		}
	}
	return resultados
}

func BuscarContacto(contactos []Contacto) {
	fmt.Println("\n--- Buscar Contacto ---")
	query := strings.ToLower(strings.TrimSpace(leerLinea("Nombre a buscar: ")))
	if query == "" {
		return
	}

	var resultados []Contacto
	for _, c := range contactos {
		if strings.Contains(strings.ToLower(c.Nombre), query) {
			resultados = append(resultados, c)
		}
	}

	if len(resultados) == 0 {
		fmt.Println("No se encontraron contactos con ese criterio.")
		return
	}

	fmt.Printf("Se encontraron %d contacto(s):\n", len(resultados))
	ImprimirTablaContactos(resultados)
}

func EliminarContacto(contactos []Contacto) []Contacto {
	fmt.Println("\n--- Eliminar Contacto ---")
	nombre := strings.ToLower(strings.TrimSpace(leerLinea("Nombre exacto a eliminar: ")))
	if nombre == "" {
		return contactos
	}

	for i, c := range contactos {
		if strings.ToLower(strings.TrimSpace(c.Nombre)) == nombre {
			contactos = append(contactos[:i], contactos[i+1:]...)
			fmt.Println("Contacto eliminado con éxito.")
			return contactos
		}
	}

	fmt.Println("No se encontró un contacto con ese nombre exacto.")
	return contactos
}

// toma los contactos del main, los ordena cronológicamente desde hoy y los regresa ordenados
func OrdenarPorCumpleanos(contactos []Contacto) []Contacto {
	ahora := time.Now()
	mesActual := int(ahora.Month())
	diaActual := ahora.Day()

	// Creamos una copia del slice para no alterar el orden original
	listaOrdenada := make([]Contacto, len(contactos))
	copy(listaOrdenada, contactos)

	// sort.Slice para reordenar los elementos basándonos en la cercanía de la fecha
	sort.Slice(listaOrdenada, func(i, j int) bool {
		var diaI, mesI, diaJ, mesJ int

		// se parsea el string "dd/mm" a números enteros usando Sscanf
		fmt.Sscanf(listaOrdenada[i].Cumpleanos, "%d/%d", &diaI, &mesI)
		fmt.Sscanf(listaOrdenada[j].Cumpleanos, "%d/%d", &diaJ, &mesJ)

		// Calculamos cuántos meses faltan para el cumpleaños de I
		mesesRestantesI := mesI - mesActual
		if mesesRestantesI < 0 || (mesesRestantesI == 0 && diaI < diaActual) {
			mesesRestantesI += 12 // Si ya pasó este año, se cuenta para el próximo ciclo de 12 meses
		}

		// Calculamos cuántos meses faltan para el cumpleaños de J
		mesesRestantesJ := mesJ - mesActual
		if mesesRestantesJ < 0 || (mesesRestantesJ == 0 && diaJ < diaActual) {
			mesesRestantesJ += 12
		}

		// Si están en el mismo mes de espera, se ordena por el que tenga el día más cercano
		if mesesRestantesI == mesesRestantesJ {
			return diaI < diaJ
		}

		return mesesRestantesI < mesesRestantesJ
	})

	return listaOrdenada
}

func OrdenarAlfabeticamente(contactos []Contacto) []Contacto {
	listaOrdenada := make([]Contacto, len(contactos))
	copy(listaOrdenada, contactos)

	// Ordenamos comparando los nombres convertidos a minúsculas
	sort.Slice(listaOrdenada, func(i, j int) bool {
		return strings.ToLower(listaOrdenada[i].Nombre) < strings.ToLower(listaOrdenada[j].Nombre)
	})

	return listaOrdenada
}
