# Trabajo práctico - sistema de análisis de redes aeroportuarias

Este proyecto es un sistema de análisis y optimización de rutas de vuelo desarrollado en **Python**. Utiliza **Teoría de Grafos** para modelar la red de aeropuertos y conexiones, permitiendo resolver problemas complejos de logística, rutas óptimas y análisis de importancia en la red.

## ⚙️ Funcionalidades y Algoritmos Implementados

El sistema modela los datos a través de grafos pesados y dirigidos, implementando los siguientes algoritmos clásicos:

* **Búsqueda de Rutas Óptimas:** Encuentra el vuelo más barato o el más rápido entre dos ciudades utilizando el algoritmo de **Dijkstra**.
* **Minimización de Escalas:** Calcula la ruta con la menor cantidad de conexiones intermedias aplicando Búsqueda en Anchura (**BFS**).
* **Análisis de Importancia:** Determina cuáles son los aeropuertos más cruciales de la red calculando la **Centralidad** del grafo y filtrando el top utilizando **Heaps (Colas de Prioridad)**.
* **Planificación de Itinerarios:** Genera un recorrido lógico para visitar múltiples ciudades sin ciclos, utilizando **Orden Topológico**.
* **Diseño de Nueva Aerolínea:** Calcula la red de rutas más económica que conecte todos los aeropuertos sin redundancias, implementando el algoritmo de **Prim (Árbol de Tendido Mínimo / MST)**.
* **Exportación Geoespacial:** Permite exportar las rutas calculadas a formato `.kml` para su visualización en herramientas como Google Earth o Google Maps.

##Tecnologías y Arquitectura
* **Lenguaje:** Python 3
* **Estructuras de Datos:** Grafos (implementación propia), Heaps, Diccionarios.
* **Manejo de Archivos:** Lectura de datasets en CSV y generación dinámica de archivos KML.

## Uso del Sistema

El sistema procesa bases de datos de aeropuertos y vuelos al inicializarse. Dependiendo de los comandos ingresados, el mapa de conexiones se evalúa por precio, tiempo o frecuencia para devolver la respuesta computacionalmente más eficiente.
