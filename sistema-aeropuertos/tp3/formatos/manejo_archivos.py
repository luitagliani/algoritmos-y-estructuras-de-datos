import csv

AEROPUERTO_I = 0
AEROPUERTO_J = 1
TIEMPO_PROMEDIO = 2
PRECIO = 3
VUELOS_INTERMEDIOS = 4
CIUDAD = 0
CODIGO_AEROPUERTO = 1
LATITUD = 2
LONGITUD = 3
CIUDAD_I = 0
CIUDAD_J = 1
PRIMERA_LINEA = 0
COMA = ","

ENCABEZADO_XML = '<?xml version="1.0" encoding="UTF-8"?>'
DECLARACION_KML = '<kml xmlns="http://earth.google.com/kml/2.1">'
INICIO_DOCUMENTO = "    <Document>"
ETIQUETA_INICIO_NOMBRE = "        <name>"
ETIQUETA_FIN_NOMBRE = "</name>"
ETIQUETA_INICIO_DESCRIPCION = "        <description>"
ETIQUETA_FIN_DESCRIPCION = "</description>"
ETIQUETA_INICIO_LUGAR = "        <Placemark>"
ETIQUETA_FIN_LUGAR = "        </Placemark>"
ETIQUETA_INICIO_PUNTO = "            <Point>"
ETIQUETA_FIN_PUNTO = "            </Point>"
ETIQUETA_INICIO_CORDS = "                <coordinates>"
ETIQUETA_FIN_CORDS = "</coordinates>"
ETIQUETA_INICIO_LINEA = "            <LineString>"
ETIQUETA_FIN_LINEA = "            </LineString>"
FIN_DOCUMENTO = "    </Document>"
FIN_KML = "</kml>"


def string_a_vuelo(linea_csv):
    return {
        "aeropuerto_i": linea_csv[AEROPUERTO_I],
        "aeropuerto_j": linea_csv[AEROPUERTO_J],
        "tiempo_promedio": float(linea_csv[TIEMPO_PROMEDIO]),
        "precio": float(linea_csv[PRECIO]),
        "vuelos_intermedios": float(linea_csv[VUELOS_INTERMEDIOS]),
    }

def string_a_aeropuerto(linea_csv):
    return {
        "ciudad": linea_csv[CIUDAD],
        "codigo": linea_csv[CODIGO_AEROPUERTO],
        "latitud": float(linea_csv[LATITUD]),
        "longitud": float(linea_csv[LONGITUD]),
    }

def obtener_vuelos(archivo):
    with open(archivo, newline='') as arch:
        lector = csv.reader(arch)
        lineas_csv = list(lector)

    vuelos = []
    for linea in lineas_csv:
        vuelo = string_a_vuelo(linea)
        vuelos.append(vuelo)

    return vuelos

def obtener_aeropuertos(archivo):
    with open(archivo, newline='') as arch:
        lector = csv.reader(arch)
        lineas_csv = list(lector)

    aeropuertos = []
    for linea in lineas_csv:
        aeropuerto = string_a_aeropuerto(linea)
        aeropuertos.append(aeropuerto)

    return aeropuertos

def obtener_info_ciudades(archivo):
    with open(archivo) as arch:
        lineas = arch.readlines()

    primera_linea = lineas[PRIMERA_LINEA].strip()
    ciudades = [c.strip() for c in primera_linea.split(COMA)]

    conexiones = []
    for linea in lineas[1:]:
        partes = [x.strip() for x in linea.strip().split(COMA)]
        if len(partes) < 2:
            continue
        conexiones.append((partes[CIUDAD_I], partes[CIUDAD_J]))

    return ciudades, conexiones

def guardar_encabezado_kml(writer):
    writer.write(f"{ENCABEZADO_XML}\n")
    writer.write(f"{DECLARACION_KML}\n")
    writer.write(f"{INICIO_DOCUMENTO}\n")
    writer.write(f"{ETIQUETA_INICIO_NOMBRE}Mapa Exportado{ETIQUETA_FIN_NOMBRE}\n")

def guardar_puntos(posiciones_de_aeropuertos, writer, recorrido, ruta_vuelos):
    for aeropuerto in ruta_vuelos:
        latitud, longitud = posiciones_de_aeropuertos[aeropuerto]
        recorrido.append(aeropuerto)

        writer.write(f"{ETIQUETA_INICIO_LUGAR}\n")
        writer.write(f"    {ETIQUETA_INICIO_NOMBRE}{aeropuerto}{ETIQUETA_FIN_NOMBRE}\n")
        writer.write(f"{ETIQUETA_INICIO_PUNTO}\n")
        writer.write(f"{ETIQUETA_INICIO_CORDS}{longitud}, {latitud}{ETIQUETA_FIN_CORDS}\n")
        writer.write(f"{ETIQUETA_FIN_PUNTO}\n")
        writer.write(f"{ETIQUETA_FIN_LUGAR}\n")

def guardar_recorridos(posiciones_de_aeropuertos, writer, recorrido):
    for i in range(len(recorrido) - 1):
        aeropuertoActual = recorrido[i]
        aeropuertoSig = recorrido[i+1]

        latitud1, longitud1 = posiciones_de_aeropuertos[aeropuertoActual]
        latitud2, longitud2 = posiciones_de_aeropuertos[aeropuertoSig]

        writer.write(f"{ETIQUETA_INICIO_LUGAR}\n")
        writer.write(f"{ETIQUETA_INICIO_LINEA}\n")
        writer.write(f"{ETIQUETA_INICIO_CORDS}{longitud1}, {latitud1} {longitud2}, {latitud2}{ETIQUETA_FIN_CORDS}\n")
        writer.write(f"{ETIQUETA_FIN_LINEA}\n")
        writer.write(f"{ETIQUETA_FIN_LUGAR}\n")

def guardar_cierre_kml(writer):
    writer.write(f"{FIN_DOCUMENTO}\n")
    writer.write(f"{FIN_KML}")
