package pila_test

import (
	"github.com/stretchr/testify/require"
	TDAPila "tdas/pila"
	"testing"
)

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())

	require.Panics(t, func() {
		pila.VerTope()
	}, "VerTope() debe generar panic cuando la pila está vacía")

	require.Panics(t, func() {
		pila.Desapilar()
	}, "Desapilar() debe generar panic cuando la pila está vacía")
}

func TestOrdenInt(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()

	pila.Apilar(1)
	require.Equal(t, 1, pila.VerTope())
	pila.Apilar(2)
	require.Equal(t, 2, pila.VerTope())
	pila.Apilar(3)
	require.Equal(t, 3, pila.VerTope())

	require.Equal(t, 3, pila.Desapilar())
	require.Equal(t, 2, pila.VerTope())
	require.Equal(t, 2, pila.Desapilar())
	require.Equal(t, 1, pila.VerTope())
	require.Equal(t, 1, pila.Desapilar())
	require.True(t, pila.EstaVacia())
}

func TestOrdenString(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[string]()

	pila.Apilar("Hola")
	require.Equal(t, "Hola", pila.VerTope())
	pila.Apilar("Hello")
	require.Equal(t, "Hello", pila.VerTope())
	pila.Apilar("Aloha")
	require.Equal(t, "Aloha", pila.VerTope())

	require.Equal(t, "Aloha", pila.Desapilar())
	require.Equal(t, "Hello", pila.VerTope())
	require.Equal(t, "Hello", pila.Desapilar())
	require.Equal(t, "Hola", pila.VerTope())
	require.Equal(t, "Hola", pila.Desapilar())
	require.True(t, pila.EstaVacia())
}

func TestDeVolumen(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	total := 10000

	for i := 0; i < total; i++ {
		pila.Apilar(i)
		require.Equal(t, i, pila.VerTope())
	}

	for i := total; i > 0; i-- {
		require.Equal(t, i-1, pila.Desapilar())
		if !pila.EstaVacia() {
			require.Equal(t, i-2, pila.VerTope())
		}
	}

	require.Panics(t, func() {
		pila.VerTope()
	}, "VerTope() debe generar panic cuando la pila está vacía")

	require.Panics(t, func() {
		pila.Desapilar()
	}, "Desapilar() debe generar panic cuando la pila está vacía")

}
