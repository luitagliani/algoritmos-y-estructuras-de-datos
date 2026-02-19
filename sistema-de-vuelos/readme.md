# Trabajo práctico - Sistema de gestión de vuelos

Este proyecto es la implementación de un sistema de administración y búsqueda de vuelos desarrollado en **Go**. Permite gestionar rutas, consultar prioridades, buscar conexiones entre aeropuertos y administrar el historial de vuelos por rangos de fechas.

## Funcionalidades Principales (Interfaz `Sistema`)

El sistema implementa una arquitectura basada en interfaces para garantizar un bajo acoplamiento. Las operaciones principales incluyen:

* **Registro de vuelos:** Permite dar de alta nuevos vuelos en el sistema.
* **Búsqueda de rutas:** Encuentra el próximo vuelo disponible entre un origen y un destino.
* **Gestión por prioridades:** Obtiene los vuelos de mayor prioridad según los criterios establecidos.
* **Filtrado temporal:** Consulta información de vuelos dentro de un rango de fechas específico.
* **Administración de historial:** Permite la eliminación masiva de registros en un rango de fechas.

## Tecnologías Utilizadas
* **Lenguaje:** Go (Golang)
* **Testing:** pruebas.sh (Automatización de pruebas y comparación de flujos `stdout`/`stderr`)
