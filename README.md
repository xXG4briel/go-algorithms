# go-algorithms

Complexibilidade big O

![](https://dicionariotec.com/public_html/images/big_o_notation.png)

---
## Ordenação e pesquisa

### Busca linear
A busca linear é um algoritmo de busca que percorre um vetor elemento por elemento até encontrar o elemento desejado.

Complexidade: <span style="color: #fcff33;">**O(n)**</span>

Ou seja, o tempo de execução do algoritmo é proporcional ao tamanho do vetor.

Exemplo **(Considerando o pior caso)**:

Um vetor de n tamanho precisa ser executado n vezes para encontrar o elemento desejado.
- 512 = 512
- 256 = 256

---

### Busca binária
A busca binária é um algoritmo de busca que divide o vetor em duas partes e verifica em qual das partes o elemento desejado se encontra.

Complexidade: <span style="color: #58ff33;">**O(log n)**</span>

Ou seja, o tempo de execução do algoritmo é proporcional ao logaritmo do tamanho do vetor.

Exemplo **(Considerando o pior caso)**:

Um vetor de X tamanho precisa ser executado n vezes para encontrar o elemento desejado.
- log2(512) = 9
- log2(256) = 8
- log2(128) = 7


