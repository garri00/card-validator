# Card validator

Request
```
{
    "cardNumber": "400010201203",
    "expirationMont" : "01",
    "expirationYear": "2021"
}
```

Response
```
{
    "valid": false,
    "error": {
        "code": 110,
        "message": "card number not vatid luhn validation"
    }
}
```

