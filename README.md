<div align="center" >
  <img src="assets/logo.svg" width="100px" alt="Astral logo" />
</div>

# Astral

a minimal digital content marketplace with powered by
1. Go
2. pg
3. sqlc
4. datastar
5. goose

```
├── assets
│   └── logo.svg
├── cmd
│   └── web
│       ├── handler.go
│       ├── main.go
│       └── routes.go
├── data
│   ├── migrations
│   │   └── 20260917123431_astral.sql
│   └── queries
│       └── users.sql
├── internal
│   └── db
│       ├── db.go
│       ├── models.go
│       └── users.sql.go
├── views
│   ├── static
│   │   ├── assets
│   │   │   └── logo.svg
│   │   ├── datastar.js
│   │   └── styles.css
│   ├── templates
│   │   ├── styles
│   │   │   └── index.css
│   │   ├── head.templ
│   │   ├── head_templ.go
│   │   ├── nav.templ
│   │   └── nav_templ.go
│   ├── index.templ
│   ├── index_templ.go
│   ├── shop.templ
│   ├── shop_templ.go
│   └── views.go
├── .air.toml
├── .env.example
├── .gitignore
├── go.mod
├── go.sum
├── README.md
└── sqlc.yaml
```
