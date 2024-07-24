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

# Домашнее задание №2 Нагрузочное тестирование
1. Для генерации анкет использовала файл [people.v2.csv](utils/faker/people.v2.csv) и написала утилиту [faker](utils/faker/) для вставки из файла в Postgres.
2. Добавила функционал для поиска анкет по префиксу имени и фамилии UsersGetByPrefixFirstNameAndSecondName() с сортировкой по id.
3. Графики, тесты для jmeter, запросы на создание индекса и сам отчет лежат в директории [homework2-jmeter](homework2-jmeter/) 

# Домашнее задание №3 Репликация
1. Все файлы, использованные в ДЗ положила в директорию[homework3-replication](homework3-replication/).
2. На существующем приложения, до использования репликации сделала нагрузочные тесты. Графики нагрузки CPU,Mem,I/O docker контейнера лежат в папке [before_async_replication](homework3-replication/image/before_async_replication/).
3. Согласно инструкции из файла [guide.md](homework3-replication/guide.md) настроила асинхронную репликацию по схеме 1 мастер и два слэйва. 
4. В  [приложении](homework-backend/) перевела методы для чтения (/user/get/{id} и /user/search из спецификации) на слэйв. Конкретный слэйв указывается в [файле конфигурации](homework-backend/internal/config/config_for_docker.yaml) параметром storage_path_for_read.
5. Провела нагрузочные тесты по этим методам и удостоверилась, что в момент чтения нагрузка перешла с контейнера pg_master на контейнер pg_slave, который я указала в конфигурационном файле. Графики нагрузки CPU,Mem,I/O docker контейнера pg_master лежат в папке [after_async_replication](homework3-replication/image/after_async_replication/).
6. Согласно инструкции из файла [guide.md](homework3-replication/guide.md) настроила асинхронную репликацию по схеме 1 мастер и два слэйва. 
7. Согласно инструкции настроила кворумную синхронную репликацию и провела тесты с помощью [jmeter](homework3-replication/jmeter/Thread%20Group.jmx). 

# Домашнее задание №4 Кэширование

1. Все файлы, использованные в ДЗ положила в директорию[homework4-cache](homework4-cache/).
2. Добавила grpc service [post](homework-backend/internal/services/post/) и [proto](homework-backend/internal/proto/post.proto).
3. В этом service реализовала функционал  добавление/удаление друга методы /friendSet, /friendDelete. А также CRUD для постов пользователей (методы /postCreate, /postUpdate, /postDelete, /postGet аналогичные методам REST из спецификации). Реализован метод для отображения ленты постов друзей (метод /post/feed из спецификации). 
4. При запуске приложения выполняю процедуру заполнения кеша лентой друзей в БД Redis. Исползуемый файл [docker-compose](homework4-cache/docker-compose.yml) 
5. В ленте на каждого пользователя держу последние 1000 записей.
6. В случаях изменения, добавления, удаления постов перестраиваю ленту только для пользователей которые подписаны на автора.
7. В случае добавления друга перестраиваю ленту пользователей всех друзей добавленного друга.
8. В случае удаления друга перестраиваю всю ленту.
9. Формирование ленты произвожу через goroutine без очереди. Нагрузочный тест проблем при Эффекте Леди Гаги не выявил. Постановку задачи в очередь на часть друзей решила реализовать в следующем ДЗ по очередям.
10. Под перестройкой ленты из СУБД предполагала то, что очистка Данных БД Redis спровоцирует постепенную перестройку кеша, поэтоу дополнительный функционал для перестройки кеша из СУБД не делала.

# Домашнее задание №5 Очереди

1. Для упрощения тестирования добавила обвязку вокруг gRPC сервера в виде gRPC-gateway. Однако обвязку сделала только вокруг сервиса, который работает с сущностями Post и Friend, описано в [proto](homework-backend/internal/proto/post.proto). Настройки для обвязки в виде tcp порта указаны в файле настроек [config_for_docker.yaml](/homework-backend/internal/config/config_for_docker.yaml) параметр web_port.
2. Реализован функционал по публикации вновь созданного поста в RabbitMQ. Пост публикуется в очередь с именем равным user_id всех подписанных на автора поста пользователей. Методы публикации в очередь и вычитвания из нее описаны в файле [rabbit.go](/homework-backend/internal/storage/rabbit/rabbit.go). Настройки для подключения к RabbitMQ указаны в файле [config_for_docker.yaml](/homework-backend/internal/config/config_for_docker.yaml). Параметр rqueue_path.
3. Подключение пользователей к сервису происходит путем подключения к websocket. На каждого подключенного пользователя запускается goroutine, которая мониторит и вычитывает все появляющиеся сообщения из очереди RabbitMQ с именем равным user_id подключенного пользователя. user_id подключенного пользователя получаем из JWT токена. tcp порт для websocket указан в файле настроек [config_for_docker.yaml](/homework-backend/internal/config/config_for_docker.yaml) параметр ws_port. Под линейной масштабируемостью сервиса я предполагала возможность запуска обработки подключения к websocket отдельно для каждого пользователя.
4. Для запуска использовала [docker-compose](/homework5-queries/docker-compose.yml).

# Домашнее задание №7 Шардирование

1. Реализовала микросервис сообщений [homework-dialogs](homework-dialogs/), реализующий согласно спецификации методы отправки сообщения пользоввателю и получение диалога между двумя пользователями.
2. Отдельно стоит заметить, что я получаю id_пользователя от имени которого пишу или читаю сообщения посредством метода основного приложения из JWT токена, не уверена правильное ли это архитектурное решение.
3. Также проделала с целью изучения возможности горизонтального масштабирования путем шардирования пункты из [readme.md](homework7-sharding/readme.md) файла, используя таблицу, созданную для микросервиса сообщений. Все проделанные шаги описаны в указанном файле readme.md. Хотела как ключ шардирования использовать составной ключ из полей "from","to", но citus не разрешил. Поэтому в текущей реализации использую ключ шардирования по первичному ключу id из таблицы dialogs. В целом все работает, однако в таком случае предполагаю, что сообщения двух пользователей могут храниться на разных шардах, а этого хотелось избежать, чтобы не ходить во все шарды сразу при запросе.
4. Файл [docker-compose.yml](homework7-sharding/docker-compose.yml) для запуска всего стэка. [Dokerfile](homework-dialogs/build/Dockerfile) для создания образа с микросервисом dialogs. Я создавала с помощью задачи "Build homework-dialogs docker image" из файла [tasks.json](.vscode/tasks.json).


# Домашнее задание №6 In-memory СУБД. tarantool.

1. Провела нагрузочные тесты микросервиса [homework-dialogs](homework-dialogs/) с помощью Jmeter. [.jmx файл](homework6-tarantool/Homework6-tarantool.jmx) тестов, как и все остальные связанные с ДЗ файлы, лежат в директории [homework6-tarantool](homework6-tarantool/). Графики с результатами тестов до переезда сервиса на тарантул по критериям latencies,transactions и response time лежат в директории [graphs](homework6-tarantool/wt_tarantool/graphs/). 
2. Образ для docker с tarantool со стартовым скриптом [app.lua](homework6-tarantool/tarantool/app.lua) описан в файле [Dockerfile](homework6-tarantool/tarantool/Dockerfile).
3. Добавила в микросервис [homework-dialogs](homework-dialogs/) возможность работать с tarantool для методов Send и List. Файл настроек приложения для Docker образа - [config_dialogs.yaml](homework6-tarantool/configs/config_dialogs.yaml).
4. Файл [docker-compose.yml](homework6-tarantool/docker-compose.yml) для запуска всего стэка.
5. Графики с результатами тестов после переезда сервиса на тарантул по критериям latencies,transactions и response time лежат в директории [graphs](homework6-tarantool/tarantool/graphs/). Странно но результаты не сильно отличаются. Я ожидала что как минимум чтение будет из RAM быстрее. Возможно, что в моем ноутбуке очень быстрый nvme диск.
