# brand_monitor_go_lang_test

## Teste para a vaga de desenvolvedor junior em go lang

Aplicação feita em go lang v1.20. Crud com envio de emails.

- OBS -> templates para emails aleatórios tirados da stripo.
- OBS -> envio de email para mailtrap.

### Casos de envio de email

- Ao criar um usuário um email é enviado como agradecimento
- Ao adicionar um novo produto um envio para todos os usuário cadastrados na base de dados

### Envio de email em massa

- A primeira versão que fiz com o envio de email em massa estava rodando em um go routine mesmo assim a performance não estava tão boa optei por usar um job usando redis para o envio de email em massa.

### Testes

- Para rodar os testes que precisam de arquivos internos como templates html etc. Precisa trocar os caminhos pois não consegui achar uma de pegar o caminho dinamicamente.
- O .env file não era lido quando fazia testes então optava por sempre colocar a variáveis necessárias em hard code.
- Para rodar os testes é necessário configurar esta lista mencionada a cima.

### Para rodar o projeto

```shell
# docker
docker compose up -d

# verificar se os containers estão rodando na sua máquina
docker container ls || docker ps

#saida
CONTAINER ID   IMAGE              COMMAND                  CREATED        STATUS             PORTS                               NAMES
b3a301b5ed76   mysql:8.0          "docker-entrypoint.s…"   21 hours ago   Up About an hour   0.0.0.0:3306->3306/tcp, 33060/tcp   mysql_db
64a9e353a637   redis:6.2-alpine   "docker-entrypoint.s…"   21 hours ago   Up About an hour   0.0.0.0:6379->6379/tcp              redis

# instalar o air
go install github.com/cosmtrek/air@latest

# instalar deps
go mod tidy

# rodar o comando
air

```

## Demonstração

Insira um gif ou um link de alguma demonstração
![GIF](https://res.cloudinary.com/dlf01tbs6/image/upload/v1684686670/2023-05-21-12-35-56_kag4ku.gif)

## Feedback

Se você tiver algum feedback, por favor nos deixe saber por meio de icarovsilva1@gmail.com
