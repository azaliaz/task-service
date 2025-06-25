# **Сервис задач**

## Запуск

- Запустить сервис можно с помощью команды `make deploy`
- Останавливает и удаляет развернутое окружение Docker вместе с его образами `make rollback`


## Примеры запросов
### Создание задачи
```
curl -X POST http://localhost:8080/tasks      
```

### Запрос на результат задачи
```
curl http://localhost:8080/tasks/{id}
```

### Удаление задачи
``` 
curl -X DELETE http://localhost:8080/tasks/{id}
```

