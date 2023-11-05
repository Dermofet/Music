# Music test backend
## Описание

Проект реализует API сервиса по хранению музыки пользователей.  

## Требования

Для успешного запуска вашего проекта на Golang необходимо следующее:

- **Golang**: Версия 1.21.0 и выше.

## Установка

1. Клонируйте репозиторий с проектом:
   ```bash
   git clone https://github.com/Dermofet/Music
   ```
   
2. Перейдите в директорию проекта:
    ```bash
    cd Music
    ```
    
3. Установите зависимости:
    ```bash
    go get
    ```
    
## Запуск

Для запуска Docker контейнеров выполните следующие команды:
```bash
cd ./dev/
docker compose up -d --build
```

Для запуска проекта вернитесь в корневую директорию проекта
```bash
cd ..
```
и выполните следующую команду:
```bash
go run ./cmd/music-backend-test/
```

Для обновления документации Swagger выполните следующую команду:
```bash
swag init -g .\cmd\music-backend-test\main.go --parseInternal
```

Для просмотра документации Swagger перейдите по ссылке:
**http://localhost:8000/docs/index.html**

## Конфигурация

Для конфигурации проекта используется файл **.env**.

#### Переменные окружения

Для настройки http сервера:
- HTTP_HOST=localhost
- HTTP_PORT=8000

Для настройки базы данных:
- DB_HOST=localhost
- DB_PORT=5433
- DB_NAME=devdb
- DB_USER=devuser
- DB_PASS=devpass 
- DB_SSLMODE=disable

## Архитектура базы данных

**СУБД**: PostgreSQL
**Таблицы базы данных:**
- users
  - id (uuid)
  - username (varchar(255))
  - password (varchar(255))
  - role (varchar(255))

- music
  - id (uuid)
  - name (varchar)
  - release_date (date)
  - size (numeric)
  - duration (interval)

- user_music
  - user_id (uuid)
  - music_id (uuid)

## Структура проекта

  - cmd/ 
    - music-backend-test/
      - config/
        - <span style="color: lightblue;">Файлы конфигураций</span>
      - main.go
  - dev/
    - .env
    - docker-compose.yml
  - Dockerfile
  - docs/
    - <span style="color: lightblue;">Файлы документации Swagger</span>
  - go.mod
  - go.sum
  - internal/
    - api/
      - http/
        - handlers/
          - <span style="color: lightblue;">Файлы обработчиков запросов</span>
        - middlewares/
          - <span style="color: lightblue;">Файлы middlewares</span>
        - presenter/
          - <span style="color: lightblue;">Файлы, для преобразования данных</span>
        - router.go
        - server.go
        - view/
          - <span style="color: lightblue;">Файлы представлений сущностей</span>
    - app/
      - app.go
      - migrate.go
      - migrations/
        - <span style="color: lightblue;">Файлы миграций</span>
    - db/
      - <span style="color: lightblue;">Файлы для работы с базой данных</span>
    - entity/
      - <span style="color: lightblue;">Файлы сущностей</span>
    - repository/
      - <span style="color: lightblue;">Файлы репозиториев</span>
    - usecase/
      - <span style="color: lightblue;">Файлы сервисов</span>
    - utils/
      - <span style="color: lightblue;">Файлы утилит</span>
  - README.md
  - run.sh

## API-документация

Здесь вы найдете описание доступных API-эндпоинтов, их методов и параметров запросов. Для каждого эндпоинта предоставьте подробное описание.

## Эндпоинты для авторизации

#### Эндпоинт 1: Аутентификация и регистрация

**Путь**: /auth/signup

**Метод**: POST

**Описание**: Этот эндпоинт предназначен для регистрации нового пользователя. Пользователь должен предоставить свои учетные данные, такие как имя пользователя и пароль. После успешной регистрации пользователь получает токен доступа.

**Тело запроса:**
```text
{
  "username": <имя_пользователя>,
  "password": <пароль_пользователя>
}
```

**Пример запроса:**
```text
POST /auth/signup
Content-Type: application/text

{
  "username": <имя_пользователя>,
  "password": <пароль_пользователя>
}
```

**Примеры ответов**
- Статус 201 Created
    ```text
    {
      "token": <токен_доступа>
    }
    ```
- Статус 401 Unauthorized
- Статус 409 Conflict
- Статус 422 UnprocessableEntity
- Статус 500 InternalServerError

#### Эндпоинт 2: Вход пользователя

**Путь**: /auth/signin

**Метод**: POST

**Описание**: Этот эндпоинт предназначен для аутентификации пользователя. Пользователь должен предоставить свой адрес электронной почты и пароль. После успешной аутентификации пользователь получает токен доступа.

**Тело запроса:**
```text
{
  "username": <имя_пользователя>,
  "password": <пароль_пользователя>
}
```

**Пример запроса:**
```text
POST /auth/signup
Content-Type: application/text

{
  "username": <имя_пользователя>,
  "password": <пароль_пользователя>
}
```

**Примеры ответов:**
- Статус 200 OK
    ```text
    {
      "token": <токен_доступа>
    }
    ```
- Статус 401 Unauthorized
- Статус 404 NotFound
- Статус 422 UnprocessableEntity
- Статус 500 InternalServerError

## Эндпоинты для работы с пользователями

#### Эндпоинт 1: Получение информации о текущем пользователе

**Путь**: /users/me

**Метод**: GET

**Описание**: Этот эндпоинт предназначен для получения данных текущего пользователя

**Пример запроса:**
```text
GET /users/me
Authorization: Bearer <токен_доступа>
```

**Примеры ответов:**
- Статус 200 OK
  ```text
  {
    "id": <id_пользователя>,
    "role": <роль_пользователя>,
    "username": <имя_пользователя>,
  }
  ```
- Статус 401 Unauthorized
- Статус 403 Forrbiden
- Статус 404 NotFound
- Статус 422 UnprocessableEntity
- Статус 500 InternalServerError

#### Эндпоинт 2: Обновление информации о текущем пользователе

**Путь**: /users/me

**Метод**: PUT

**Описание**: Этот эндпоинт предназначен для обновления информации о текущем авторизованном пользователе, такой как имя пользователя и адрес электронной почты.

**Тело запроса:**
{
  "username": <имя_пользователя>,
  "password": <пароль_пользователя>
}

**Пример запроса:**
```text
PUT /users/me
Authorization: Bearer <токен_доступа>
Content-Type: application/text

{
  "username": <имя_пользователя>,
  "password": <пароль_пользователя>
}
```

**Примеры ответов:**
- Статус 200 OK
    ```text
    {
      "id": <id_пользователя>,
      "role": <роль_пользователя>,
      "username": <имя_пользователя>,
    }
    ```
- Статус 401 Unauthorized
- Статус 403 Forrbiden
- Статус 404 NotFound
- Статус 422 UnprocessableEntity
- Статус 500 InternalServerError

#### Эндпоинт 3: Удаление текущего пользователя

**Путь**: /users/me

**Метод**: DELETE

**Описание**: Этот эндпоинт предназначен для удаления текущего авторизованного пользователя.

**Пример запроса:**
```text
DELETE /users/me
Authorization: Bearer <токен_доступа>
```

**Примеры ответов:**
- Статус 204 NoContent
- Статус 401 Unauthorized
- Статус 403 Forrbiden
- Статус 404 NotFound
- Статус 422 UnprocessableEntity
- Статус 500 InternalServerError

#### Эндпоинт 4: Получение информации о пользователе по id

**Путь**: /users/id/{id}

**Метод**: GET

**Описание**: Этот эндпоинт предназначен для получения пользователя по id.

**Параметры запроса:**
- id (uuid): id пользователя

**Пример запроса:**
```text
GET /users/id/{id}
Authorization: Bearer <токен_доступа>
```

**Примеры ответов:**
- Статус 200 OK
    ```text
    {
      "id": <id_пользователя>,
      "role": <роль_пользователя>,
      "username": <имя_пользователя>,
    }
    ```
- Статус 401 Unauthorized
- Статус 403 Forrbiden
- Статус 404 NotFound
- Статус 422 UnprocessableEntity
- Статус 500 InternalServerError

#### Эндпоинт 5: Получение информации о пользователе по username

**Путь**: /users/username/{username}

**Метод**: GET

**Описание**: Этот эндпоинт предназначен для получения данных о пользователе по username.

**Параметры запроса:**
- username (строка): username пользователя

**Пример запроса:**
```text
GET /users/username/{username}
Authorization: Bearer <токен_доступа>
```

**Примеры ответов:**
- Статус 200 OK
    ```text
    {
      "id": <id_пользователя>,
      "role": <роль_пользователя>,
      "username": <имя_пользователя>,
    }
    ```
- Статус 401 Unauthorized
- Статус 403 Forrbiden
- Статус 404 NotFound
- Статус 422 UnprocessableEntity
- Статус 500 InternalServerError

#### Эндпоинт 6: Обновление информации пользователя по id

**Путь**: /users/{id}

**Метод**: PUT

**Описание**: Этот эндпоинт предназначен для обновления пользователя, полученного по id.

**Параметры запроса:**
- id (uuid): id пользователя

**Тело запроса:**
```text
{
  "password": <пароль_пользователя>,
  "username": <имя_пользователя>
}
```

**Пример запроса:**
```text
PUT /users/{id}
Authorization: Bearer <токен_доступа>

{
  "password": <пароль_пользователя>,
  "username": <имя_пользователя>
}
```

**Примеры ответов:**
- Статус 200 OK
    ```text
    {
      "id": <id_пользователя>,
      "role": <роль_пользователя>,
      "username": <имя_пользователя>,
    }
    ```
- Статус 401 Unauthorized
- Статус 403 Forrbiden
- Статус 404 NotFound
- Статус 422 UnprocessableEntity
- Статус 500 InternalServerError

#### Эндпоинт 7: Удаление пользователя по id

**Путь**: /users/{id}

**Метод**: DELETE

**Описание**: Этот эндпоинт предназначен для удаления пользователя по id.

**Параметры запроса:**
- id (uuid): id пользователя

**Пример запроса:**
```text
DELETE /users/{id}
Authorization: Bearer <токен_доступа>
```

**Примеры ответов:**
- Статус 204 NoContemt
- Статус 401 Unauthorized
- Статус 403 Forrbiden
- Статус 422 UnprocessableEntity
- Статус 500 InternalServerError

#### Эндпоинт 8: Добавление трека пользователем в понравившиеся

**Путь**: /users/add-track/{id}

**Метод**: POST

**Описание**: Этот эндпоинт предназначен для добавления трека пользователем в понравившиеся

**Параметры запроса:**
- id (uuid): id трека

**Пример запроса:**
```text
POST /users/add-track/{id}
Authorization: Bearer <токен_доступа>
```

**Примеры ответов:**
- Статус 200 OK
- Статус 401 Unauthorized
- Статус 403 Forrbiden
- Статус 422 UnprocessableEntity
- Статус 500 InternalServerError

#### Эндпоинт 9: Удаление трека пользователем из понравившихся

**Путь**: /users/remove-track/{id}

**Метод**: DELETE

**Описание**: Этот эндпоинт предназначен для удаления трека пользователем из понравившихся

**Параметры запроса:**
- id (uuid): id трека

**Пример запроса:**
```text
DELETE /users/remove-track/{id}
Authorization: Bearer <токен_доступа>
```

**Примеры ответов:**
- Статус 204 OK
- Статус 401 Unauthorized
- Статус 403 Forrbiden
- Статус 422 UnprocessableEntity
- Статус 500 InternalServerError

#### Эндпоинт 10: Получение списка треков, которые понравились пользователю

**Путь**: /users/get-liked-tracks

**Метод**: GET

**Описание**: Этот эндпоинт предназначен для получения списка треков, которые понравились пользователю

**Пример запроса:**
```text
GET /users/get-liked-tracks
Authorization: Bearer <токен_доступа>
```

**Примеры ответов:**
- Статус 200 OK
  ```text
  [
    {
      "duration": <продолжительность_трека>,
      "id": <id_трека>,
      "name": <название_трека>,
      "size": <размер_файла>
    }
  ]
  ```
- Статус 401 Unauthorized
- Статус 403 Forrbiden
- Статус 422 UnprocessableEntity
- Статус 500 InternalServerError

## Эндпоинты для работы с треками

#### Эндпоинт 1: Получение списка треков, которые понравились пользователю

**Путь**: /music/catalog

**Метод**: GET

**Описание**: Этот эндпоинт предназначен для получения списка всех треков.

**Пример запроса:**
```text
GET /music/catalog
Authorization: Bearer <токен_доступа>
```

**Примеры ответов:**
- Статус 200 OK
  ```text
  [
    {
      "duration": <продолжительность_трека>,
      "id": <id_трека>,
      "name": <название_трека>,
      "size": <размер_файла>
    }
  ]
  ```
- Статус 401 Unauthorized
- Статус 403 Forrbiden
- Статус 422 UnprocessableEntity
- Статус 500 InternalServerError

#### Эндпоинт 2: Получение списка треков, отсортированных по популярности

**Путь**: /music/popular

**Метод**: GET

**Описание**: Этот эндпоинт предназначен для получения списка треков, отсортированных по популярности. Популярность определяется по количеству добавлений трека в понравившиеся.

**Пример запроса:**
```text
GET /music/popular
Authorization: Bearer <токен_доступа>
```

**Примеры ответов:**
- Статус 200 OK
  ```text
  [
    {
      "duration": <продолжительность_трека>,
      "id": <id_трека>,
      "name": <название_трека>,
      "size": <размер_файла>
    }
  ]
  ```
- Статус 401 Unauthorized
- Статус 403 Forrbiden
- Статус 422 UnprocessableEntity
- Статус 500 InternalServerError

#### Эндпоинт 3: Получение списка треков, отсортированных по дате релиза

**Путь**: /music/release

**Метод**: GET

**Описание**: Этот эндпоинт предназначен для получения списка треков, отсортированных по дате релиза.

**Пример запроса:**
```text
GET /music/release
Authorization: Bearer <токен_доступа>
```

**Примеры ответов:**
- Статус 200 OK
  ```text
  [
    {
      "duration": <продолжительность_трека>,
      "id": <id_трека>,
      "name": <название_трека>,
      "size": <размер_файла>
    }
  ]
  ```
- Статус 401 Unauthorized
- Статус 403 Forrbiden
- Статус 422 UnprocessableEntity
- Статус 500 InternalServerError

#### Эндпоинт 4: Создание трека

**Путь**: /music/new

**Метод**: POST

**Описание**: Этот эндпоинт предназначен для получения списка треков, отсортированных по дате релиза.

**Форма в запросе:**
- name (string): название трека
- release (time.Time): дата релиза
- file (file): файл трека

**Пример запроса:**
```text
POST /music/new
Authorization: Bearer <токен_доступа>

{
  "name": <название_трека>,
  "release": <дата_релиза>,
  "file": <файл_трека>
}
```

**Примеры ответов:**
- Статус 201 Created
- Статус 401 Unauthorized
- Статус 403 Forrbiden
- Статус 422 UnprocessableEntity
- Статус 500 InternalServerError

#### Эндпоинт 5: Обновление трека

**Путь**: /music/{id}

**Метод**: PUT

**Описание**: Этот эндпоинт предназначен для обновления трека по id.

**Форма в запросе:**
- id (uuid): id трека
- name (string): название трека
- release (time.Time): дата релиза
- file (file): файл трека

**Пример запроса:**
```text
PUT /music/{id}
Authorization: Bearer <токен_доступа>

{
  "id": <id_трека>,
  "name": <название_трека>,
  "release": <дата_релиза>,
  "file": <файл_трека>
}
```

**Примеры ответов:**
- Статус 200 Created
- Статус 401 Unauthorized
- Статус 403 Forrbiden
- Статус 422 UnprocessableEntity
- Статус 500 InternalServerError

#### Эндпоинт 6: Удаление трека

**Путь**: /music/{id}

**Метод**: DELETE

**Описание**: Этот эндпоинт предназначен для удаления трека по id.

**Пример запроса:**
```text
DELETE /music/{id}
Authorization: Bearer <токен_доступа>
```

**Примеры ответов:**
- Статус 204 NoContent
- Статус 401 Unauthorized
- Статус 403 Forrbiden
- Статус 422 UnprocessableEntity
- Статус 500 InternalServerError

