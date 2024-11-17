# Энергогарант

## Структура проекта

```
kafka
├── broker                  // Имплементация брокера
├── config                  // Конфигурационные файлы
├── handlers                // http хэндлеры
│   └── producer
└── tests                   // тесты
```

## Схема Apache Kafka

У нас есть 1 брокер, 1 продюсер, 2 консюмер группы по 1 консюмеру.
Название топика зависит от того, куда надо отправить данные

![Схема Apache Kafka](https://gitlab.cbitrix.com/clients/eg-tyk-gw/-/blob/tyk-gateway/internal/kafka/docs/kafka-scheme.jpg)

## Запуск команд

dev

```
go run cmd/producer/main.go --config=./config/dev.yaml
```

prod

```
go run cmd/producer/main.go --config=./config/prod.yaml
```
