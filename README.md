# Algoritmos y Estructuras de Datos

Este repositorio agrupa una serie de proyectos enfocados en el modelado de datos, la eficiencia algorítmica y la aplicación de la Teoría de Grafos para la resolución de problemas complejos. Las soluciones están desarrolladas en **Go** y **Python**.

# Estructura del Repositorio

El código está dividido en tres grandes módulos prácticos:

### 1. TDAs (Go)
Implementación desde cero de estructuras de datos fundamentales para el manejo eficiente de memoria y referencias.
* **Estructuras incluidas:** Pilas (Stack), Listas Enlazadas y Diccionarios (Hash Maps).
### 2. Sistema de Gestión de Vuelos (Go)
Sistema de administración y búsqueda que maneja información temporal y colas de prioridad.
* **Funcionalidades:** Registro de rutas, búsqueda del próximo vuelo disponible, filtrado por rangos de fechas y gestión de prioridades.
* **Testing Automatizado:** Incluye un script de validación en **Bash** (`pruebas.sh`) que inyecta casos de prueba, compara salidas estándar (`stdout`/`stderr`) y captura excepciones de memoria.

### 3. Análisis de Redes Aeroportuarias (Python)
Sistema que modela una red de aeropuertos mediante grafos pesados y dirigidos para resolver problemas de logística.
* **Algoritmos implementados:**
  * Búsqueda de la ruta más rápida o barata (**Dijkstra**).
  * Minimización de escalas entre destinos (**BFS**).
  * Cálculo de los aeropuertos más críticos de la red (**Centralidad** utilizando Heaps).
  * Planificación de itinerarios sin ciclos (**Orden Topológico**).
  * Diseño de redes de bajo costo sin conexiones redundantes (**Prim / Árbol de Tendido Mínimo**).
* **Visualización:** Exportación de las rutas calculadas a formato `.kml` para su representación geoespacial.
