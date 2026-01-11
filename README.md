# 🚀 Go-clean-architecture — CRUD com Go, Gin, GORM, PostgreSQL e Air

Este projeto é um **CRUD simples em Go**, criado com o objetivo de **estudar e aplicar uma arquitetura organizada e escalável**, usando ferramentas muito comuns no ecossistema Go para APIs.

O foco **não é só o CRUD**, mas **a forma como o projeto é estruturado**, separando bem responsabilidades e facilitando manutenção e crescimento do sistema.

---

## 🧰 Tecnologias utilizadas

- **Go** — linguagem principal
- **Gin** — framework HTTP rápido e minimalista
- **GORM** — ORM para acesso ao banco de dados
- **PostgreSQL** — banco de dados relacional
- **Air** — hot reload para desenvolvimento

---

## 📂 Estrutura de pastas (arquitetura)

```text
internal/
├── app/            # Inicialização da aplicação
├── config/         # Configurações (env, gorm, migrations)
├── router/         # Setup global de rotas e versões da API
├── user/           # Módulo de usuário (exemplo de domínio)
│   ├── handler.go      # Camada HTTP (Gin)
│   ├── service.go      # Regras de negócio
│   ├── repository.go   # Interface do repositório
│   ├── repository_impl.go # Implementação (GORM / DB) - Atualmente em mock
│   ├── model.go        # Entidade / Model
│   └── routes.go       # Rotas do módulo
└── tmp/            # Usado pelo Air (hot reload)

main.go             # Entry point da aplicação
```

---

### ▶️ Como rodar o projeto

Siga os passos abaixo para executar a aplicação em ambiente de desenvolvimento:

**Configurar variáveis de ambiente**  
Renomeie o arquivo `.env.example` para `.env` e ajuste as configurações conforme seu ambiente.

```bash
cp .env.example .env
```

 **Instalar o AIR** 
 ```bash
 go install github.com/air-verse/air@latest

 ```

 **Executar** 

 Rodar no terminal o comando
 ```bash
 air

