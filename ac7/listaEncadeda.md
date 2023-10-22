# Lista Encadeada

## Função Add

A função `Add` insere um novo nó na lista encadeada:

1. Crie um novo nó.
2. Verifique se o nó inicial está vazio.
   - Se o nó inicial estiver vazio, aponte para o novo nó criado.
   - Se o nó final não estiver vazio, aponte o próximo nó para o novo nó criado.
3. Atribua o nó final ao nó criado.

## Função MostraLista

A função `MostraLista` exibe os elementos da lista encadeada:

1. Crie um nó com o valor do nó inicial.
2. Enquanto o nó não estiver vazio, faça:
   - Imprima o valor do nó.
   - Atualize o nó para o próximo nó.

## Função Remove

A função `Remove` remove um nó da lista encadeada com base em um valor fornecido:

1. Se o valor do nó inicial for igual ao valor buscado, faça:
   - Aponte o valor inicial para o próximo.
   - Pare.
2. Enquanto o próximo nó não for nulo, faça:
   - Se o valor do nó é igual ao valor procurado, faça:
     - Aponte o valor do nó anterior para o próximo.
     - Pare.
   - Atualize o valor anterior com o nó.
   - Atualize o nó com o próximo valor do nó.
3. Se o último nó for igual ao valor, faça:
   - Passe o valor final para o nó anterior.
