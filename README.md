<div align="center">
  
  ## Imersão Full Cycle - 🎟 Venda de ingressos
  
</div>

### Projeto: Sistema de Venda de Ingressos

- Gerenciamento de eventos / processamento de reservas		
  - Integração com sistemas de parceiros
- Sistemas de parceiros
- Frontend
- API Gateway

#
<img src="https://github.com/bruno-silverio/fullcycle-devticket/assets/27282770/7a5e8ef0-f6dc-418d-a248-e065abe55416" />
    
<div align="center">
  
  ## 📋 Casos de usos, arquitetura e wireframes
  
  ### Arquitetura geral do sistema
  
  <img src="https://github.com/bruno-silverio/fullcycle-devticket/assets/27282770/4b207083-c252-4c02-9637-8fd380e8b18f" />
  
  ### Operação principal - Reservar ingresso
  
  <img src="https://github.com/bruno-silverio/fullcycle-devticket/assets/27282770/804de197-07be-4d1d-82ec-e5f2161e6d9c" />

</div>

## 🚦 How to use the project (instructions)

How to run (Open 4 terminals for better viewing)

#### Docker
./fullcycle-devticket >
```bash 
docker compose up
```
#### Golang
./fullcycle-devticket >
```bash 
docker compose exec golang sh
```
/app
```bash
go run cmd/events/main.go
```
#### Partner1
./fullcycle-devticket >
```bash 
docker compose exec nestjs bash
```
~/app$
```bash 
npm install
```
```bash
npm run migrate:partner1
```
```bash
npm run migrate:partner2
```
```bash
npm run start partner1-fixture
```
```bash
npm run start partner2-fixture
```
```bash
npm run start:dev
```
#### Partner2
./fullcycle-devticket >
```bash 
docker compose exec nestjs bash
```
~/app$
```bash
npm run start:dev partner2
```
#### Nextjs
./fullcycle-devticket >
```bash 
docker compose exec nextjs bash
```
~/app$
```bash
npm run dev
```
Após iniciar as aplicações frontend e backend, basta acessar 'localhost' na porta [8000/nextjs](http://localhost:8000/nextjs) para ser direcionado à aplicação Next.js e realizar os devidos testes.

## Ordem do desenvolvimento
- [x] API de Parceiros - Nest.js
- [x] API Gateway - Kong
- [x] Sistema de gerenciamento de integremos - Golang
- [x] Frontend - Next.js
- [x] Integração de todos os sistemas

## Autores

- [@argentinaluiz](https://github.com/argentinaluiz)
- [@wesleywillians](https://github.com/wesleywillians)

##
<p align="center">
  <a href="https://skillicons.dev">
    <img src="https://skillicons.dev/icons?i=next,nest,go,docker" >
    <img src="https://github.com/bruno-silverio/fullcycle-devticket/assets/27282770/5d8c1da1-379a-4673-aea1-67c9040c4900"  width="55" height="55"/>
  </a>
</p>
