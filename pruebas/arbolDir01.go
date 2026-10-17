package pruebas

import (
	"fmt"

	"github.com/goforj/godump"
)

type Carpeta struct {
	Nombre string
	Hijos  []*Carpeta
}

func (c *Carpeta) AgregarHijo(hijo *Carpeta) {
	c.Hijos = append(c.Hijos, hijo)
}

func main() {
	raiz := &Carpeta{Nombre: "fotos"}

	hijo1 := &Carpeta{Nombre: "viajes"}
	hijo2 := &Carpeta{Nombre: "cumpleaños"}
	raiz.AgregarHijo(hijo1)
	raiz.AgregarHijo(hijo2)
	for _, hijo := range raiz.Hijos {
		fmt.Println(hijo.Nombre)
	}

	godump.Dump(raiz)
}
