## Запуск
### Через Docker
1. Склонируйте репозиторий

2. В корне проекта, в консоли выполните команду:
```
docker compose -f docker-compose.yaml up -d
```

3. Также в консоли выполните команду:
```
docker exec -it client bash
```

4. Запустите клиента командой:
```
./main -a=server -p=7070
```

5. Done
