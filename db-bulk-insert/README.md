# Bulk insert

Execute o seguinte comando em um terminal `unix` na raiz do projeto:

```bash
export $(cat .env | xargs); sudo -E env PATH="$PATH" make all;
```

Veja o arquivo `Makefile` para mais comandos.

Execute `sudo -E env PATH="$PATH" make teardown` para encerrar o programa.

Execute `sudo -E env PATH="$PATH" make purge` para apagar o volume de dados do banco de dados.
