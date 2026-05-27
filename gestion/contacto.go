package gestion

import "fmt"

type Contacto struct {
	Nombre     string `json:"nombre"`
	Cumpleanos string `json:"cumpleanos"`
	Email      string `json:"email"`
	Telefono   string `json:"telefono"`
}

func (c *Contacto) StringContacto() string {
	return fmt.Sprintf("Nombre: %-25s\nCumpleaños: %s\nEmail: %s\nTeléfono: %s",
		c.Nombre, c.Cumpleanos, c.Email, c.Telefono)
}
