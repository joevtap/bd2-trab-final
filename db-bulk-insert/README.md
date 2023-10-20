# Bulk insert

Execute o seguinte comando em um terminal `unix` na raiz do projeto:

```bash
export $(cat .env | xargs) && env PATH="$PATH" make all
```

Veja o arquivo `Makefile` para mais comandos.

Execute `env PATH="$PATH" make teardown` para encerrar o programa.

Execute `env PATH="$PATH" make purge` para apagar o volume de dados do banco de dados.
