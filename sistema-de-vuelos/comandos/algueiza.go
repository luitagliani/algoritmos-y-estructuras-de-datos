package comandos

import (
	"fmt"
	"os"
	"strings"
	manejoArchivos "tp2/archivos"
	FORMATOS "tp2/formatos"
	SISTEMA "tp2/sistema"
)

const (
	ASCENDENTE       = "asc"
	DESCENDENTE      = "desc"
	AGREGAR_ARCHIVO  = "agregar_archivo"
	VER_TABLERO      = "ver_tablero"
	INFO_VUELO       = "info_vuelo"
	PRIORIDAD_VUELOS = "prioridad_vuelos"
	SIGUIENTE_VUELO  = "siguiente_vuelo"
	BORRAR           = "borrar"
)

func ImprimirErrorComando(nombre string) {
	fmt.Fprintf(os.Stderr, "Error en comando %s\n", nombre)
}

func ocurrioUnError(err error, nombreComando string) bool {
	if err != nil {
		if strings.Contains(err.Error(), "inválido") {
			fmt.Fprintln(os.Stderr, err.Error())
		} else {
			ImprimirErrorComando(nombreComando)
		}
		return true
	}
	return false
}

func AgregarArchivo(sistema SISTEMA.Sistema, archivo string) {
	if !strings.HasSuffix(archivo, ".csv") {
		ocurrioUnError(fmt.Errorf("formato de archivo incorrecto"), AGREGAR_ARCHIVO)
		return
	}

	vuelos, err := manejoArchivos.ParsearCSV(archivo)
	if ocurrioUnError(err, AGREGAR_ARCHIVO) {
		return
	}

	for _, v := range vuelos {
		sistema.AgregarVuelo(v)
	}

	fmt.Println("OK")
}

func PrioridadVuelos(sistema SISTEMA.Sistema, cantidadVuelos string) {
	cantidad, err := FORMATOS.CadenaAEntero(cantidadVuelos)
	if ocurrioUnError(err, PRIORIDAD_VUELOS) {
		return
	}

	if cantidad < 0 {
		if ocurrioUnError(fmt.Errorf("cantidad negativa"), PRIORIDAD_VUELOS) {
			return
		}
	}

	vuelos := sistema.ObtenerVuelosPrioritarios(cantidad)

	for i := 0; i < len(vuelos); i++ {
		fmt.Printf("%d - %d\n", vuelos[i].Prioridad, vuelos[i].Nro_vuelo)
	}

	fmt.Println("OK")
}

func ImprimirInfoVuelo(vuelo *SISTEMA.Vuelo) {
	fmt.Printf("%d %s %s %s %s %d %v %d %d %d\n",
		vuelo.Nro_vuelo, vuelo.Aerolinea, vuelo.Aeropuerto_origen,
		vuelo.Aeropuerto_destino, vuelo.Matricula, vuelo.Prioridad,
		vuelo.Fecha.Format(FORMATOS.REFERENCIA), vuelo.Retraso_partida,
		vuelo.Duracion, vuelo.Cancelado)
}

func InfoVuelo(sistema SISTEMA.Sistema, numero string) {
	nroVuelo, err := FORMATOS.CadenaAEntero(numero)
	if ocurrioUnError(err, INFO_VUELO) {
		return
	}

	vuelo := sistema.ObtenerInfoVuelo(nroVuelo)
	if vuelo == nil {
		if ocurrioUnError(fmt.Errorf("vuelo inexistente"), INFO_VUELO) {
			return
		}
	}

	ImprimirInfoVuelo(vuelo)
	fmt.Println("OK")
}

func SiguienteVuelo(sistema SISTEMA.Sistema, origen, destino, fecha string) {
	vuelo := sistema.BuscarSiguienteVuelo(origen, destino, fecha)

	if vuelo == nil {
		fmt.Printf("No hay vuelo registrado desde %s hacia %s desde %s\n", origen, destino, fecha)
	} else {
		ImprimirInfoVuelo(vuelo)
	}

	fmt.Println("OK")
}

func VerTablero(sistema SISTEMA.Sistema, cantidad, modo, desde, hasta string) {
	cant, err := FORMATOS.CadenaAEntero(cantidad)
	if ocurrioUnError(err, VER_TABLERO) {
		return
	}

	fechaDesde, err := FORMATOS.CadenaAFecha(desde)
	if ocurrioUnError(err, VER_TABLERO) {
		return
	}

	fechaHasta, err := FORMATOS.CadenaAFecha(hasta)
	if ocurrioUnError(err, VER_TABLERO) {
		return
	}

	if cant <= 0 || fechaDesde.After(fechaHasta) {
		ImprimirErrorComando(VER_TABLERO)
		return
	}

	var asc bool

	switch modo {
	case ASCENDENTE:
		asc = true
	case DESCENDENTE:
		asc = false
	default:
		ImprimirErrorComando(VER_TABLERO)
		return
	}

	vuelos := sistema.ObtenerDatosVueloPorFecha(fechaDesde, fechaHasta, cant, asc)

	if cant > len(vuelos) {
		cant = len(vuelos)
	}

	if !asc {
		invertirVuelos(&vuelos)
	}

	for i := 0; i < cant; i++ {
		fmt.Printf("%s - %s\n", vuelos[i].Fecha, vuelos[i].Numero)
	}

	fmt.Println("OK")
}

func invertirVuelos(vuelo *[]SISTEMA.VueloPorFecha) {
	for i, j := 0, len((*vuelo))-1; i < j; i, j = i+1, j-1 {
		(*vuelo)[i], (*vuelo)[j] = (*vuelo)[j], (*vuelo)[i]
	}
}

func Borrar(sistema SISTEMA.Sistema, desdeStr, hastaStr string) {
	fechaDesde, err := FORMATOS.CadenaAFecha(desdeStr)
	if ocurrioUnError(err, BORRAR) {
		return
	}

	fechaHasta, err := FORMATOS.CadenaAFecha(hastaStr)
	if ocurrioUnError(err, BORRAR) {
		return
	}

	if fechaDesde.After(fechaHasta) {
		ImprimirErrorComando(BORRAR)
		return
	}

	vuelos := sistema.BorrarRango(fechaDesde, fechaHasta)

	for i := 0; i < len(vuelos); i++ {
		ImprimirInfoVuelo(vuelos[i])
	}

	fmt.Println("OK")
}
