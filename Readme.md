Сервис событий
--------------

Ваша задача написать сервис, который реализует API создания и завершения разных событий.
События могу быть разных типов. В базе данных может быть только одно незавершенное событие одного типа.

### API

*POST /v1/start*

Тело запроса

    {
        "type": "..." // тип события
    }

Метод создает событие переданного типа. Если в базе уже есть незавершенное событие переданного типа, то новое событие создавать не нужно, и в ответ не должно приходить сообщение об ошибке.

*POST /v1/finish*

    {
        "type": "..." // тип события
    }

Метод завершает событие переданного типа

### БД

База данных должна быть MongoDB.
События всех типов должны хранится отдельными документами в одной коллекции

Пример документа в коллекции:

     {
        "_id": "...",
        "type": "...",
        "state": 0 // state 0 - для незавершенных событий , 1 - для завершенных
    }
    
## Комментарии к задаче
В условии не указано ожидаемое поведение в случае, если приходит начало события, типа, который раньше был завершен. 
Будем считать, что в таком случае создается не новое событие, а старое событие становится незавершенным (state=0).  

## Запуск сервиса
Сервис запускается командой `go run ./cmd/events/main.go ./cmd/events/app.go`, при этом необходимо сначала корректно 
настроить конфигурацию в файле .env и запустить `dep ensure` для загрузки необходимых библиотек.
Приложение работает с коллекцией `events` mongoDB.