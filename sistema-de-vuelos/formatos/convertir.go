package formatos

import (
	"strconv"
	"time"
)

const (
	REFERENCIA = "2006-01-02T15:04:05"
)

func CadenaAEntero(cadena string) (int, error) {
	return strconv.Atoi(cadena)
}

func CadenaAFecha(cadena string) (time.Time, error) {
	return time.Parse(REFERENCIA, cadena)
}

func FechaACadena(fecha time.Time) string {
	return fecha.Format(REFERENCIA)
}

func EnteroACadena(entero int) string {
	return strconv.Itoa(entero)
}
