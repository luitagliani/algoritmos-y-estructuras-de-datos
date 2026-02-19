import heapq
from collections import deque

INF = float("inf")

def dijkstra(grafo, origen):
    distancias = {}
    padres = {}
    
    for v in grafo.obtener_vertices():
        distancias[v] = INF
        padres[v] = None

    distancias[origen] = 0
    padres[origen] = None

    heap = [(0, origen)]

    while heap:
        dist_actual, vertice_actual = heapq.heappop(heap)
        
        if dist_actual != distancias[vertice_actual]:
            continue

        for w in grafo.obtener_adyacentes(vertice_actual):
            peso = grafo.peso_arista(vertice_actual, w)
            nueva_dist = dist_actual + peso
            
            if nueva_dist < distancias[w]:
                distancias[w] = nueva_dist
                padres[w] = vertice_actual
                heapq.heappush(heap, (nueva_dist, w))
    return distancias, padres

def bfs(grafo, origen):
    padres = {}
    distancias = {}
    cola = deque()

    padres[origen] = None
    distancias[origen] = 0
    cola.append(origen)

    while cola:
        v = cola.popleft()
        for w in grafo.obtener_adyacentes(v):
            if w not in padres:
                padres[w] = v
                distancias[w] = distancias[v] + 1
                cola.append(w)
    return distancias, padres
