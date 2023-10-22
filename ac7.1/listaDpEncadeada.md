# Lista Duplamente encadeada Encadeada

## Função Remover:

- Se o tamanho da lista for maior ou igual à posição, então faça:
  1. `p` aponta para o nó inicial.
  2. Enquanto `i` for menor que `posicao-1`, faça:
     - `p` aponta para o próximo nó.
  3. O próximo nó aponta para o próximo nó (próximo = próximo do próximo).
  4. Se o próximo for diferente de nulo, faça:
     - O anterior do próximo.anterior aponta para `p` atual.
  5. Decrementa 1.
- Senão:
  - Mostra "Não foi encontrado".

## Função Adicionar:

- Crie um nó apontando para um novo nó com o número digitado na principal.
- Ponteiro `p` criado aponta para o nó inicial.
- Enquanto `i` < `posicao-1`, faça:
  1. `p` criado aponta para próximo nó.
  2. `novo_nó.proximo` aponta para `p.próximo`.
- Se `p.próximo` for diferente de nulo, faça:
  - `p.próximo.anterior` aponta para `novo_nó`.
- `novo_nó.anterior` aponta para `p`.
- `p.próximo` aponta para `novo_nó`.
- Comprimento aumenta 1.

## Função Mostrar Lista:

- Ponteiro `nó` aponta para `início.próximo`.
- Enquanto `nó` for diferente de nulo, faça:
  - Imprimir valor do nó.
  - `nó` aponta para `nó.próximo`.

## Função Buscar:

- Ponteiro `nó` aponta para `início.próximo`.
- Enquanto `nó` for diferente de nulo, faça:
  - Se `nó.número` for igual a `numero`, retorne `nó`.
  - `nó` aponta para `nó.próximo`.
- Retorne nulo.
