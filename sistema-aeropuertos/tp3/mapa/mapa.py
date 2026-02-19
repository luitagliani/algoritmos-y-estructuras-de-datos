from grafo.grafo import Grafo
from formatos.manejo_archivos import obtener_vuelos, obtener_aeropuertos, obtener_info_ciudades
from grafos_lib.camino_minimo import dijkstra, bfs
from grafos_lib.centralidad import centralidad
from grafos_lib.datos_adicionales import obtener_aristas
from grafos_lib.mst import mst_prim
from grafos_lib.orden import orden_topologico
from formatos.manejo_archivos import guardar_encabezado_kml, guardar_puntos, guardar_recorridos, guardar_cierre_kml

import heapq
import csv
import sys
from io import StringIO

CIUDAD = "ciudad"
CODIGO = "codigo"
LATITUD = "latitud"
LONGITUD = "longitud"
AEROPUERTO_I = "aeropuerto_i"
AEROPUERTO_J = "aeropuerto_j"
TIEMPO_PROMEDIO = "tiempo_promedio"
PRECIO = "precio"
VUELOS_INTERMEDIOS = "vuelos_intermedios"
BARATO = "barato"
RAPIDO = "rapido"
COMA_ESPACIO = ", "
PRIMERA_CELDA = 0
SEGUNDA_CELDA = 1
TERCERA_CELDA = 2

class Mapa:
    def __init__(self, grafo_precio, grafo_tiempo, grafo_frecuencia, aeropuertos_ciudad, datos_vuelo, posiciones_aeropuertos):
        self.grafo_por_precio = grafo_precio
        self.grafo_por_tiempo = grafo_tiempo
        self.grafo_por_frecuencia = grafo_frecuencia
        self.aeropuertos_por_ciudad = aeropuertos_ciudad
        self.datos_de_vuelo = datos_vuelo
        self.posiciones_de_aeropuertos = posiciones_aeropuertos

    def crear_mapa(archivo_aeropuertos, archivo_vuelos):
        aeropuertos = obtener_aeropuertos(archivo_aeropuertos)
        vuelos = obtener_vuelos(archivo_vuelos)

        grafo_precio = Grafo(False)
        grafo_tiempo = Grafo(False)
        grafo_frecuencia = Grafo(False)
        aeropuertos_de_ciudad = {}
        posiciones_aeropuertos = {}
        datos_vuelo = {}

        for aeropuerto in aeropuertos:
            ciudad = aeropuerto[CIUDAD]
            codigo = aeropuerto[CODIGO]
            latitud = aeropuerto[LATITUD]
            longitud = aeropuerto[LONGITUD]

            posiciones_aeropuertos[codigo] = (latitud, longitud)
            grafo_precio.agregar_vertice(codigo)
            grafo_tiempo.agregar_vertice(codigo)
            grafo_frecuencia.agregar_vertice(codigo)

            if ciudad not in aeropuertos_de_ciudad:
                aeropuertos_de_ciudad[ciudad] = []
            aeropuertos_de_ciudad[ciudad].append(codigo)


        for vuelo in vuelos:
            aeropuerto_i = vuelo[AEROPUERTO_I]
            aeropuerto_j = vuelo[AEROPUERTO_J]
            tiempo_promedio = vuelo[TIEMPO_PROMEDIO]
            precio = vuelo[PRECIO]
            vuelos_intermedios = vuelo[VUELOS_INTERMEDIOS]

            grafo_precio.agregar_arista(aeropuerto_i, aeropuerto_j, precio)
            grafo_tiempo.agregar_arista(aeropuerto_i, aeropuerto_j, tiempo_promedio)

            if vuelos_intermedios > 0:
                grafo_frecuencia.agregar_arista(aeropuerto_i, aeropuerto_j, 1/vuelos_intermedios)

            clave_vuelo = f"{aeropuerto_i}-{aeropuerto_j}"
            datos_vuelo[clave_vuelo] = (vuelo[TIEMPO_PROMEDIO], vuelo[VUELOS_INTERMEDIOS])

        return Mapa(grafo_precio, grafo_tiempo, grafo_frecuencia, aeropuertos_de_ciudad, datos_vuelo, posiciones_aeropuertos)


    def camino_mas(self, modo, origen, destino):
        if modo == BARATO:
            grafo = self.grafo_por_precio
        elif modo == RAPIDO:
            grafo = self.grafo_por_tiempo
        else:
            print("Modo inválido")
            return

        if origen not in self.aeropuertos_por_ciudad or destino not in self.aeropuertos_por_ciudad:
            print("Ciudad inválida")
            return

        aeropuertos_origen = self.aeropuertos_por_ciudad[origen]
        aeropuertos_destino = self.aeropuertos_por_ciudad[destino]

        mejor_camino = None
        mejor_valor = None
        encontrado = False

        for aero_origen in aeropuertos_origen:
            distancias, padres = dijkstra(grafo, aero_origen)

            for aero_destino in aeropuertos_destino:
                camino_actual = hallar_camino(aero_destino, padres)
                valor_actual = distancias[aero_destino]

                if not encontrado or valor_actual < mejor_valor:
                    mejor_camino = camino_actual
                    mejor_valor = valor_actual
                    encontrado = True

        if not encontrado:
            print("No se encontró camino entre las ciudades")
            return

        print(" -> ".join(mejor_camino))
        return mejor_camino

    def camino_escalas(self, origen, destino):
        if origen not in self.aeropuertos_por_ciudad or destino not in self.aeropuertos_por_ciudad:
            print("Ciudad inválida")
            return

        mejor_camino = None
        largo_min = float('inf')
        aeropuertos_origen = self.aeropuertos_por_ciudad[origen]
        aeropuertos_destino = self.aeropuertos_por_ciudad[destino]

        for aeropuerto_origen in aeropuertos_origen:
            distancias, padres = bfs(self.grafo_por_tiempo, aeropuerto_origen)
            for aeropuerto_destino in aeropuertos_destino:
                if aeropuerto_destino in distancias and distancias[aeropuerto_destino] < largo_min:
                    largo_min = distancias[aeropuerto_destino]
                    mejor_camino = hallar_camino(aeropuerto_destino, padres)

        if mejor_camino is None:
            print("No se encontró camino entre las ciudades")
        else:
            print(" -> ".join(mejor_camino))

        return mejor_camino


    def mas_centrales(self, cantidad):
        if cantidad <= 0:
            print("Cantidad inválida")
            return
        
        centrales = centralidad(self.grafo_por_frecuencia)
        centralesOrdenados = top_centrales(centrales, cantidad)

        print(COMA_ESPACIO.join(centralesOrdenados))

    def itinerario(self, archivo):
        ciudades, conexiones = obtener_info_ciudades(archivo)
        grafo_temp = Grafo(True)

        for ciudad in ciudades:
            grafo_temp.agregar_vertice(ciudad)

        for conexion in conexiones:
            ciudad_i = conexion[PRIMERA_CELDA]
            ciudad_j = conexion[SEGUNDA_CELDA]
            grafo_temp.agregar_arista(ciudad_i, ciudad_j, 1)
        
        orden = orden_topologico(grafo_temp)

        if len(orden) != grafo_temp.cantidad():
            print("El grafo tiene ciclos, no se puede realizar un itinerario")
            return
        
        print(", ".join(orden))

        for i in range(len(orden) - 1):
            origen = orden[i]
            destino = orden[i + 1]
            self.camino_mas(RAPIDO, origen, destino)

    def nueva_aerolinea(mapa, ruta_salida):
        try:
            with open(ruta_salida, mode='w', newline='', encoding='utf-8') as archivo_salida:
                writer = csv.writer(archivo_salida)

                grafo_nueva_aerolinea = mst_prim(mapa.grafo_por_precio)
                vuelos_necesarios = obtener_aristas(grafo_nueva_aerolinea)

                for vuelo in vuelos_necesarios:
                    aeropuerto_i = vuelo[PRIMERA_CELDA]
                    aeropuerto_j = vuelo[SEGUNDA_CELDA]
                    peso = vuelo[TERCERA_CELDA] 

                    clave_vuelo = f"{aeropuerto_i}-{aeropuerto_j}"
                    if clave_vuelo not in mapa.datos_de_vuelo:
                        clave_vuelo = f"{aeropuerto_j}-{aeropuerto_i}"

                    tiempo_promedio, cant_vuelos_intermedios = mapa.datos_de_vuelo[clave_vuelo]

                    linea = [
                        aeropuerto_i,
                        aeropuerto_j,
                        str(int(tiempo_promedio)),
                        str(int(peso)),
                        str(int(cant_vuelos_intermedios))
                    ]

                    writer.writerow(linea)
            print("OK")

        except Exception as e:
            print("Error al escribir linea:", e)
    
    def exportar_kml(self, ruta_salida, ruta_vuelos):
        if not ruta_vuelos:
            print("No hay un camino para exportar", file=sys.stderr)
            return
        
        writer = StringIO()
        recorrido = []
        
        guardar_encabezado_kml(writer)
        
        guardar_puntos(self.posiciones_de_aeropuertos, writer, recorrido, ruta_vuelos)
        guardar_recorridos(self.posiciones_de_aeropuertos, writer, recorrido)
        
        guardar_cierre_kml(writer)
        
        try:
            with open(ruta_salida, 'w', encoding='utf-8') as f:
                f.write(writer.getvalue())
            print("OK")
        except IOError as e:
            print(f"Error al escribir archivo: {e}", file=sys.stderr)

def top_centrales(dicCentrales, cantidad):
    heapCentrales = []
    for aeropuerto, centralidad in dicCentrales.items():
        if len(heapCentrales) < cantidad:
            heapq.heappush(heapCentrales, (centralidad, aeropuerto))
        elif centralidad > heapCentrales[0][0]:
            heapq.heappop(heapCentrales)
            heapq.heappush(heapCentrales, (centralidad, aeropuerto))
    
    centrales = []
    
    while len(heapCentrales) > 0:
        _, aeropuerto = heapq.heappop(heapCentrales)
        centrales.append(aeropuerto)
    
    return centrales[::-1]

def hallar_camino(destino, padres):
    camino = []
    actual = destino
    while actual != None:
        camino.append(actual)
        actual = padres[actual]
    return camino[::-1]
