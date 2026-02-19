
def obtener_aristas(grafo):
    aristas = []
    visitados = set()
    for v in grafo.obtener_vertices():
        for w in grafo.obtener_adyacentes(v):
            if w not in visitados:
                peso = grafo.peso_arista(v, w)
                aristas.append((v, w, peso))
        visitados.add(v)
    return aristas  