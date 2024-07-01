<div align="center">
  
  ## Frontend com React e Next.js
  
</div>

### Descrição

Repositório do FrontEnd da aplicação feito em Next.js

### Tecnologias utilizadas

- TypeScript | JavaScript
- Next.js e React.js
- React Server Components
- Tailwind
  
### Ordem do desenvolvimento
- [x] Criar o projeto Next.js
- [x] Criar a página HOME (Listagem de eventos)
- [x] Criar a página de layout dos lugares disponíveis de um evento
- [x] Criar a página do checkout
- [x] Criar a página de sucesso

### Rodar a aplicação

Dentro da pasta `nextjs-frontend` execute o comando abaixo para rodar o container `Docker`:
```
docker compose up
```

Quando o container estiver pronto, precisamos acessar o container do `nextjs` e executar a aplicação:

```
// entrar no container:
docker compose exec nextjs bash

// instalar as dependências:
npm install

// executar a aplicação:
npm run dev

```

### Para Windows 

Lembrar de instalar o WSL2 e Docker. Vejo o vídeo: [https://www.youtube.com/watch?v=btCf40ax0WU](https://www.youtube.com/watch?v=btCf40ax0WU) 

Siga o guia rápido de instalação: [https://github.com/codeedu/wsl2-docker-quickstart](https://github.com/codeedu/wsl2-docker-quickstart)
## Learn More

To learn more about Next.js, take a look at the following resources:

- [Next.js Documentation](https://nextjs.org/docs) - learn about Next.js features and API.
