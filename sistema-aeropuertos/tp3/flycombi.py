#!/usr/bin/env python3

import sys
from mapa.mapa import Mapa

CAMINO_MAS = "camino_mas"
CANT_ARGS_CAMINO_MAS = 3
CAMINO_ESCALAS = "camino_escalas"
CANT_ARGS_CAMINO_ESCALAS = 2
CENTRALIDAD = "centralidad"
CANT_ARGS_CENTRALIDAD = 1
NUEVA_AEROLINEA = "nueva_aerolinea"
CANT_ARGS_NUEVA_AEROLINEA = 1
ITINERARIO = "itinerario"
CANT_ARGS_ITINERARIO = 1
EXPORTAR_KML = "exportar_kml"
CANT_ARGS_EXPORTAR_KML = 1
COMA = ","
ESPACIO = " "
ARGS_ARCHIVO1 = 1
ARGS_ARCHIVO2 = 2

def main():
    if len(sys.argv) != 3:
        print("Formato incorrecto: acordate que es ./ejecutable <csv1> <csv2>")
        return

    archivo1, archivo2 = sys.argv[ARGS_ARCHIVO1], sys.argv[ARGS_ARCHIVO2]

    if not archivo1.endswith(".csv") or not archivo2.endswith(".csv"):
        print("Formato incorrecto: los archivos deben ser .csv")
        return

    mapa = Mapa.crear_mapa(archivo1, archivo2)
    ruta_en_cache = []
    
    for linea in sys.stdin:
        linea = linea.strip()
        if not linea:
            continue

        campos = linea.split(ESPACIO, 1)
        if len(campos) < 2:
            continue

        comando, args_cadena = campos
        args = [arg.strip() for arg in args_cadena.split(COMA)]

        if comando == CAMINO_MAS:
            if len(args) != CANT_ARGS_CAMINO_MAS:
                print("Uso esperado: camino_mas <modo>,<origen>,<destino>", file=sys.stderr)
                continue
            modo, origen, destino = args
            ruta_en_cache = mapa.camino_mas(modo, origen, destino)

        elif comando == CAMINO_ESCALAS:
            if len(args) != CANT_ARGS_CAMINO_ESCALAS:
                print("Uso esperado: camino_escalas <origen>,<destino>", file=sys.stderr)
                continue
            origen, destino = args
            ruta_en_cache = mapa.camino_escalas(origen, destino)

        elif comando == CENTRALIDAD:
            if len(args) != CANT_ARGS_CENTRALIDAD:
                print("Uso esperado: centralidad <cantidad>", file=sys.stderr)
                continue
            cantidad = args[0]
            mapa.mas_centrales(int(cantidad))

        elif comando == NUEVA_AEROLINEA:
            if len(args) != CANT_ARGS_NUEVA_AEROLINEA:
                print("Uso esperado: nueva_aerolinea <ruta>", file=sys.stderr)
                continue
            ruta = args[0]
            mapa.nueva_aerolinea(ruta)
        elif comando == ITINERARIO:
            if len(args) != CANT_ARGS_ITINERARIO:
                print("Uso esperado: itinerario <ruta>", file=sys.stderr)
                continue
            ruta = args[0]
            mapa.itinerario(ruta)
        elif comando == EXPORTAR_KML:
            if len(args) != CANT_ARGS_EXPORTAR_KML:
                print("Uso esperado: exportar_kml <ruta>", file=sys.stderr)
                continue
            ruta = args[0]
            mapa.exportar_kml(ruta, ruta_en_cache)

if __name__ == "__main__":
    try:
        main()
    except KeyboardInterrupt:
        print("\nPrograma interrumpido. Chau :)")