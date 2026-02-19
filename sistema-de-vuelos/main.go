package main

import (
	"bufio"
	"os"
	"strings"
	COMANDOS "tp2/comandos"
	SISTEMA "tp2/sistema"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sistemaVuelos := SISTEMA.InicializarSistema()

	for scanner.Scan() {
		linea := scanner.Text()
		input := strings.Fields(linea)
		if len(input) == 0 {
			continue
		}

		comando := input[0]
		args := input[1:]

		switch comando {
		case COMANDOS.AGREGAR_ARCHIVO:
			if len(args) < 1 {
				COMANDOS.ImprimirErrorComando(COMANDOS.AGREGAR_ARCHIVO)
				continue
			}

			for i := 0; i < len(args); i++ {
				COMANDOS.AgregarArchivo(sistemaVuelos, args[i])
			}
		case COMANDOS.VER_TABLERO:
			if len(args) < 4 {
				COMANDOS.ImprimirErrorComando(COMANDOS.VER_TABLERO)
				continue
			}

			COMANDOS.VerTablero(sistemaVuelos, args[0], args[1], args[2], args[3])
		case COMANDOS.INFO_VUELO:
			if len(args) != 1 {
				COMANDOS.ImprimirErrorComando(COMANDOS.INFO_VUELO)
				continue
			}

			COMANDOS.InfoVuelo(sistemaVuelos, args[0])
		case COMANDOS.PRIORIDAD_VUELOS:
			if len(args) != 1 {
				COMANDOS.ImprimirErrorComando(COMANDOS.PRIORIDAD_VUELOS)
				continue
			}

			COMANDOS.PrioridadVuelos(sistemaVuelos, args[0])
		case COMANDOS.SIGUIENTE_VUELO:
			if len(args) != 3 {
				COMANDOS.ImprimirErrorComando(COMANDOS.SIGUIENTE_VUELO)
				continue
			}
			COMANDOS.SiguienteVuelo(sistemaVuelos, args[0], args[1], args[2])
		case COMANDOS.BORRAR:
			if len(args) != 2 {
				COMANDOS.ImprimirErrorComando(COMANDOS.BORRAR)
				continue
			}
			COMANDOS.Borrar(sistemaVuelos, args[0], args[1])
		}
	}
}
