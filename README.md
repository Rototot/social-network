# social-network

## Возможности

### HW 1: Заготовка для социальной сети

#### Функциональные требования:

* Авторизация по паролю

* Страница регистрации, где указывается следующая информация:

1. Имя
2. Фамилия
3. Возраст
4. Пол
5. Интересы
6. Город


* Страницы с анкетой.

## Структура проекта
```
├── backend
│   ├── bin
│   ├── cmd
│   ├── db
│   │   └── migrations
│   ├── pkg
│   │   ├── common
│   │   ├── config
│   │   ├── ping
│   │   └── users
│   ├── scripts
│   └── tests
│       ├── e2e
│       ├── factories
│       ├── fixtures
│       └── utils
└── frontend
    ├── build
    │   └── static
    ├── public
    └── src
```
#### Backend
Используется [go-kit](https://gokit.io/)
* Golang
* MySQL
* Redis

#### Frontend
* React
* Nginx

### Install

```bash
make init
make build
make migrate-up
make start
```

### Run

### Development


#### Create Migrations

```bash
make migrate-create
```

### Run migrations

```bash
make build
make migrate-up
```

#### Install Dep

#### Backend cli instrument

```
go install github.com/abice/go-enum
go install golang.org/x/tools/cmd/stringer
go install github.com/golang/mock/mockgen@v1.6.0
go install github.com/golang-migrate/migrate/v4
go install github.com/joho/godotenv/cmd/godotenv
go get github.com/mailru/easyjson && go install github.com/mailru/easyjson/...@latest
```
