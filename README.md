Структура проекта:
    api:
        docs - документация swagger
    cmd:
        main.go - точка входа в программу , обьявление самых основных структур
    config:
        config.go - чтение файла .env и запись переменных окружения в Config{}
    docker:
        Dockerfile
    intern:
        authService:
            delivery - реализация handlers и роутеров
            repository - реализация authRepository для работы с db
            usecase - реализация authUseCase (бизнес правила приложения)
        emailService - моковые данные email сервиса записывает письма внутри себя в map[user_id]email 
        models - модели используемые в intern (некоторые используются только для документации swagger)
        server - настройка а также запуск сервера
    migrations - миграция базы данных перед запуском сервера, запускается в cmd/main.go
    pkg:
        db - реализация подключения к базе данных
        httpErrors - юзерские ошибки 
        logger - базовая реализация логгера 
        utils - содержит работу с контекстом запроса, валидацию структур и генерацию access и refresh токенов
    .env - заполнен, можно не менять
    docker-compose.yml - запуск контейнера с postgres и самим сервером
    
Запуск проекта:
`docker compose up --build`

Тестирование проекта:
Для удобства подключил swagger и он доступен по пути /swagger/
api/v1/auth/tokens :
    В файле миграции создается юзер с user_id=b3ac3626-7e37-4026-b789-9a081b252dd1
    При необходимости просто добавить нового юзера в migrations/202411... со своим user_id uuid.UUID
    Есть возможность указывания в headers IPAddress, из-за использования docker сервер будет видеть будто все запросы делаются с одного ip adress, подключать nginx для такого задания посчитал излишним решением.

api/v1/auth/tokens/refresh :
    Вставляем данные которые получим при использовании /token

Написал также интеграционные тесты, для запуска:
`go test ./...`

Для проверки покрытия тестами:
`go test -coverprofile=coverage.out ./...`

Выполнение условий тз:
 - Access token (jwt) генерируется при помощи jwt.SigningMethodHS512 (см. pkg/utils/tokens.go) и не храниться в базе
 - Refresh token передается в форме base64 и в базе храниться исключительно в виде bcrypt хэша
 - Access и Refresh token взаимно связаны при помощи refresh_token_id и поэтому можно сделать refresh только для того Access токена с которым он связан, refresh обновляет так же refresh_token_id сделано для безопасности. 
 - Payload токенов содержит ip адресс и при его изменении отправляется сообщение на моковый emailService (пришедшее сообщение видно в логах)

При возникновении ошибок или вопросов пишите на venskiandre32@gmail.com либо в telegram @ban_ka

