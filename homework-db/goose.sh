# Добавление файла миграций 
goose -dir homework-db/migrations create new_user_table sql
# Приминение миграций из папки
goose -dir homework-db/migrations postgres "postgresql://postgres:example@127.0.0.1:5432/otus_homework?sslmode=disable" up
# Откат миграции
goose -dir homework-db/migrations postgres "postgresql://postgres:example@127.0.0.1:5432/otus_homework?sslmode=disable" down-to 20240302213613