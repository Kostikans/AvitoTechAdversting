## AvitoTech Advertising.

## Принятые решения 
- добавил пагинацию через offset limit и запрашивание страничек с помощью query параметра "page"

## Запуск проекта 
make start (локальная постгреха должна быть остановлена "sudo service postgresql stop")

## Запуск тестов
make tests (покрытие около 75%)

## документация 
[Swagger-doc](http://localhost:9000/docs/index.html)

## Архитектура сервиса
![Service Architecture](https://github.com/Kostikans/AvitoTechAdvertising/raw/master/diagram.jpg)

## p.s.
env файл закинул в репу чтобы можно было сразу протестить проект, так то его понятно там быть не должно)