package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCrearListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
}

func TestListaInvariante(t *testing.T) {
	listaEnteros := TDALista.CrearListaEnlazada[int]()
	largoInicial := listaEnteros.Largo()
	cadenaEnteros := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := range cadenaEnteros {
		listaEnteros.InsertarPrimero(cadenaEnteros[i])
		require.EqualValues(t, cadenaEnteros[i], listaEnteros.VerPrimero())
		require.EqualValues(t, cadenaEnteros[0], listaEnteros.VerUltimo())
	}

	for i := 10; i > 0; i-- {
		require.EqualValues(t, listaEnteros.VerPrimero(), listaEnteros.BorrarPrimero())
	}

	for i := range cadenaEnteros {
		listaEnteros.InsertarUltimo(cadenaEnteros[i])
		require.EqualValues(t, cadenaEnteros[i], listaEnteros.VerUltimo())
		require.EqualValues(t, cadenaEnteros[0], listaEnteros.VerPrimero())
	}

	for i := 10; i > 0; i-- {
		require.EqualValues(t, listaEnteros.VerPrimero(), listaEnteros.BorrarPrimero())
	}

	require.True(t, listaEnteros.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaEnteros.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaEnteros.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaEnteros.VerUltimo() })
	require.EqualValues(t, largoInicial, listaEnteros.Largo())

	listaStrings := TDALista.CrearListaEnlazada[string]()
	largoInicial = listaStrings.Largo()
	cadenaStrings := [10]string{"Hola", "Como", "estas", "?", "yo", "bien", "por", "suerte", "gracias"}

	for i := range cadenaStrings {
		listaStrings.InsertarPrimero(cadenaStrings[i])
		require.EqualValues(t, cadenaStrings[i], listaStrings.VerPrimero())
		require.EqualValues(t, cadenaStrings[0], listaStrings.VerUltimo())
	}

	for i := 10; i > 0; i-- {
		require.EqualValues(t, listaStrings.VerPrimero(), listaStrings.BorrarPrimero())
	}

	for i := range cadenaStrings {
		listaStrings.InsertarUltimo(cadenaStrings[i])
		require.EqualValues(t, cadenaStrings[i], listaStrings.VerUltimo())
		require.EqualValues(t, cadenaStrings[0], listaStrings.VerPrimero())
	}

	for i := 10; i > 0; i-- {
		require.EqualValues(t, listaStrings.VerPrimero(), listaStrings.BorrarPrimero())
	}

	require.True(t, listaStrings.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaStrings.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaStrings.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaStrings.VerUltimo() })
	require.EqualValues(t, largoInicial, listaStrings.Largo())

	listaFlotantes := TDALista.CrearListaEnlazada[float64]()
	largoInicial = listaFlotantes.Largo()
	cadenaFlotantes := [10]float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9, 10.10}

	for i := range cadenaFlotantes {
		listaFlotantes.InsertarPrimero(cadenaFlotantes[i])
		require.EqualValues(t, cadenaFlotantes[i], listaFlotantes.VerPrimero())
		require.EqualValues(t, cadenaFlotantes[0], listaFlotantes.VerUltimo())
	}

	for i := 10; i > 0; i-- {
		require.EqualValues(t, listaFlotantes.VerPrimero(), listaFlotantes.BorrarPrimero())
	}

	for i := range cadenaFlotantes {
		listaFlotantes.InsertarUltimo(cadenaFlotantes[i])
		require.EqualValues(t, cadenaFlotantes[i], listaFlotantes.VerUltimo())
		require.EqualValues(t, cadenaFlotantes[0], listaFlotantes.VerPrimero())
	}

	for i := 10; i > 0; i-- {
		require.EqualValues(t, listaFlotantes.VerPrimero(), listaFlotantes.BorrarPrimero())
	}

	require.True(t, listaFlotantes.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaFlotantes.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaFlotantes.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaFlotantes.VerUltimo() })
	require.EqualValues(t, largoInicial, listaFlotantes.Largo())
}

func TestVolumen(t *testing.T) {
	listaEnteros := TDALista.CrearListaEnlazada[int]()

	for i := 0; i <= 10000; i++ {
		listaEnteros.InsertarPrimero(i)
	}

	for i := 10000; i >= 0; i-- {
		require.EqualValues(t, i, listaEnteros.BorrarPrimero())
	}

	for i := 0; i <= 10000; i++ {
		listaEnteros.InsertarUltimo(i)
		require.EqualValues(t, i, listaEnteros.VerUltimo())
	}

	for i := 0; i <= 10000; i++ {
		require.EqualValues(t, i, listaEnteros.BorrarPrimero())
	}
}

func TestListaVaciadaComoRecienCreada(t *testing.T) {
	listaEnteros := TDALista.CrearListaEnlazada[int]()

	for i := 0; i <= 100; i++ {
		listaEnteros.InsertarPrimero(i)
	}

	for i := 100; i >= 0; i-- {
		require.EqualValues(t, i, listaEnteros.BorrarPrimero())
	}

	require.True(t, listaEnteros.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaEnteros.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaEnteros.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaEnteros.BorrarPrimero() })
}

func TestLlamadadasInvalidasListaCreada(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
}
func TestInternSinCorte(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	cadenaEnteros := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := range cadenaEnteros {
		lista.InsertarUltimo(cadenaEnteros[i])
	}

	var guardarPrimerPar int
	lista.Iterar(func(n int) bool {
		if n%2 == 0 {
			guardarPrimerPar = n
			return false
		}
		return true
	})

	require.Equal(t, 2, guardarPrimerPar)
}

func TestInternoConCorte(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	cadenaEnteros := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := range cadenaEnteros {
		lista.InsertarUltimo(cadenaEnteros[i])
	}

	var contador int
	lista.Iterar(func(n int) bool {
		contador += n
		return true
	})

	require.Equal(t, 55, contador)
}

func TestInternoSumaPares(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	cadenaEnteros := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := range cadenaEnteros {
		lista.InsertarUltimo(cadenaEnteros[i])
	}

	var contador int
	lista.Iterar(func(n int) bool {
		if n%2 == 0 {
			contador += n
		}
		return true
	})

	require.Equal(t, 30, contador)
}

func TestExternoInsertarInicio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	largoInicial := lista.Largo()
	iterador := lista.Iterador()

	iterador.Insertar(1)
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())
	iterador.Borrar()
	require.EqualValues(t, lista.Largo(), largoInicial)

	lista.InsertarPrimero(1)
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())

}

func TestExternoInsertarFinal(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	cadenaEnteros := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := range cadenaEnteros {
		lista.InsertarUltimo(cadenaEnteros[i])
	}

	iterador := lista.Iterador()

	for iterador.HaySiguiente() {
		iterador.Siguiente()
	}
	largoAnterior := lista.Largo()

	iterador.Insertar(11)
	require.Equal(t, 11, lista.VerUltimo())
	iterador.Borrar()
	require.EqualValues(t, largoAnterior, lista.Largo())

	lista.InsertarUltimo(11)
	require.Equal(t, 11, lista.VerUltimo())

}

func TestExternoInsertarMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	cadenaEnteros := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := range cadenaEnteros {
		lista.InsertarUltimo(cadenaEnteros[i])
	}

	iterador := lista.Iterador()

	for iterador.HaySiguiente() && iterador.VerActual() != 5 {
		iterador.Siguiente()
	}
	largoAnterior := lista.Largo()

	iterador.Insertar(90)
	require.Equal(t, 90, iterador.VerActual())

	iterador.Siguiente()
	require.Equal(t, 5, iterador.VerActual())
	require.EqualValues(t, largoAnterior+1, lista.Largo())
}

func TestExternoEliminarInicio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	cadenaEnteros := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := range cadenaEnteros {
		lista.InsertarUltimo(cadenaEnteros[i])
	}
	largoAnterior := lista.Largo()

	iterador := lista.Iterador()

	nroBorrado := iterador.Borrar()
	require.Equal(t, 2, lista.VerPrimero())
	require.Equal(t, 1, nroBorrado)
	require.EqualValues(t, largoAnterior-1, lista.Largo())
}

func TestExternoEliminarMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	cadenaEnteros := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := range cadenaEnteros {
		lista.InsertarUltimo(cadenaEnteros[i])
	}
	largoAnterior := lista.Largo()

	iterador := lista.Iterador()

	for iterador.HaySiguiente() && iterador.VerActual() != 5 {
		iterador.Siguiente()
	}

	iterador.Borrar()

	require.NotEqual(t, 5, iterador.VerActual())
	require.Equal(t, 6, iterador.VerActual())
	require.EqualValues(t, largoAnterior-1, lista.Largo())

}

func TestIteradorListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iterador := lista.Iterador()

	require.False(t, iterador.HaySiguiente())
	require.Panics(t, func() { iterador.VerActual() })
	require.Panics(t, func() { iterador.Siguiente() })
	require.Panics(t, func() { iterador.Borrar() })
}

func TestIteradorCoincideConLista(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(1)

	iterador := lista.Iterador()
	require.Equal(t, lista.VerPrimero(), iterador.VerActual())

	for iterador.HaySiguiente() {
		iterador.Siguiente()
	}
	iterador.Insertar(4)
	require.Equal(t, lista.VerUltimo(), iterador.VerActual())
}

func TestVaciarListaConIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < 5; i++ {
		lista.InsertarUltimo(i)
	}

	iterador := lista.Iterador()
	for iterador.HaySiguiente() {
		iterador.Borrar()
	}
	require.True(t, lista.EstaVacia())
	require.EqualValues(t, 0, lista.Largo())
}

func TestIteradorEliminaImpares(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 1; i < 7; i++ {
		lista.InsertarUltimo(i)
	}

	iterador := lista.Iterador()
	for iterador.HaySiguiente() {
		if iterador.VerActual() % 2 != 0 {
			iterador.Borrar()
		} else {
			iterador.Siguiente()
		}
	}

	var resultado []int
	lista.Iterar(func(n int) bool {
		resultado = append(resultado, n)
		return true
	})

	require.Equal(t, []int{2, 4, 6}, resultado)
}
