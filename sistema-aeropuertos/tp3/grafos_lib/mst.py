import heapq
from grafo.grafo import Grafo

def mst_prim(grafo):
    vertice = grafo.vertice_aleatorio()
    visitados = set([vertice])
    heap = []
    mst = Grafo(False)

    mst.agregar_vertice(vertice)

    for v in grafo.obtener_adyacentes(vertice):
        mst.agregar_vertice(v)
    
    for w in grafo.obtener_adyacentes(vertice):
        peso = grafo.peso_arista(vertice, w)
        heapq.heappush(heap, (peso, vertice, w))

    while heap:
        peso, v, w = heapq.heappop(heap)
        
        if w in visitados:
            continue
        
        visitados.add(w)
        mst.agregar_vertice(w)
        mst.agregar_arista(v, w, peso)

        for u in grafo.obtener_adyacentes(w):
            if u not in visitados:
                peso = grafo.peso_arista(w, u)
                heapq.heappush(heap, (peso, w, u))
    return mst
