package main

import "fmt"

type Contacto struct {
	Nombre     string
	Cumpleanos string
	Email      string
	Telefono   string
}

func (c Contacto) StringContacto() string {
	return fmt.Sprintf("Nombre: %-25s\nCumpleaños: %s\nEmail: %s\nTeléfono: %s",
		c.Nombre, c.Cumpleanos, c.Email, c.Telefono)
}
