package main

import (
	"strings"

	"github.com/goforj/godump"
)

type Carpeta struct {
	Nombre string
	Hijos  []*Carpeta
}

func (c *Carpeta) ObtenerOCrearHijo(nombre string) *Carpeta {
	// 1. ¿Ya existe este hijo en la lista?
	for _, hijo := range c.Hijos {
		if hijo.Nombre == nombre {
			return hijo // ¡Existe! Devolvemos el puntero a ese hijo existente
		}
	}

	// 2. Si no existe, creamos uno nuevo con &
	nuevo := &Carpeta{Nombre: nombre}

	// 3. Lo agregamos como hijo de la carpeta actual
	c.Hijos = append(c.Hijos, nuevo)

	// 4. Devolvemos el puntero al nuevo hijo
	return nuevo
}

func main() {
	raiz := &Carpeta{Nombre: "fotos"}
	cursor := raiz

	ruta := "viajes/2024/verano"
	partesRuta := strings.Split(ruta, "/")
	for _, parte := range partesRuta {
		cursor = cursor.ObtenerOCrearHijo(parte)
	}

	godump.Dump(raiz)
}
