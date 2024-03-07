# OTUS_HOMEWORK

Приложение для домашнего задания курса OTUS Higload Architect.

Приложение представляет из себя заготовку социальной сети. Backend приложения написано на языке GO версии 1.22.0. Используемая БД - postgresql. 

# gRPC

Вместо REST/HTTP использовала gRPC. Реализованы следующие процедуры:
- /login (gRPC: UserLogin)
- /user/register (gRPC: UserRegister)
- /user/get/{id} (gRPC: UserGetById)

Описание процедур и типов сообщений в файле [homework-backend/internal/proto/autorization.proto](homework-backend/internal/proto/autorization.proto). Backend приложение докеризовано с помощью [Dockerfile](homework-backend/build/Dockerfile)

# Deploy

Для запуска использую задачу "Run docker-compose up" из файла [tasks.json](.vscode/tasks.json) в среде VSCode. 
После запуска задачи "Run docker-compose up" необходимо создать БД otus_homework в контейнере с образом postgresql и выполнить миграции из папки [homework-db/migrations](homework-db/migrations/). Для работы с миграциями я использовала [goose](https://github.com/pressly/goose).
