package comandos

type Comandos interface {

	// Procesa de forma completa un archivo de .csv que contiene datos de vuelos.
	AgregarArchivo(string)

	// Muestra los K vuelos ordenados por fecha de forma ascendente (asc) o descendente (desc), cuya fecha de despegue esté dentro de el
	// intervalo <desde> <hasta> (inclusive).
	VerTablero(string, string, string, string)

	// Muestra toda la información posible en sobre el vuelo que tiene el código pasado por parámetro.
	InfoVuelo(string)

	// Muestra los códigos de los K vuelos que tienen mayor prioridad.
	PrioridadVuelos(string)

	// Muestra la información del vuelo (tal cual en info_vuelo) del próximo vuelo directo que conecte los aeropuertos de origen y destino,
	// a partir de la fecha indicada (inclusive). Si no hay un siguiente vuelo cargado, imprimir No hay vuelo registrado desde <aeropuerto origen>
	// hacia <aeropuerto destino> desde <fecha> (con los valores que correspondan).
	SiguienteVuelo(string, string, string)

	// Borra todos los vuelos cuya fecha de despegue estén dentro del intervalo <desde> <hasta> (inclusive).
	Borrar(string, string)
}
