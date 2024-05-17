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


# Домашнее задание №3
1. Все файлы, использованные в ДЗ плолжила в директорию[homework3-replication](homework3-replication/).
2. На существующем приложения, до использования репликации сделала нагрузочные тесты. Графики нагрузки CPU,Mem,I/O docker контейнера лежат в папке [before_async_replication](homework3-replication/image/before_async_replication/).
3. Согласно инструкции из файла [guide.md](homework3-replication/guide.md) настроила асинхронную репликацию по схеме 1 мастер и два слэйва. 
4. В  [приложении](homework-backend/) перевела методы для чтения (/user/get/{id} и /user/search из спецификации) на слэйв. Конкретный слэйв указывается в [файле конфигурации](homework-backend/internal/config/config_for_docker.yaml) параметром storage_path_for_read.
5. Провела нагрузочные тесты по этим методам и удостоверилась, что в момент чтения нагрузка перешла с контейнера pg_master на контейнер pg_slave, который я указала в конфигурационном файле. Графики нагрузки CPU,Mem,I/O docker контейнера pg_master лежат в папке [after_async_replication](homework3-replication/image/after_async_replication/).
6. Согласно инструкции из файла [guide.md](homework3-replication/guide.md) настроила асинхронную репликацию по схеме 1 мастер и два слэйва. 
7. Согласно инструкции настроила кворумную синхронную репликацию и провела тесты с помощью [jmeter](homework3-replication/jmeter/Thread%20Group.jmx). 
