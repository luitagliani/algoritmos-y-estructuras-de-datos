package pila

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

const _CAPACIDAD_INICIAL int = 50

func CrearPilaDinamica[T any]() Pila[T] {
	pila := new(pilaDinamica[T])
	pila.datos = make([]T, _CAPACIDAD_INICIAL)
	return pila
}

func (pila *pilaDinamica[T]) EstaVacia() bool {
	return pila.cantidad == 0
}

func (pila *pilaDinamica[T]) VerTope() T {
	pila.chequeoPilaVacia()
	return pila.datos[pila.cantidad-1]
}

func (pila *pilaDinamica[T]) Apilar(elemento T) {
	if pila.cantidad == cap(pila.datos) && !pila.EstaVacia() {
		pila.redimensionar(pila.cantidad * 2)
	}

	pila.datos[pila.cantidad] = elemento
	pila.cantidad++
}

func (pila *pilaDinamica[T]) Desapilar() T {
	pila.chequeoPilaVacia()
	var dato_borrado T = pila.datos[pila.cantidad-1]
	pila.cantidad--

	if (pila.cantidad*4) <= cap(pila.datos) && pila.cantidad >= _CAPACIDAD_INICIAL {
		pila.redimensionar(cap(pila.datos) / 2)
	}
	return dato_borrado
}

func (pila *pilaDinamica[T]) chequeoPilaVacia() {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
}

func (pila *pilaDinamica[T]) redimensionar(nueva_capacidad int) {
	redimension := make([]T, nueva_capacidad)
	copy(redimension, pila.datos)
	pila.datos = redimension
}
