# GO_FINAL_PROJECT

## 📁 Структура проекта

```
GO_FINAL_PROJECT/
├── pkg/
│   ├── api/          # Обработчики API запросов
│   │   ├── nextdate.go
│   │   ├── task.go
│   │   ├── tasks.go
│   │   └── updateTask.go
│   ├── db/           # Работа с базой данных
│   │   ├── db.go
│   │   └── task.go
│   └── server/       # Запуск сервера
│       └── server.go
├── .env              # Конфигурация сервиса
├── compose.yaml
├── Dockerfile
├── go.mod
├── go.sum
├── main.go           # Точка входа приложения
├── README.md
└── scheduler.db      # База данных SQLite
```

**Тесты** находятся в директории `tests`:

```
tests/
├── addtask_4_test.go
├── app_1_test.go
├── db_2_test.go
├── nextdate_3_test.go
├── settings.go        # Содержит токен для тестов
├── task_6_test.go
├── task_7_test.go
└── tasks_5_test.go
```

**Фронтенд** располагается в папке `web`:

```
web/
├── css/
│   ├── style.css
│   └── theme.css
├── js/
│   ├── axios.min.js
│   └── scripts.min.js
├── favicon.ico
├── index.html
└── login.html
```

---

## 🔧 Настройки проекта

Конфигурация сделана с использование `.env` файла:

```
TODO_PORT=7540
TODO_DBFILE=scheduler.db
TODO_PASSWORD=12345
```

- **TODO_PORT** – порт для сервера  
- **TODO_DBFILE** – путь к базе данных SQLite  
- **TODO_PASSWORD** – пароль для авторизации  

---

## 💻 Технологический стек

- Язык: **Go**  
- База данных: **SQLite**  
- API формат: **JSON**  

---

## 🦡 Тестирование проекта

> Для корректной работы тестов необходимо обновить токен в `settings.go` (токен генерируется на 8 часов).
> Тесты запускаются при запущенном проекте.  

### Основные тесты:

```bash
go test -run ^TestNextDate$ ./tests      # Проверка функции NextDate
go test -run ^TestAddTask$ ./tests       # Проверка создания новой задачи
go test -run ^TestTasks$ ./tests         # Проверка получения списка задач
go test -run ^TestEditTask$ ./tests      # Проверка обновления задачи
go test -run ^TestDone$ ./tests          # Проверка отметки задачи выполненной
go test -run ^TestDelTask$ ./tests       # Проверка удаления задачи
go test ./tests                          # Запуск всех тестов
```

---

## 💣 Запуск проекта через Docker
Проект опубликован на DockerHub
**DockerHub:** [go_final_project](https://hub.docker.com/r/evgzor/go_final_project/tags)

```bash
docker run -p 7540:7540 go_final_project:v1.0.0
```

**Запуск проекта через Docker Compose:**

```bash
docker compose up
```

---

## 🔑 Ключевые моменты

- API и сервер написаны на **Go**  
- База данных **SQLite**, путь настраивается через механиз `.env` файла
- Фронтенд: HTML/CSS/JS, готовый для интеграции с API  
- Полный набор **тестов** для всех операций с задачами  
- Возможность запуска через **Docker** и **Docker Compose**
- Часть методов документирована для использования **Swagger**