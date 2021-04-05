# TODOS

Тестовое задание от NapoleonIT

## Description
---

### Основное задание

Необходимо реализовать два микросервиса. Каждый из микросервисов должен представлять из себя программу-сервер. Написание исходного кода на Golang. База данных - Postgres.
Users
Микросервис Users - серверная программа отвечающая за взаимодействия с пользователями. 

Модель User должна содержать:

* id
* username
* password (не в открытом виде)

Сервис предоставляет REST API для выполнения следующих задач:
* регистрация пользователя
* авторизация пользователя и получение токена  (предоставляет токен для взаимодействия с защищенным маршрутами)
* получение пользователя по его ID

Задача функции регистрации безопасно создать нового пользователя в базе данных для последующего использования данных. При отсутствии ошибок необходимо вернуть 201-ый HTTP статус.

Путь ”Авторизация пользователя” принимает только POST запросы со следующей структурой: 
```json
{
    username: str,
    password: str
}
```

Задача функции авторизации - безопасно проверить наличие пользователя в базе данных и верность пароля. При верном username  и password необходимо вернуть JWT token. При неверном логине или пароле необходимо возвращать 401-ый HTTP статус.

Путь ”Получение пользователя” предоставляет безопасную информацию о пользователе и должен быть защищен с помощью middleware авторизации.

### Todos

Микросервис Todos - серверная программа отвечающая за взаимодействия с задачами пользователя.

Модель Todo должна содержать:
* информацию о пользователе, которому она принадлежит
* заголовок задачи
* подробное описание задачи
* время, до которого нужно исполнить задачу

Сервис представляет из себя сервер с путями, выполняющие следующие функции:
создание задач пользователя
обновление задач пользователя
удаление задач пользователя
получение Todo по id
получение всех задач пользователя ( должен возвращать массив Todos пользователя отсортированных по возрастанию времени исполнения задачи, т.е. ближайшие - первые в списке)
получение задач пользователя, которые нужно исполнить до времени заданного в Body запроса (должен возвращать массив Todos пользователя отсортированных по возрастанию времени исполнения задачи)

Все маршруты должны быть защищены с помощью middleware авторизации, т.е. только пользователь, получивший токены в сервисе Users, имеет доступ к принадлежащей ему информации.

Задача маршрута создания Todo безопасно создать новое Todo в базе для последующего использования данных. При отсутствии ошибок необходимо вернуть 201-ый HTTP статус. Каждую задачу необходимо сопоставлять с пользователем с помощью ID пользователя. Сопоставление от одного (User) ко многим (Todos).

Задача маршрута получения Todo предоставить информацию о Todos пользователя. 
Если Todos не найдены, то необходимо возвращать пустой массив.

Пользователя должен иметь доступ только до задач, созданных им.

**Рекомендации**

описать API сервисов в Swagger спецификации и приложить к проекту
создать docker-compose.yml для запуска сервисов и БД
написать тесты на маршруты сервисов

## TODO

- [ ] User service
  - [X] Код
  - [ ] Транспорт http
    - [X] Код
    - [ ] Тесты маршрутов
    - [ ] Выделить аутентификацию в отдельный middleware
  - [X] Транспорт RPC
  - [ ] Тесты
- [ ] Todos service
  - [X] Код
  - [ ] Транспорт
    - [X] Код
    - [ ] Тесты маршрутов
    - [ ] Выделить аутентификацию в отдельный middleware
  - [ ] Тесты
- [ ] Storage
  - [X] Cache
    - [X] Код
    - [X] Интерфейс
    - [X] Мок
  - [X] Database
    - [X] Код User 
    - [X] Код Task
    - [X] Интерфейс
  - [X] Интерфейс
  - [ ] Тесты
- [ ] JWT
  - [X] Код
  - [ ] Тесты
- [ ] Password hash
  - [X] Код
  - [ ] Тесты
- [X] Sender 
  - [X] SMTP
    - [X] Код 
  - [X] Интерфейс
  - [X] Мок
- [ ] Документация
- [X] docker-compose