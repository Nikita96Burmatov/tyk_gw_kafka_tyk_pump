Реализация доставки асинхронных запросов средствами Tyk-gateway + kafka + redis + Tyk-pump + prometheus
1 Tyk-gateway принимает запрос отправляет в Kafka
2 Redis складывается информация о запросе 
3 Tyk-pump собирает данные из Redis
4 Prometheus забирает метрики из Tyk-pump что позволяет отслеживать состояния запроса