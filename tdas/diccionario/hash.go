package diccionario

import (
	"fmt"
	"hash/fnv"
	TDALista "tdas/lista"
)

const (
	_CANTIDAD_INICIAL = 101
	_FACTOR_CARGA_MAX = 1
	_POSICION_INICIAL = -1
	_FACTOR_CARGA_MIN = 0.25
)

type parClaveValor[K comparable, V any] struct {
	clave K
	dato  V
}

type hashAbierto[K comparable, V any] struct {
	tabla    []TDALista.Lista[parClaveValor[K, V]]
	tam      int
	cantidad int
}

type iteradorHash[K comparable, V any] struct {
	diccionario     *hashAbierto[K, V]
	indiceActual    int
	iterListaActual TDALista.IteradorLista[parClaveValor[K, V]]
}

func crearTablaHash[K comparable, V any](tam int) []TDALista.Lista[parClaveValor[K, V]] {
	return make([]TDALista.Lista[parClaveValor[K, V]], tam)
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	listas := crearTablaHash[K, V](_CANTIDAD_INICIAL)
	for i := range listas {
		listas[i] = TDALista.CrearListaEnlazada[parClaveValor[K, V]]()
	}

	hashAbierto := &hashAbierto[K, V]{
		tabla:    listas,
		tam:      _CANTIDAD_INICIAL,
		cantidad: 0,
	}

	return hashAbierto
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func funcionHash(data []byte) uint64 {
	hasher := fnv.New64a()
	hasher.Write(data)
	return hasher.Sum64()
}

func (diccionario *hashAbierto[K, V]) crearParClaveValor(clave K, dato V) parClaveValor[K, V] {
	par := parClaveValor[K, V]{
		clave: clave,
		dato:  dato,
	}
	return par
}

func (diccionario *hashAbierto[K, V]) encontrarPosicion(clave K) int {
	claveBits := convertirABytes(clave)
	hash := funcionHash(claveBits)
	posicion := int(hash % uint64(diccionario.tam))
	return posicion
}

func esPrimo(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func siguientePrimo(n int) int {
	encontrado := false
	var sig_primo int
	for !encontrado {
		n++
		if esPrimo(n) {
			encontrado = true
			sig_primo = n
		}
	}
	return sig_primo
}

func anteriorPrimo(n int) int {
	encontrado := false
	var ant_primo int
	for !encontrado {
		n--
		if esPrimo(n) {
			encontrado = true
			ant_primo = n
		}
	}
	return ant_primo
}

func (diccionario *hashAbierto[K, V]) Redimensionar(nuevo_tam int) {
	tope_anterior := diccionario.tam
	tabla_anterior := diccionario.tabla
	nueva_tabla := crearTablaHash[K, V](nuevo_tam)
	diccionario.tam = nuevo_tam

	for i := range nueva_tabla {
		nueva_tabla[i] = TDALista.CrearListaEnlazada[parClaveValor[K, V]]()
	}

	for i := 0; i < tope_anterior; i++ {
		iterador := tabla_anterior[i].Iterador()
		for iterador.HaySiguiente() {
			actual := iterador.VerActual()
			posicion := diccionario.encontrarPosicion(actual.clave)
			nueva_tabla[posicion].InsertarUltimo(actual)
			iterador.Siguiente()
		}
	}

	diccionario.tabla = nueva_tabla
}

func (diccionario *hashAbierto[K, V]) Guardar(clave K, dato V) {
	if float64(diccionario.cantidad)/float64(diccionario.tam) > _FACTOR_CARGA_MAX {
		nuevo_tam := siguientePrimo(diccionario.tam * 2)
		diccionario.Redimensionar(nuevo_tam)
	}

	posicion := diccionario.encontrarPosicion(clave)
	nuevo_elemento := diccionario.crearParClaveValor(clave, dato)

	if !diccionario.tabla[posicion].EstaVacia() {
		iteradorLista := diccionario.tabla[posicion].Iterador()

		for iteradorLista.HaySiguiente() {
			actual := iteradorLista.VerActual()
			if actual.clave == clave {
				iteradorLista.Borrar()
				diccionario.tabla[posicion].InsertarUltimo(nuevo_elemento)
				return
			}
			iteradorLista.Siguiente()
		}
		diccionario.tabla[posicion].InsertarUltimo(nuevo_elemento)
		diccionario.cantidad++
	} else {
		diccionario.tabla[posicion].InsertarPrimero(nuevo_elemento)
		diccionario.cantidad++
	}
}

func (diccionario *hashAbierto[K, V]) Pertenece(clave K) bool {
	pertenece := false
	posicion := diccionario.encontrarPosicion(clave)

	if !diccionario.tabla[posicion].EstaVacia() {
		iteradorLista := diccionario.tabla[posicion].Iterador()

		for iteradorLista.HaySiguiente() {
			actual := iteradorLista.VerActual()
			if actual.clave == clave {
				pertenece = true
				break
			}
			iteradorLista.Siguiente()
		}
	}
	return pertenece
}

func (diccionario *hashAbierto[K, V]) Obtener(clave K) V {
	var dato V
	pertenece := false
	posicion := diccionario.encontrarPosicion(clave)

	iteradorLista := diccionario.tabla[posicion].Iterador()

	for iteradorLista.HaySiguiente() {
		actual := iteradorLista.VerActual()
		if actual.clave == clave {
			dato = actual.dato
			pertenece = true
		}
		iteradorLista.Siguiente()
	}

	if !pertenece {
		panic("La clave no pertenece al diccionario")
	}

	return dato
}

func (diccionario *hashAbierto[K, V]) Borrar(clave K) V {
	pertenece := false
	var dato V
	posicion := diccionario.encontrarPosicion(clave)
	iteradorLista := diccionario.tabla[posicion].Iterador()

	for iteradorLista.HaySiguiente() {
		actual := iteradorLista.VerActual()
		if actual.clave == clave {
			dato = actual.dato
			iteradorLista.Borrar()
			diccionario.cantidad--
			pertenece = true
			break
		}
		iteradorLista.Siguiente()
	}

	if float64(diccionario.cantidad)/float64(diccionario.tam) < _FACTOR_CARGA_MIN && diccionario.tam/2 > _CANTIDAD_INICIAL {
		nuevo_tam := anteriorPrimo(diccionario.tam / 2)

		diccionario.Redimensionar(nuevo_tam)

	}

	if !pertenece {
		panic("La clave no pertenece al diccionario")
	}

	return dato
}

func (diccionario *hashAbierto[K, V]) Cantidad() int {
	return diccionario.cantidad
}

func (diccionario *hashAbierto[K, V]) Iterar(visitar func(K, V) bool) {
	for i := 0; i < diccionario.tam; i++ {
		if !diccionario.tabla[i].EstaVacia() {
			iteradorLista := diccionario.tabla[i].Iterador()
			for iteradorLista.HaySiguiente() {
				actual := iteradorLista.VerActual()
				if !visitar(actual.clave, actual.dato) {
					return
				}
				iteradorLista.Siguiente()
			}
		}
	}
}

func (iterDicci *iteradorHash[K, V]) encontrarSigListaNoVacia() {
	for i := iterDicci.indiceActual + 1; i < iterDicci.diccionario.tam; i++ {
		if !iterDicci.diccionario.tabla[i].EstaVacia() {
			iterDicci.iterListaActual = iterDicci.diccionario.tabla[i].Iterador()
			iterDicci.indiceActual = i
			return
		}
	}
	iterDicci.iterListaActual = nil
}

func (diccionario *hashAbierto[K, V]) Iterador() IterDiccionario[K, V] {
	iterador := &iteradorHash[K, V]{
		diccionario:  diccionario,
		indiceActual: _POSICION_INICIAL,
	}

	iterador.encontrarSigListaNoVacia()
	return iterador
}

func (iterador *iteradorHash[K, V]) HaySiguiente() bool {
	return iterador.iterListaActual != nil
}

func (iterador *iteradorHash[K, V]) VerActual() (K, V) {
	iterador.checkHaySiguiente()
	return iterador.iterListaActual.VerActual().clave, iterador.iterListaActual.VerActual().dato
}

func (iterador *iteradorHash[K, V]) Siguiente() {
	iterador.checkHaySiguiente()
	iterador.iterListaActual.Siguiente()

	if !iterador.iterListaActual.HaySiguiente() {
		iterador.encontrarSigListaNoVacia()
	}

}

func (iterador *iteradorHash[K, V]) checkHaySiguiente() {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
}
