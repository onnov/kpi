# Реализация методов:
- /_api/facts/save_fact
- /_api/indicators/get_facts

# Для разработки используется mariadb в docker
### Запустить docker:
- `docker compose up -d`
- или `make up`

### Остановить docker:
- `docker compose down`
- или `make down`

# Миграции
Для управленя миграциями используется инструент https://github.com/pressly/goose  
Установка `go install github.com/pressly/goose/v3/cmd/goose@latest`  
Миграция создана в каталоге `migrations`
Нужно "накатить" миграцию `./migrate.sh`

# Параметры запуска задаются переменными окружения
- `LISTENED_ADDRESS=":3000"` ip:port на котором будет "слушать" наш сервер
- `DB_HOST=127.0.0.1:3366` db host:port
- `DB_NAME=kpi` db name
- `DB_USER=kpi` db user
- `DB_PASSWORD=kpi` db password

# dev Запуск
Скрипт `run.sh` билдит, устанавливает переменные окружения и запускает 

# Запросы для тестирования:
```shell
curl --request POST \
  --url http://127.0.0.1:3000/_api/facts/save_fact \
  --header 'Authorization: Bearer 48ab34464a5573519725deb5865cc74c' \
  --header 'Content-Type: application/x-www-form-urlencoded' \
  --header 'User-Agent: insomnia/2023.5.8' \
  --data period_start=2024-05-01 \
  --data period_end=2024-05-30 \
  --data period_key=month \
  --data indicator_to_mo_id=227373 \
  --data indicator_to_mo_fact_id=0 \
  --data value=1 \
  --data fact_time=2024-05-31 \
  --data is_plan=0 \
  --data auth_user_id=400 \
  --data 'comment=buffer Last_name'
```
```shell
curl --request POST \
  --url http://127.0.0.1:3000/_api/indicators/get_facts \
  --header 'Authorization: Bearer 48ab34464a5573519725deb5865cc74c' \
  --header 'Content-Type: application/x-www-form-urlencoded' \
  --header 'User-Agent: insomnia/2023.5.8' \
  --data period_start=2024-05-01 \
  --data period_end=2024-05-30 \
  --data period_key=month \
  --data indicator_to_mo_id=227373
```
