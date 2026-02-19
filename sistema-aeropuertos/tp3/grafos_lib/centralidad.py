from grafos_lib.camino_minimo import dijkstra

INF = float("inf")

def ordenar_vertices(grafo, distancias):
    vertices_validos = [v for v in grafo.obtener_vertices() if distancias[v] != INF]
    vertices_validos.sort(key=lambda v: distancias[v], reverse=True)
    return vertices_validos

def centralidad(grafo):
    centralidad = {}
    vertices = grafo.obtener_vertices()
    
    for v in vertices:
        centralidad[v] = 0

    for v in vertices:
        distancias, padres = dijkstra(grafo, v)
        cent_aux = {}

        for w in vertices:
            cent_aux[w] = 0

        vertices_ordenados = ordenar_vertices(grafo, distancias)
        
        for w in vertices_ordenados:
            if w == v:
                continue
            cent_aux[padres[w]] += 1 + cent_aux[w]
        
        for w in vertices:
            if w != v:
                centralidad[w] += cent_aux[w]  
    return centralidad