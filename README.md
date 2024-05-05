# OTUS_HOMEWORK

Приложение для домашнего задания курса OTUS Higload Architect.

Приложение представляет из себя заготовку социальной сети. Backend приложения написан на языке GO версии 1.22.0. Используемая БД - postgresql. 

# gRPC

Вместо REST/HTTP использовала gRPC. Реализованы следующие процедуры:
- /login (gRPC: UserLogin)
- /user/register (gRPC: UserRegister)
- /user/get/{id} (gRPC: UserGetById)

Описание процедур и типов сообщений в файле [homework-backend/internal/proto/autorization.proto](homework-backend/internal/proto/autorization.proto). Backend приложение докеризовано с помощью [Dockerfile](homework-backend/build/Dockerfile)

# Deploy

Для запуска использую задачу "Run docker-compose up" из файла [tasks.json](.vscode/tasks.json) в среде VSCode. 
После запуска задачи "Run docker-compose up" необходимо создать БД otus_homework в контейнере с образом postgresql и выполнить миграции из папки [homework-db/migrations](homework-db/migrations/). Для работы с миграциями я использовала [goose](https://github.com/pressly/goose).

# Домашнее задание №2
1. Для генерации анкет использовала файл [people.v2.csv](utils/faker/people.v2.csv) и написала утилиту [faker](utils/faker/) для вставки из файла в Postgres.
2. Добавила функционал для поиска анкет по префиксу имени и фамилии UsersGetByPrefixFirstNameAndSecondName() с сортировкой по id.
3. Графики, тесты для jmeter, запросы на создание индекса и сам отчет лежат в директории [homework2-jmeter](homework2-jmeter/) 
