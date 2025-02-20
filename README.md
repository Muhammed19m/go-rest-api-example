# golang junior dev



## Rest api

Rest api обрабатывающий 3 вида запрос:
- получение баланса по id
- положить деньги на счет
- взять деньги со счета

``` http
GET http://localhost:8000/api/v1/wallets/1 HTTP/1.1
```
  
``` http
POST http://localhost:8000/api/v1/wallet HTTP/1.1
content-type: application/json
{
    "walletId": 1,
    "operationType": "DEPOSIT",
    "amount": 10
}
```
  
  
``` http
POST http://localhost:8000/api/v1/wallet HTTP/1.1
content-type: application/json
{
    "walletId": 1,
    "operationType": "WITHDRAW",
    "amount": 10
}
```

## Запуск приложения 

```
docker-compose up
```


