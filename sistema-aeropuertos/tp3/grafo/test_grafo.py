import unittest
from grafo.grafo import Grafo

vertices = [1, 2, 3, 4, 5, 6, 7, 8, 9]

def encontrado_en_arr(elemento, arr):
    return elemento in arr

class TestGrafo(unittest.TestCase):

    def test_grafo_vacio(self):
        grafo = Grafo(False)
        self.assertEqual(grafo.cantidad(), 0)
        self.assertEqual(len(grafo.obtener_vertices()), 0)

    def test_agregar_vertices(self):
        grafo = Grafo(False)

        for i, v in enumerate(vertices):
            grafo.agregar_vertice(v)
            self.assertTrue(grafo.existe_vertice(v))
            self.assertEqual(grafo.cantidad(), i + 1)

        vs = grafo.obtener_vertices()
        for v in vs:
            self.assertTrue(encontrado_en_arr(v, vs))
            self.assertEqual(len(grafo.obtener_adyacentes(v)), 0)

    def test_agregar_vertice_duplicado(self):
        grafo = Grafo(False)
        grafo.agregar_vertice(1)
        grafo.agregar_vertice(1) 
        self.assertEqual(grafo.cantidad(), 1)

    def test_borrar_vertice(self):
        grafo = Grafo(False)
        grafo.agregar_vertice(1)
        grafo.borrar_vertice(1)
        self.assertEqual(grafo.cantidad(), 0)
        self.assertFalse(grafo.existe_vertice(1))

        with self.assertRaises(ValueError):
            grafo.borrar_vertice(1)

    def test_agregar_arista_no_pesado(self):
        grafo = Grafo(False)

        with self.assertRaises(ValueError):
            grafo.agregar_arista(1, 2, 0)

        grafo.agregar_vertice(1)
        grafo.agregar_vertice(2)
        grafo.agregar_arista(1, 2, 0)
        self.assertTrue(grafo.hay_arista(1, 2))
        self.assertTrue(grafo.hay_arista(2, 1))


        self.assertIn(2, grafo.obtener_adyacentes(1))
        self.assertIn(1, grafo.obtener_adyacentes(2))

    def test_borrar_arista(self):
        grafo = Grafo(False)
        grafo.agregar_vertice(1)
        grafo.agregar_vertice(2)

        with self.assertRaises(ValueError):
            grafo.borrar_arista(1, 2)

        grafo.agregar_arista(1, 2, 0)
        grafo.borrar_arista(1, 2)

        self.assertFalse(grafo.hay_arista(1, 2))
        self.assertNotIn(2, grafo.obtener_adyacentes(1))
        self.assertNotIn(1, grafo.obtener_adyacentes(2))

    def test_obtener_adyacentes_en_grafo_dirigido(self):
        grafo = Grafo(True)
        self.assertTrue(grafo.es_dirigido)

        grafo.agregar_vertice(vertices[0])
        for i in range(1, len(vertices)):
            grafo.agregar_vertice(vertices[i])
            grafo.agregar_arista(vertices[i-1], vertices[i], 0)

        for i in range(1, len(vertices)):
            self.assertIn(vertices[i], grafo.obtener_adyacentes(vertices[i-1]))
            self.assertNotIn(vertices[i-1], grafo.obtener_adyacentes(vertices[i]))
            self.assertTrue(grafo.hay_arista(vertices[i-1], vertices[i]))
            self.assertFalse(grafo.hay_arista(vertices[i], vertices[i-1]))

    def test_grafo_completo_no_dirigido(self):
        grafo = Grafo(False)
        for v in vertices:
            grafo.agregar_vertice(v)

        aristas = [
            (2, 1), (2, 3), (2, 9), (2, 4), (2, 7),
            (1, 5), (9, 5), (3, 6), (4, 6), (7, 8)
        ]
        for a, b in aristas:
            grafo.agregar_arista(a, b, 0)

        self.assertEqual(len(grafo.obtener_adyacentes(2)), 5)
        self.assertEqual(len(grafo.obtener_adyacentes(1)), 2)
        self.assertEqual(len(grafo.obtener_adyacentes(5)), 2)
        self.assertEqual(len(grafo.obtener_adyacentes(9)), 2)
        self.assertEqual(len(grafo.obtener_adyacentes(3)), 2)
        self.assertEqual(len(grafo.obtener_adyacentes(6)), 2)
        self.assertEqual(len(grafo.obtener_adyacentes(4)), 2)
        self.assertEqual(len(grafo.obtener_adyacentes(7)), 2)
        self.assertEqual(len(grafo.obtener_adyacentes(8)), 1)

        grafo.borrar_vertice(2)
        self.assertFalse(grafo.existe_vertice(2))
        self.assertTrue(grafo.hay_arista(1, 5))
        self.assertTrue(grafo.hay_arista(9, 5))
        self.assertTrue(grafo.hay_arista(3, 6))
        self.assertTrue(grafo.hay_arista(4, 6))
        self.assertTrue(grafo.hay_arista(7, 8))

    def test_grafo_pesado(self):
        grafo = Grafo(False)
        for v in vertices:
            grafo.agregar_vertice(v)

        for i in range(1, len(vertices)):
            grafo.agregar_arista(vertices[0], vertices[i], i * 2)
            self.assertTrue(grafo.hay_arista(vertices[0], vertices[i]))
            self.assertEqual(grafo.peso_arista(vertices[0], vertices[i]), i * 2)

if __name__ == '__main__':
    unittest.main()
