package storage

import (
	"bufio"
	"fmt"
	"os"
	"proyecto_tlp/gestion"
	"strings"
)

const NombreArchivo = "contactos.txt"

func sanitizarCampo(valor string) string {
	valor = strings.TrimSpace(valor)
	valor = strings.ReplaceAll(valor, "\t", " ")
	valor = strings.ReplaceAll(valor, "\r", " ")
	valor = strings.ReplaceAll(valor, "\n", " ")
	return valor
}

func encodeContacto(c gestion.Contacto) string {
	// Formato simple por línea (TSV): nombre\tcumpleanos\ttelefono\temail
	// Elegimos tabulador como separador para permitir nombres con espacios.
	return fmt.Sprintf("%s\t%s\t%s\t%s",
		sanitizarCampo(c.Nombre),
		sanitizarCampo(c.Cumpleanos),
		sanitizarCampo(c.Telefono),
		sanitizarCampo(c.Email),
	)
}

func decodeContacto(linea string) (gestion.Contacto, bool) {
	linea = strings.TrimSpace(linea)
	if linea == "" {
		return gestion.Contacto{}, false
	}

	partes := strings.Split(linea, "\t")
	if len(partes) != 4 {
		// Soporte mínimo si el archivo viene con '|'
		partes = strings.Split(linea, "|")
		if len(partes) != 4 {
			return gestion.Contacto{}, false
		}
	}

	return gestion.Contacto{
		Nombre:     strings.TrimSpace(partes[0]),
		Cumpleanos: strings.TrimSpace(partes[1]),
		Telefono:   strings.TrimSpace(partes[2]),
		Email:      strings.TrimSpace(partes[3]),
	}, true
}

func CargarContactos() ([]gestion.Contacto, error) {
	fmt.Println("Buscando datos guardados...")
	if _, err := os.Stat(NombreArchivo); os.IsNotExist(err) {
		fmt.Println("No se encontraron datos guardados.")
		return []gestion.Contacto{}, nil
	}

	archivo, err := os.Open(NombreArchivo)
	if err != nil {
		return nil, err
	}
	defer archivo.Close()

	scanner := bufio.NewScanner(archivo)
	// Aumentar tamaño de línea por si hay nombres largos
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)

	var contactosCargados []gestion.Contacto
	lineaN := 0
	for scanner.Scan() {
		lineaN++
		c, ok := decodeContacto(scanner.Text())
		if !ok {
			// Permitimos líneas vacías; si no es vacía y no parsea, marcamos error.
			if strings.TrimSpace(scanner.Text()) == "" {
				continue
			}
			return nil, fmt.Errorf("formato inválido en %s (línea %d)", NombreArchivo, lineaN)
		}
		contactosCargados = append(contactosCargados, c)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	fmt.Printf("Datos cargados con éxito. %d contactos encontrados.\n", len(contactosCargados))
	return contactosCargados, nil
}

func GuardarContactos(contactos []gestion.Contacto) error {
	fmt.Println("Guardando los contactos en 'contactos.txt'...")

	archivo, err := os.Create(NombreArchivo)
	if err != nil {
		return err
	}
	defer archivo.Close()

	w := bufio.NewWriter(archivo)
	for _, c := range contactos {
		if _, err := fmt.Fprintln(w, encodeContacto(c)); err != nil {
			return err
		}
	}
	if err := w.Flush(); err != nil {
		return err
	}
	fmt.Println("Contactos guardados con éxito.")
	return nil
}
