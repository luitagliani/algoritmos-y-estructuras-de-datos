package lista

type Lista[T any] interface {

	// Devuelve true si la lista esta vacia, false en caso contrario
	EstaVacia() bool

	// Inserta en el primer nodo de la lista el elemento pasado por parametro.
	InsertarPrimero(T)

	// Inserta en el ultimo nodo de la lista el elemento pasado por parametro
	InsertarUltimo(T)

	// Borra el primer nodo de la lista. En caso de que la lista este vacia, entra en panico con un mensaje
	// "La lista esta vacia"
	BorrarPrimero() T

	// Devuelve el elemento en el primer nodo de la lista. En caso de que la lista este vacia, entra en panico con un mensaje
	// "La lista esta vacia"
	VerPrimero() T

	// Devuelve el elemento en el ultimo nodo de la lista. En caso de que la lista este vacia, entra en panico con un mensaje
	// "La lista esta vacia"
	VerUltimo() T

	// Largo devuelve la cantidad de elementos de la lista
	Largo() int

	//Aplica la función pasada por parámetro a todos los elementos de la lista, hasta que no hayan más elementos, o la función devuelva false.
	Iterar(visitar func(T) bool)

	//Crea un iterador para la lista.
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {
	//Devuelve el dato del nodo sobre el que está parado el iterador y si ya terminó de iterar entra en pánico con el mensaje
	// "El iterador termino de iterar".
	VerActual() T

	//Devuelve true si hay algo más para ver y si ya terminó de iterar entra en pánico con el mensaje "El iterador termino de iterar".
	HaySiguiente() bool

	// Itera sobre el siguiente nodo del actual.
	Siguiente()

	//Inserta un elemento sobre la posición en la que está parado el iterador.
	Insertar(T)

	//Borra el elemento en el que está parado el iterador.
	Borrar() T
}
