package diccionario

import (
    TDAPila "tdas/pila"
)

type funcCmp[K any] func(K, K) int

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	dato      V
}

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      funcCmp[K]
}
type iteradorAbb[K comparable, V any] struct {
    pila TDAPila.Pila[*nodoAbb[K, V]]
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) *abb[K, V] {
    return &abb[K, V]{
        cmp: funcion_cmp,
    }
}
//func (abb *abb[K, V]) Guardar(clave K, dato V){ 

// }


// func (abb *abb[K, V]) Pertenece(clave K) bool{ 

// }

// func (abb *abb[K, V]) Obtener(clave K) V{ 

// }

// func (abb *abb[K, V]) Borrar(clave K) V{ 

// }

func (abb *abb[K, V]) Cantidad() int{
    return abb.cantidad
}

// func (abb *abb[K, V]) Iterar(func(clave K, dato V) bool){

// }

func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] {
	pila_hizq := TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()
	actual := abb.raiz

	for actual != nil {
		pila_hizq.Apilar(actual)
		actual = actual.izquierdo
	}

	return &iteradorAbb[K, V]{
        pila: pila_hizq,
    }
}


func (it *iteradorAbb[K, V]) HaySiguiente() bool{
    return !it.pila.EstaVacia()
}


func (it *iteradorAbb[K, V]) VerActual() (K, V){
    it.checkHaySiguiente()
    nodo := it.pila.VerTope()
    return nodo.clave, nodo.dato
}

func (it *iteradorAbb[K, V]) Siguiente(){
    it.checkHaySiguiente()
    nodo := it.pila.Desapilar()

    if nodo.derecho != nil {
        actual := nodo.derecho
        for actual != nil{
            it.pila.Apilar(actual)
            actual = nodo.izquierdo
        }
    }
}


// func (it *iteradorAbb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool){

// }

// func (it *iteradorAbb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V]{

// }

func (it *iteradorAbb[K, V]) checkHaySiguiente() {
	if !it.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
}