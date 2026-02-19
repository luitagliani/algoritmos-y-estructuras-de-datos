package sistema

import (
	"strings"
	TDAHeap "tdas/cola_prioridad"
	TDADiccionario "tdas/diccionario"
	TDALista "tdas/lista"
	"time"
	FORMATOS "tp2/formatos"
)

type VueloPorFecha struct {
	Fecha  string
	Numero string
}

type Vuelo struct {
	Nro_vuelo          int
	Aerolinea          string
	Aeropuerto_origen  string
	Aeropuerto_destino string
	Matricula          string
	Prioridad          int
	Fecha              time.Time
	Retraso_partida    int
	Duracion           int
	Cancelado          int
}

type ListaDeVuelos = TDALista.Lista[*Vuelo]
type VuelosPorFechaLista = TDADiccionario.DiccionarioOrdenado[time.Time, ListaDeVuelos]
type VuelosPorDestino = TDADiccionario.Diccionario[string, VuelosPorFechaLista]
type VuelosPorOrigen = TDADiccionario.Diccionario[string, VuelosPorDestino]

type sistemaVuelos struct {
	vuelosPorClave  TDADiccionario.Diccionario[int, *Vuelo]
	vuelosPorFecha  TDADiccionario.DiccionarioOrdenado[time.Time, TDADiccionario.DiccionarioOrdenado[string, *Vuelo]]
	vuelosPorCamino VuelosPorOrigen
	cantidadVuelos  int
}

const (
	_PRIMERO_MAYOR = 1
	_PRIMERO_MENOR = -1
)

func compararPrioridad(vuelo1, vuelo2 *Vuelo) int {
	prioridad1 := vuelo1.Prioridad
	numeroVuelo1 := FORMATOS.EnteroACadena(vuelo1.Nro_vuelo)
	prioridad2 := vuelo2.Prioridad
	numeroVuelo2 := FORMATOS.EnteroACadena(vuelo2.Nro_vuelo)

	if prioridad1 > prioridad2 || (prioridad1 == prioridad2 && strings.Compare(numeroVuelo1, numeroVuelo2) < 0) {
		return _PRIMERO_MAYOR
	}

	if prioridad1 < prioridad2 || (strings.Compare(numeroVuelo1, numeroVuelo2) > 0) {
		return _PRIMERO_MENOR
	}

	return prioridad1 - prioridad2
}

func compararFechas(a, b time.Time) int {
	if a.After(b) {
		return _PRIMERO_MAYOR
	}

	if b.After(a) {
		return _PRIMERO_MENOR
	}

	return 0
}

func compNroVuelo(a, b string) int {
	return strings.Compare(a, b)
}

func InicializarSistema() Sistema {
	vueloClave := TDADiccionario.CrearHash[int, *Vuelo]()
	vueloFecha := TDADiccionario.CrearABB[time.Time, TDADiccionario.DiccionarioOrdenado[string, *Vuelo]](compararFechas)
	vueloCamino := TDADiccionario.CrearHash[string, TDADiccionario.Diccionario[string, TDADiccionario.DiccionarioOrdenado[time.Time, TDALista.Lista[*Vuelo]]]]()

	return &sistemaVuelos{
		vuelosPorClave:  vueloClave,
		vuelosPorFecha:  vueloFecha,
		vuelosPorCamino: vueloCamino,
	}
}

func (sistema *sistemaVuelos) AgregarVueloPorCamino(vuelo *Vuelo) {
	origen := vuelo.Aeropuerto_origen
	destino := vuelo.Aeropuerto_destino
	fecha := vuelo.Fecha

	if !sistema.vuelosPorCamino.Pertenece(origen) {
		sistema.vuelosPorCamino.Guardar(origen, TDADiccionario.CrearHash[string, TDADiccionario.DiccionarioOrdenado[time.Time, TDALista.Lista[*Vuelo]]]())
	}

	dicDestino := sistema.vuelosPorCamino.Obtener(origen)

	if !dicDestino.Pertenece(destino) {
		dicDestino.Guardar(destino, TDADiccionario.CrearABB[time.Time, TDALista.Lista[*Vuelo]](compararFechas))
	}

	dicFecha := dicDestino.Obtener(destino)

	if !dicFecha.Pertenece(fecha) {
		listaVuelos := TDALista.CrearListaEnlazada[*Vuelo]()
		dicFecha.Guardar(fecha, listaVuelos)
	}

	listaVuelos := dicFecha.Obtener(fecha)
	listaVuelos.InsertarUltimo(vuelo)
}

func (sistema *sistemaVuelos) agregarVueloPorFecha(vuelo *Vuelo) {
	if !sistema.vuelosPorFecha.Pertenece(vuelo.Fecha) {
		sistema.vuelosPorFecha.Guardar(vuelo.Fecha, TDADiccionario.CrearABB[string, *Vuelo](compNroVuelo))
	}

	diccVuelos := sistema.vuelosPorFecha.Obtener(vuelo.Fecha)
	nro := FORMATOS.EnteroACadena(vuelo.Nro_vuelo)

	diccVuelos.Guardar(nro, vuelo)
}

func (sistema *sistemaVuelos) ActualizarVueloPorFecha(vueloNuevo *Vuelo, fechaVieja time.Time) {
	nro := FORMATOS.EnteroACadena(vueloNuevo.Nro_vuelo)

	if fechaVieja != vueloNuevo.Fecha {
		if sistema.vuelosPorFecha.Pertenece(fechaVieja) {
			diccViejo := sistema.vuelosPorFecha.Obtener(fechaVieja)
			if diccViejo.Pertenece(nro) {
				diccViejo.Borrar(nro)
			}
		}
	}

	if !sistema.vuelosPorFecha.Pertenece(vueloNuevo.Fecha) {
		sistema.vuelosPorFecha.Guardar(vueloNuevo.Fecha, TDADiccionario.CrearABB[string, *Vuelo](compNroVuelo))
	}

	diccNuevo := sistema.vuelosPorFecha.Obtener(vueloNuevo.Fecha)
	diccNuevo.Guardar(nro, vueloNuevo)
}

func (sistema *sistemaVuelos) eliminarVueloPorCamino(vuelo *Vuelo) {
	origen := vuelo.Aeropuerto_origen

	if !sistema.vuelosPorCamino.Pertenece(origen) {
		return
	}

	destino := vuelo.Aeropuerto_destino
	fecha := vuelo.Fecha

	dicDestino := sistema.vuelosPorCamino.Obtener(origen)
	if !dicDestino.Pertenece(destino) {
		return
	}

	dicFecha := dicDestino.Obtener(destino)
	if !dicFecha.Pertenece(fecha) {
		return
	}

	listaVuelos := dicFecha.Obtener(fecha)

	iter := listaVuelos.Iterador()
	for iter.HaySiguiente() {
		v := iter.VerActual()
		if v == vuelo {
			iter.Borrar()
			break
		}
		iter.Siguiente()
	}

	if listaVuelos.EstaVacia() {
		dicFecha.Borrar(fecha)
	}
}

func (sistema *sistemaVuelos) actualizarVueloPorCamino(vueloViejo *Vuelo, vueloNuevo *Vuelo) {
	sistema.eliminarVueloPorCamino(vueloViejo)
	sistema.AgregarVueloPorCamino(vueloNuevo)
}

func (sistema *sistemaVuelos) AgregarVuelo(vuelo *Vuelo) {
	vueloExistente := sistema.vuelosPorClave.Pertenece(vuelo.Nro_vuelo)

	if vueloExistente {
		//si el vuelo que se quiere agregar ya pertenece al sistema, lo vamos a tratar como "vuelo viejo" ya que probablemente será actualizado
		vueloViejo := sistema.vuelosPorClave.Obtener(vuelo.Nro_vuelo)

		//si se cumple cualquiera de estas condiciones, tengo que actualizar el diccionario de vuelos por camino
		if vueloViejo.Aeropuerto_origen != vuelo.Aeropuerto_origen ||
			vueloViejo.Aeropuerto_destino != vuelo.Aeropuerto_destino ||
			vueloViejo.Fecha != vuelo.Fecha {
			sistema.actualizarVueloPorCamino(vueloViejo, vuelo)
		}

		//siempre me aseguro de actualizar el diccionario de vuelos por fecha, por si cambió algo
		sistema.ActualizarVueloPorFecha(vuelo, vueloViejo.Fecha)

	} else {
		//sino, es un vuelo nuevo y lo actualizo normalmente
		sistema.agregarVueloPorFecha(vuelo)
		sistema.AgregarVueloPorCamino(vuelo)
		sistema.cantidadVuelos++
	}

	//vuelos por clave, al guardar toda la estructura del vuelo y no tener un orden, si es nuevo se guarda y si ya pertenecía se actualiza
	sistema.vuelosPorClave.Guardar(vuelo.Nro_vuelo, vuelo)
}

func (sistema *sistemaVuelos) crearArrayDeVuelos() []*Vuelo {
	vuelosTotales := make([]*Vuelo, 0)
	iterFechas := sistema.vuelosPorFecha.Iterador()

	for iterFechas.HaySiguiente() {
		_, diccVuelos := iterFechas.VerActual()

		iterVuelos := diccVuelos.Iterador()
		for iterVuelos.HaySiguiente() {
			_, vuelo := iterVuelos.VerActual()
			vuelosTotales = append(vuelosTotales, vuelo)
			iterVuelos.Siguiente()
		}

		iterFechas.Siguiente()
	}

	return vuelosTotales
}

func (sistema *sistemaVuelos) ObtenerVuelosPrioritarios(cantidadAMostrar int) []*Vuelo {
	vuelosTotales := sistema.crearArrayDeVuelos()
	heapPriotarios := TDAHeap.CrearHeapArr(vuelosTotales, compararPrioridad)
	var vuelosOrdenados []*Vuelo

	for j := 0; j < cantidadAMostrar && !heapPriotarios.EstaVacia(); j++ {
		vuelosOrdenados = append(vuelosOrdenados, heapPriotarios.Desencolar())
	}

	return vuelosOrdenados
}

func (sistema *sistemaVuelos) ObtenerInfoVuelo(clave_vuelo int) *Vuelo {
	if !sistema.vuelosPorClave.Pertenece(clave_vuelo) {
		return nil
	}

	vuelo := sistema.vuelosPorClave.Obtener(clave_vuelo)

	return vuelo
}

func (sistema *sistemaVuelos) BuscarSiguienteVuelo(origen, destino, fecha string) *Vuelo {
	if !sistema.vuelosPorCamino.Pertenece(origen) {
		return nil
	}
	dicOrigen := sistema.vuelosPorCamino.Obtener(origen)

	if !dicOrigen.Pertenece(destino) {
		return nil
	}

	dicDestino := dicOrigen.Obtener(destino)
	fechaBuscada, _ := time.Parse(FORMATOS.REFERENCIA, fecha)
	iterador := dicDestino.IteradorRango(&fechaBuscada, nil)

	if iterador.HaySiguiente() {
		_, listaVuelos := iterador.VerActual()
		if listaVuelos != nil && !listaVuelos.EstaVacia() {
			vuelo := listaVuelos.VerPrimero()
			return vuelo
		}
	}

	return nil
}

func crearVueloPorFecha(fecha time.Time, nro int) VueloPorFecha {
	return VueloPorFecha{
		Fecha:  FORMATOS.FechaACadena(fecha),
		Numero: FORMATOS.EnteroACadena(nro),
	}
}

func (sistema *sistemaVuelos) ObtenerDatosVueloPorFecha(fechaDesde, fechaHasta time.Time, cantVuelosAMostrar int, ascendente bool) []VueloPorFecha {
	iteradorVuelosPorFecha := sistema.vuelosPorFecha.IteradorRango(&fechaDesde, &fechaHasta)
	var vuelos []VueloPorFecha

	for iteradorVuelosPorFecha.HaySiguiente() {
		fecha, dicVuelos := iteradorVuelosPorFecha.VerActual()
		iteradorDic := dicVuelos.Iterador()

		for iteradorDic.HaySiguiente() {
			_, vuelo := iteradorDic.VerActual()
			datosActual := crearVueloPorFecha(fecha, vuelo.Nro_vuelo)

			vuelos = append(vuelos, datosActual)
			iteradorDic.Siguiente()
		}
		iteradorVuelosPorFecha.Siguiente()
	}
	return vuelos
}

func (sistema *sistemaVuelos) eliminarVuelo(vuelo *Vuelo) {
	sistema.vuelosPorClave.Borrar(vuelo.Nro_vuelo)
	sistema.eliminarVueloPorCamino(vuelo)
	sistema.cantidadVuelos--
}

func (sistema *sistemaVuelos) BorrarRango(fechaDesde, fechaHasta time.Time) []*Vuelo {
	iteradorABB := sistema.vuelosPorFecha.IteradorRango(&fechaDesde, &fechaHasta)

	var fechasABorrar []time.Time
	var vuelosBorrados []*Vuelo

	for iteradorABB.HaySiguiente() {
		fecha, listaVuelo := iteradorABB.VerActual()
		iteradorLista := listaVuelo.Iterador()

		for iteradorLista.HaySiguiente() {
			_, vueloBorrado := iteradorLista.VerActual()
			sistema.eliminarVuelo(vueloBorrado)
			vuelosBorrados = append(vuelosBorrados, vueloBorrado)

			iteradorLista.Siguiente()
		}

		fechasABorrar = append(fechasABorrar, fecha)
		iteradorABB.Siguiente()
	}

	for _, fecha := range fechasABorrar {
		sistema.vuelosPorFecha.Borrar(fecha)
	}

	return vuelosBorrados
}
