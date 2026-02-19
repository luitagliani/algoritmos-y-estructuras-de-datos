from collections import deque

def grados_entrada(grafo):
    grados = {}

    for v in grafo.obtener_vertices():
        grados[v] = 0
    
    for v in grafo.obtener_vertices():
        for w in grafo.obtener_adyacentes(v):
            grados[w] += 1
            
    return grados

def orden_topologico(grafo):
    grados = grados_entrada(grafo)
    cola = deque()
    resultado = []
    
    for v in grafo.obtener_vertices():
        if grados[v] == 0:
            cola.append(v)
    
    while cola:
        vertice = cola.popleft()
        resultado.append(vertice)
        
        for w in grafo.obtener_adyacentes(vertice):
            grados[w] -= 1
            if grados[w] == 0:
                cola.append(w)
    return resultado
