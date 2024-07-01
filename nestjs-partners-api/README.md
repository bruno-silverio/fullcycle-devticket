<div align="center">
  
  ## Microsserviços de Parceiros de ingressos com Nest.js
  
</div>

### Descrição
Repositório da API feita com Nest.js (Reserva de Ingressos)

### Design do projeto
<div align="center">

  ![image](https://github.com/bruno-silverio/fullcycle-devticket/assets/27282770/78135602-4e7f-47b7-b834-23c18b0ecdb0)
  
</div>


### Tecnologias utilizadas

- TypeScript | JavaScript
- Nest.js
- Prisma ORM
- Rest
- Autenticação

### Ordem do desenvolvimento
- [x] Criar o projeto Nest.js
- [x] Criar o banco de dados MySQL e integrar com Prisma
- [x] Criar uma lib para ficar fácil criar várias API de parceiros
- [x] Criar API dos eventos
- [x] Criar API dos spots
- [x] Criar endpoint da API para reserva
- [x] Criar vários Apps para serem as APIs dos parceiros
- [x] Permitir que se acesse as APIs via token

## Rodar a aplicação

Dentro da pasta `nest-partners-api` execute o comando abaixo para rodar os containers `Docker`:
```
docker compose up
```

Quando os containers estiverem prontos, precisamos acessar o container do `nestjs` e executar a aplicação:

```
// entrar no container:
docker compose exec nestjs bash

// instalar as dependências:
npm install

// executar a migração do primeiro parceiro:
npm run migrate:partner1

// executar a migração do segundo parceiro:
npm run migrate:partner2

// Executar o partner1 na porta 3000
npm run start:dev

// Executar o partner2 na porta 3001
npm run start:dev partner2

```

Espere os logs verdinhos do Nest para verificar se deu certo.

* Acessar o arquivo './partner1.http' para criar / listar os eventos, reservar lugares e comprar ingressos do Parceiro 1.

* Acessar o arquivo './partner2.http' para criar / listar os eventos, reservar lugares e comprar ingressos do Parceiro 2.

### Para Windows 

Lembrar de instalar o WSL2 e Docker. Vejo o vídeo: [https://www.youtube.com/watch?v=btCf40ax0WU](https://www.youtube.com/watch?v=btCf40ax0WU) 

Siga o guia rápido de instalação: [https://github.com/codeedu/wsl2-docker-quickstart](https://github.com/codeedu/wsl2-docker-quickstart)

## Support

Nest is an MIT-licensed open source project. It can grow thanks to the sponsors and support by the amazing backers. If you'd like to join them, please [read more here](https://docs.nestjs.com/support).
