package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"proyecto_tlp/gestion"
)

const NombreArchivo = "contactos.json"

func CargarContactos() ([]gestion.Contacto, error) {
	fmt.Println("Buscando datos guardados...")
	if _, err := os.Stat(NombreArchivo); os.IsNotExist(err) {
		fmt.Println("No se encontraron datos guardados.")
		return []gestion.Contacto{}, nil
	}

	datosBytes, err := os.ReadFile(NombreArchivo)
	if err != nil {
		return nil, err
	}

	var contactosCargados []gestion.Contacto
	err = json.Unmarshal(datosBytes, &contactosCargados)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Datos cargados con éxito. %d contactos encontrados.\n", len(contactosCargados))
	return contactosCargados, nil
}

func GuardarContactos(contactos []gestion.Contacto) error {
	fmt.Println("Guardando los contactos en 'contactos.json'...")

	datosJSON, err := json.MarshalIndent(contactos, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(NombreArchivo, datosJSON, 0644)
	if err != nil {
		return err
	}
	fmt.Println("Contactos guardados con éxito.")
	return nil
}
