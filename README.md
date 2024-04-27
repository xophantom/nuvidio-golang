# Nuvidio Challenge - Go

Implementação em Go que encontra o tamanho da maior substring sem repetição de caracteres em uma string fornecida. Utilizei a técnica de 'sliding window'.

### lengthOfLongestSubstring
A função lengthOfLongestSubstring recebe uma string como entrada e retorna o tamanho da maior substring que não contém caracteres repetidos. A função utiliza map para rastrear o índice mais recente de cada caractere e ajusta o início de uma "janela" deslizante para ir excluindo os caracteres repetidos.

# Executar o Projeto

## Pré-requisitos
Certifique-se de ter o Go instalado.

## Instruções de Execução
Clone ou baixe este repositório.
Navegue até o diretório do projeto.
Abra um terminal neste diretório.
Execute o comando abaixo para rodar o script:

```
go run main.go
```
