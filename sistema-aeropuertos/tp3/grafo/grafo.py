import random
class Grafo:
    def __init__(self, es_dirigido):
        self.es_dirigido = es_dirigido
        self.vertices = {}
        self.cant_vertices = 0

    def agregar_vertice(self, v):
        if v not in self.vertices:
            self.vertices[v] = {}
            self.cant_vertices += 1

    def _verificar_vertice(self, v):
        if v not in self.vertices:
            raise ValueError("El vértice no pertenece al grafo")
        
    def _verificar_arista(self, v1, v2):
        if not self.hay_arista(v1, v2):
            raise ValueError("La arista no existe en el grafo")
        
    def borrar_vertice(self, v):
        self._verificar_vertice(v)

        del self.vertices[v]
        self.cant_vertices -= 1

        for adyacentes in self.vertices.values():
            if v in adyacentes:
                del adyacentes[v]

    def cantidad(self):
        return self.cant_vertices

    def obtener_adyacentes(self, v):
        self._verificar_vertice(v)
        return list(self.vertices[v].keys())

    def existe_vertice(self, v):
        return v in self.vertices

    def hay_arista(self, v1, v2):
        self._verificar_vertice(v1)
        self._verificar_vertice(v2)

        return v2 in self.vertices[v1]

    def agregar_arista(self, v1, v2, peso=1):
        self._verificar_vertice(v1)
        self._verificar_vertice(v2)

        if self.hay_arista(v1, v2):
            raise ValueError("Ya hay una arista entre los vértices referenciados")

        self.vertices[v1][v2] = peso
        if not self.es_dirigido:
            self.vertices[v2][v1] = peso

    def borrar_arista(self, v1, v2):
        self._verificar_arista(v1, v2)

        del self.vertices[v1][v2]
        if not self.es_dirigido:
            del self.vertices[v2][v1]

    def obtener_vertices(self):
        return list(self.vertices.keys())

    def peso_arista(self, v1, v2):
        self._verificar_vertice(v1)
        self._verificar_vertice(v2)

        self._verificar_arista(v1, v2)

        return self.vertices[v1][v2]
    
    def vertice_aleatorio(self):
        if self.cant_vertices == 0:
            raise ValueError("El grafo no tiene vértices")
        
        return random.choice(list(self.vertices.keys()))