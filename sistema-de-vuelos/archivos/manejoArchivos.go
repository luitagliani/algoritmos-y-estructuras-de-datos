package manejoArchivos

import (
	"encoding/csv"
	"fmt"
	"os"
	FORMATOS "tp2/formatos"
	SISTEMA "tp2/sistema"
)

func ParsearCSV(archivo string) ([]*SISTEMA.Vuelo, error) {
	arch, err := os.Open(archivo)
	if err != nil {
		return nil, fmt.Errorf("no se pudo abrir el archivo: %v", err)

	}
	defer arch.Close()

	csvReader := csv.NewReader(arch)
	lineas_csv, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error leyendo CSV: %v", err)

	}

	var vuelos []*SISTEMA.Vuelo

	for i, linea := range lineas_csv {

		vuelo, err := StringAVuelo(linea)

		if err != nil {
			return nil, fmt.Errorf("línea %d: %w", i+1, err)
		}

		vuelos = append(vuelos, vuelo)
	}

	return vuelos, nil
}

func StringAVuelo(linea_csv []string) (*SISTEMA.Vuelo, error) {
	nro_vuelo, err := FORMATOS.CadenaAEntero(linea_csv[FORMATOS.NUMERO_VUELO])
	if err != nil {
		return nil, fmt.Errorf("número de vuelo inválido: %v", err)
	}

	prioridad, err := FORMATOS.CadenaAEntero(linea_csv[FORMATOS.PRIORIDAD])
	if err != nil {
		return nil, fmt.Errorf("número de prioridad inválido: %v", err)
	}

	fecha, err := FORMATOS.CadenaAFecha(linea_csv[FORMATOS.FECHA])
	if err != nil {
		return nil, fmt.Errorf("número de fecha inválido: %v", err)
	}

	retraso_partida, err := FORMATOS.CadenaAEntero(linea_csv[FORMATOS.RETRASO_PARTIDA])
	if err != nil {
		return nil, fmt.Errorf("número de retraso inválido: %v", err)
	}

	duracion, err := FORMATOS.CadenaAEntero(linea_csv[FORMATOS.DURACION])
	if err != nil {
		return nil, fmt.Errorf("número de duración inválido: %v", err)

	}

	cancelado, err := FORMATOS.CadenaAEntero(linea_csv[FORMATOS.ESTADO_CANCELACION])
	if err != nil {
		return nil, fmt.Errorf("número de cancelación inválido: %v", err)
	}

	return &SISTEMA.Vuelo{
		Nro_vuelo:          nro_vuelo,
		Aerolinea:          linea_csv[FORMATOS.AEROLINEA],
		Aeropuerto_origen:  linea_csv[FORMATOS.AEROPUERTO_ORIGEN],
		Aeropuerto_destino: linea_csv[FORMATOS.AEROPUERTO_DESTINO],
		Matricula:          linea_csv[FORMATOS.MATRICULA],
		Prioridad:          prioridad,
		Fecha:              fecha,
		Retraso_partida:    retraso_partida,
		Duracion:           duracion,
		Cancelado:          cancelado,
	}, nil
}
