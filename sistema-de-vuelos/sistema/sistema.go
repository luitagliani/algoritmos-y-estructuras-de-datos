package sistema

import (
	"time"
)

type Sistema interface {
	AgregarVuelo(*Vuelo)
	ObtenerVuelosPrioritarios(int) []*Vuelo
	ObtenerInfoVuelo(int) *Vuelo
	BuscarSiguienteVuelo(string, string, string) *Vuelo
	ObtenerDatosVueloPorFecha(time.Time, time.Time, int, bool) []VueloPorFecha
	BorrarRango(time.Time, time.Time) []*Vuelo
}
