package lista

type nodoLista[T any] struct {
	siguiente *nodoLista[T]
	dato      T
}

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}
type iteradorLista[T any] struct {
	actual   *nodoLista[T]
	anterior *nodoLista[T]
	lista    *listaEnlazada[T]
}

func CrearListaEnlazada[T any]() Lista[T] {
	lista := new(listaEnlazada[T])
	return lista
}

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.primero == nil
}

func crearNuevoNodo[T any](elemento T, nodoActual *nodoLista[T]) *nodoLista[T] {
	nuevoNodo := new(nodoLista[T])
	nuevoNodo.dato, nuevoNodo.siguiente = elemento, nodoActual

	return nuevoNodo
}

func (lista *listaEnlazada[T]) InsertarPrimero(elemento T) {
	nuevoNodo := crearNuevoNodo(elemento, lista.primero)

	if lista.EstaVacia() {
		lista.ultimo = nuevoNodo
		nuevoNodo.siguiente = nil
	}

	lista.primero = nuevoNodo
	lista.largo++
}

func (lista *listaEnlazada[T]) InsertarUltimo(elemento T) {
	nuevoNodo := crearNuevoNodo(elemento, nil)

	if lista.EstaVacia() {
		lista.primero = nuevoNodo
	} else {
		lista.ultimo.siguiente = nuevoNodo
	}

	lista.ultimo = nuevoNodo
	lista.largo++
}

func (lista *listaEnlazada[T]) verDato(nodo *nodoLista[T]) T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}

	return nodo.dato
}

func (lista *listaEnlazada[T]) VerPrimero() T {
	return lista.verDato(lista.primero)
}

func (lista *listaEnlazada[T]) VerUltimo() T {
	return lista.verDato(lista.ultimo)
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	datoEliminado := lista.VerPrimero()

	if lista.primero.siguiente == nil {
		lista.ultimo = nil
	}

	lista.primero = lista.primero.siguiente
	lista.largo--

	return datoEliminado
}

func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	actual := lista.primero
	for actual != nil && visitar(actual.dato) {
		actual = actual.siguiente
	}
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	iterador := &iteradorLista[T]{
		lista:    lista,
		actual:   lista.primero,
		anterior: nil,
	}
	return iterador
}

func (iterador *iteradorLista[T]) VerActual() T {
	iterador.checkHaySiguiente()
	return iterador.actual.dato
}

func (iterador *iteradorLista[T]) HaySiguiente() bool {
	return iterador.actual != nil
}

func (iterador *iteradorLista[T]) Siguiente() {
	iterador.checkHaySiguiente()
	iterador.anterior = iterador.actual
	iterador.actual = iterador.actual.siguiente
}

func (iterador *iteradorLista[T]) Insertar(elemento T) {
	nuevoNodo := crearNuevoNodo(elemento, iterador.actual)

	if iterador.anterior == nil {
		iterador.lista.primero = nuevoNodo
	} else {
		iterador.anterior.siguiente = nuevoNodo
	}

	if iterador.actual == nil {
		iterador.lista.ultimo = nuevoNodo
	}

	iterador.actual = nuevoNodo
	iterador.lista.largo++
}

func (iterador *iteradorLista[T]) Borrar() T {
	iterador.checkHaySiguiente()
	datoBorrado := iterador.actual.dato
	iterador.actual = iterador.actual.siguiente

	if iterador.anterior == nil {
		iterador.lista.primero = iterador.actual
	} else {
		iterador.anterior.siguiente = iterador.actual
	}

	if iterador.actual == nil {
		iterador.lista.ultimo = iterador.anterior
	}
	iterador.lista.largo--

	return datoBorrado
}

func (iterador *iteradorLista[T]) checkHaySiguiente() {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
}
