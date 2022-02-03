# social-network

## Возможности

### HW 1: Заготовка для социальной сети

### Возможности
* Авторизация
* Регистрация
* Создание персональной страницы
* Возможность подружиться
* Увидеть список друзей

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