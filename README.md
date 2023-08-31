# Тестовое задание на позицию стажера в Avito по направлению Backend-разработка
## Запуск
Требуется Docker.
### Запустить сервис
```shell
make up
```

### Остановить сервис
```shell
make down
```

## Примеры API

### Создать сегмент
```shell
curl --location 'http://localhost:8080/segments' \
--header 'Content-Type: application/json' \
--data '{
    "Slug": "TEST_SEGMENT"
}'
```

### Удалить сегмент
```shell
curl --location --request DELETE 'http://localhost:8080/segments/TEST_SEGMENT' \
--data ''
```

### Обновить данные пользователя
```shell
curl --location --request PUT 'http://localhost:8080/users/1/segments' \
--header 'Content-Type: application/json' \
--data '{
    "segments_add": [
        "s1", "s2"
    ],
    "segments_del": []
}'
```

### Получить сегменты пользователя
```shell
curl --location 'http://localhost:8080/users/1/segments'
```