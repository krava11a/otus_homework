# Шардирование

1) Скачаем docker-compose.yml, которым будем пользоваться в дальнейшем.

curl https://raw.githubusercontent.com/citusdata/docker/master/docker-compose_citus.yml > docker-compose.yml

2) POSTGRES_PASSWORD=pass 
```shell
    docker-compose -f docker-compose_citus.yml -p citus up --scale worker=2 -d
```

3) Подключимся к координатору:
```shell
docker exec -it citus_master psql -U postgres
```
4) Таблицу руками не создавала, а запустила микросервис [homework-dialogs](../homework-dialogs/), указав в файле конфигурации [config.yml](../homework-dialogs/internal/config/config.yaml) в параметре storage_path параметры для подключения к citus. Микросервис сам создает табличку.


5) Создадим из нее распределенную (шардированную) таблицу:
```postgresql
SELECT create_distributed_table('dialogs', 'from');
```
Mapping:

id -> hash(id) % 32 -> shard -> worker

```postgresql
SELECT shard_count FROM citus_tables WHERE table_name::text = 'dialogs';
```
6) Наполняла данными запросами из postman

7) Посмотрим план запроса. Видим, что select теперь распределенный и пойдет на все шарды:
```postgresql
explain select * from dialogs limit 10;
```
```text
QUERY PLAN
-----------------------------------------------------------------------------------------------
Limit  (cost=0.00..0.00 rows=10 width=552)
   ->  Custom Scan (Citus Adaptive)  (cost=0.00..0.00 rows=100000 width=552)
         Task Count: 32
         Tasks Shown: One of 32
         ->  Task
               Node: host=citus_worker_1 port=5432 dbname=postgres
               ->  Limit  (cost=0.00..0.81 rows=10 width=552)
                     ->  Seq Scan on dialogs_102008 dialogs  (cost=0.00..11.40 rows=140 width=552)
(8 rows)
```

8) Посмотрим план запроса по конкретному id. Видим, что такой select отправится только на один из шардов:
```
postgres=# explain select * from dialogs where "from"='e6caf552-37bb-4482-8634-2cf17506f25c' limit 10;
                                        QUERY PLAN 
------------------------------------------------------------------------------------------------------------
Custom Scan (Citus Adaptive)  (cost=0.00..0.00 rows=0 width=0)
   Task Count: 1
   Tasks Shown: All
   ->  Task
         Node: host=citus_worker_1 port=5432 dbname=postgres
         ->  Limit  (cost=0.00..11.75 rows=1 width=552)
               ->  Seq Scan on dialogs_102022 dialogs  (cost=0.00..11.75 rows=1 width=552)
                     Filter: ("from" = 'e6caf552-37bb-4482-8634-2cf17506f25c'::uuid)
(8 rows)

9) Добавим еще парочку шардов:

docker-compose -f docker-compose_citus.yml -p citus up --scale worker=5 -d

10) Посмотрим, видит ли координатор новые шарды:

SELECT master_get_active_worker_nodes();

11) Проверим, на каких узлах лежат сейчас данные:

SELECT nodename, count(*)
FROM citus_shards GROUP BY nodename;

12) Видим, что данные не переехали на новые узлы, надо запустить перебалансировку.

13) Для начала установим wal_level = logical чтобы узлы могли переносить данные:

alter system set wal_level = logical;
SELECT run_command_on_workers('alter system set wal_level = logical');

15) Перезапускаем все узлы в кластере, чтобы применить изменения wal_level.

docker-compose -f docker-compose_citus.yml restart

16) Проверим, что wal_level изменился:

docker exec -it citus-worker-1 psql -U postgres

show wal_level;

wal_level изменился однако следующий пункт выдал ошибку из-за того что у меня нет PRIMARY_KEY в таблице dialogs:

```
ERROR:  cannot use logical replication to transfer shards of the relation dialogs since it doesn't have a REPLICA IDENTITY or PRIMARY KEY
DETAIL:  UPDATE and DELETE commands on the shard will error out during logical replication unless there is a REPLICA IDENTITY or PRIMARY KEY.
HINT:  If you wish to continue without a replica identity set the shard_transfer_mode to 'force_logical' or 'block_writes'.
```


17) Запустим ребалансировку:

docker exec -it citus_master psql -U postgres

SELECT citus_rebalance_start();

18) Следим за статусом ребалансировки, пока не увидим там соообщение "task_state_counts": {"done": 18}

SELECT * FROM citus_rebalance_status();

19) Проверяем, что данные равномерно распределились по шардам:

SELECT nodename, count(*)
FROM citus_shards GROUP BY nodename;

20) Создадим референсную таблицу:

CREATE TABLE test_reference (
id bigint NOT NULL PRIMARY KEY,
data text NOT NULL
);

Создала референсную табличку так:

CREATE TABLE dialogs_reference (id SERIAL PRIMARY KEY,"from" UUID NOT NULL, "to" UUID NOT NULL, "text" VARCHAR(255) NOT NULL, "timestamp"  DATE NOT NULL);

select create_reference_table('dialogs_reference');

insert into dialogs_reference values(1,'e6caf552-37bb-4482-8634-2cf17506f25c','e4928f51-f1a3-449a-83e0-37fe0e143ed6',' a!!','2024-07-11');
insert into dialogs_reference values(2,'e6caf552-37bb-4482-8634-2cf17506f25c','e4928f51-f1a3-449a-83e0-37fe0e143ed6',' b!!','2024-07-12');
insert into dialogs_reference values(3,'e6caf552-37bb-4482-8634-2cf17506f25c','e4928f51-f1a3-449a-83e0-37fe0e143ed6',' c!!','2024-07-13');
insert into dialogs_reference values(4,'e6caf552-37bb-4482-8634-2cf17506f25c','e4928f51-f1a3-449a-83e0-37fe0e143ed6',' d!!','2024-07-14');
insert into dialogs_reference values(5,'e6caf552-37bb-4482-8634-2cf17506f25c','e4928f51-f1a3-449a-83e0-37fe0e143ed6',' e!!','2024-07-15');
<!-- insert into dialogs_reference values('e6caf552-37bb-4482-8634-2cf17506f25c','e4928f51-f1a3-449a-83e0-37fe0e143ed6',' f!!','2024-07-16');
insert into dialogs_reference values('e6caf552-37bb-4482-8634-2cf17506f25c','e4928f51-f1a3-449a-83e0-37fe0e143ed6',' g!!','2024-07-17'); -->




select * from dialogs join dialogs_reference on dialogs.id = dialogs_reference.id limit 10;

explain select * from dialogs join dialogs_reference on dialogs.id = dialogs_reference.id where dialogs.id = 1 limit 10;

21) Создадим связанные таблицы:
    CREATE TABLE dialogs_colocate (id SERIAL PRIMARY KEY,"from" UUID NOT NULL, "to" UUID NOT NULL, "text" VARCHAR(255) NOT NULL, "timestamp"  DATE NOT NULL);

SELECT create_distributed_table('dialogs_colocate', 'id', colocate_with => 'dialogs');

insert into dialogs_colocate(id, "from","to","text","timestamp")
select
i,
'e6caf552-37bb-4482-8634-2cf17506f25c','e4928f51-f1a3-449a-83e0-37fe0e143ed6',' f!!','2024-07-16'
from generate_series(1, 1000000) as i;

select * from dialogs join dialogs_colocate on dialogs.id = dialogs_colocate.id limit 10;

Посмотрим explain запроса:

explain select * from dialogs join dialogs_colocate on dialogs.id = dialogs_colocate.id limit 10;
                                                                         QUERY PLAN                                                                         
------------------------------------------------------------------------------------------------------------------------------------------------------------
 Limit  (cost=0.00..0.00 rows=10 width=1112)
   ->  Custom Scan (Citus Adaptive)  (cost=0.00..0.00 rows=100000 width=1112)
         Task Count: 32
         Tasks Shown: One of 32
         ->  Task
               Node: host=citus_worker_4 port=5432 dbname=postgres
               ->  Limit  (cost=0.29..60.49 rows=10 width=601)
                     ->  Nested Loop  (cost=0.29..782.95 rows=130 width=601)
                           ->  Seq Scan on dialogs_102041 dialogs  (cost=0.00..11.30 rows=130 width=556)
                           ->  Index Scan using dialogs_colocate_pkey_102074 on dialogs_colocate_102074 dialogs_colocate  (cost=0.29..5.94 rows=1 width=45)
                                 Index Cond: (id = dialogs.id)
(11 rows)

select * from dialogs join dialogs_colocate on dialogs.id = dialogs_colocate.id where dialogs.id = 1 limit 10;

Смотрим explain и смотрим на какое количество шардов полетят запросы:

explain select * from dialogs join dialogs_colocate on dialogs.id = dialogs_colocate.id where dialogs.id = 1 limit 10;
                                                                      QUERY PLAN                                                                      
------------------------------------------------------------------------------------------------------------------------------------------------------
 Custom Scan (Citus Adaptive)  (cost=0.00..0.00 rows=0 width=0)
   Task Count: 1
   Tasks Shown: All=   ->  Task
         Node: host=citus_worker_4 port=5432 dbname=postgres
         ->  Limit  (cost=0.43..16.48 rows=1 width=601)
               ->  Nested Loop  (cost=0.43..16.48 rows=1 width=601)
                     ->  Index Scan using dialogs_pkey_102042 on dialogs_102042 dialogs  (cost=0.14..8.16 rows=1 width=556)
                           Index Cond: (id = 1)
                     ->  Index Scan using dialogs_colocate_pkey_102075 on dialogs_colocate_102075 dialogs_colocate  (cost=0.29..8.30 rows=1 width=45)
                           Index Cond: (id = 1)
(11 rows)

22) Создадим локальную таблицу и попробуем сджойнить:

CREATE TABLE dialogs_local (id SERIAL PRIMARY KEY,"from" UUID NOT NULL, "to" UUID NOT NULL, "text" VARCHAR(255) NOT NULL, "timestamp"  DATE NOT NULL);

insert into dialogs_local values(1,'e6caf552-37bb-4482-8634-2cf17506f25c','e4928f51-f1a3-449a-83e0-37fe0e143ed6',' a!!','2024-07-11');

explain select * from dialogs join dialogs_local on dialogs.id = dialogs_local.id where dialogs.id = 1 limit 10;

----------------------------------------------------------------------------------------------------------------------------
 Custom Scan (Citus Adaptive)  (cost=0.00..0.00 rows=0 width=0)
   ->  Distributed Subplan 13_1
         ->  Custom Scan (Citus Adaptive)  (cost=0.00..0.00 rows=0 width=0)
               Task Count: 1
               Tasks Shown: All
               ->  Task
                     Node: host=citus_worker_4 port=5432 dbname=postgres
                     ->  Index Scan using dialogs_pkey_102042 on dialogs_102042 dialogs  (cost=0.14..8.16 rows=1 width=556)
                           Index Cond: (id = 1)
   Task Count: 1
   Tasks Shown: All
   ->  Task
         Node: host=localhost port=5432 dbname=postgres
         ->  Limit  (cost=0.15..20.71 rows=5 width=1112)
               ->  Nested Loop  (cost=0.15..20.71 rows=5 width=1112)
                     ->  Index Scan using dialogs_local_pkey on dialogs_local  (cost=0.14..8.16 rows=1 width=556)
                           Index Cond: (id = 1)
                     ->  Function Scan on read_intermediate_result intermediate_result  (cost=0.00..12.50 rows=5 width=556)
                           Filter: (id = 1)
(19 rows)

32) Эксперимент, в результате получим ошибку. Текст ошибки ниже:
ERROR:  relation "dialogs_102042" already exists
CONTEXT:  while executing command on citus_worker_4:5432

CREATE TABLE dialogs_102042 (id SERIAL PRIMARY KEY,"from" UUID NOT NULL, "to" UUID NOT NULL, "text" VARCHAR(255) NOT NULL, "timestamp"  DATE NOT NULL);

SELECT create_distributed_table('dialogs_102042', 'id');