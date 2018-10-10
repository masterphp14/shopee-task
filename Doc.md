# Tax Calc
----------------------------------------------------
using : github.com/joho/godotenv,github.com/labstack/echo,github.com/jinzhu/gorm
----------------------------------------------------
Get All Order, Get One Order, and Create Order
## Get One Order
```
GET localhost:port/order/14
```

### Response Fail
```json
{
    "status": "record not found"
}
```

### Response Success
```json
{
    "status": "success",
    "data": {
        "id": 14,
        "no_order": "wer2",
        "date": "2018-03-12T04:09:39+07:00",
        "items": [
            {
                "id": 5,
                "order": {
                    "id": 14,
                    "no_order": "wer2",
                    "date": "2018-03-12T04:09:39+07:00"
                },
                "order_id": 14,
                "item_name": "coklat",
                "tax_code": {
                    "id": 3,
                    "code": "3",
                    "type": "Entertainment",
                    "refundable": "no"
                },
                "tax_code_id": 3,
                "price": 3000,
                "tax": 29,
                "amount": 3029
            },
            {
                "id": 6,
                "order": {
                    "id": 14,
                    "no_order": "wer2",
                    "date": "2018-03-12T04:09:39+07:00"
                },
                "order_id": 14,
                "item_name": "coklat2",
                "tax_code": {
                    "id": 1,
                    "code": "1",
                    "type": "Food & Beverage",
                    "refundable": "yes"
                },
                "tax_code_id": 1,
                "price": 2000,
                "tax": 200,
                "amount": 2200
            }
        ],
        "price_subtotal": 5000,
        "tax_subtotal": 229,
        "grand_subtotal": 5229
    },
    "total": 1
}
```

## Create Order
```
POST localhost:port/order

```

### Request
```json
{
    "no_order": "wer2",
    "date": "2018-03-12T04:09:39+07:00",
    "items": [
            	{
                	"item_name": "coklat",
                	"tax_code_id": 3,
                	"price": 3000
            	},
            	{
                	"item_name": "coklat2",
                	"tax_code_id": 1,
                	"price": 2000
            	}
        ]
}

```
### Response Success
```json
{
    "status": "success",
    "data": {
        "id": 14,
        "no_order": "wer2",
        "date": "2018-03-12T04:09:39+07:00",
        "items": [
            {
                "id": 5,
                "order": {
                    "id": 14,
                    "no_order": "wer2",
                    "date": "2018-03-12T04:09:39+07:00"
                },
                "order_id": 14,
                "item_name": "coklat",
                "tax_code": {
                    "id": 3,
                    "code": "3",
                    "type": "Entertainment",
                    "refundable": "no"
                },
                "tax_code_id": 3,
                "price": 3000,
                "tax": 29,
                "amount": 3029
            },
            {
                "id": 6,
                "order": {
                    "id": 14,
                    "no_order": "wer2",
                    "date": "2018-03-12T04:09:39+07:00"
                },
                "order_id": 14,
                "item_name": "coklat2",
                "tax_code": {
                    "id": 1,
                    "code": "1",
                    "type": "Food & Beverage",
                    "refundable": "yes"
                },
                "tax_code_id": 1,
                "price": 2000,
                "tax": 200,
                "amount": 2200
            }
        ],
        "price_subtotal": 5000,
        "tax_subtotal": 229,
        "grand_subtotal": 5229
    },
    "total": 1
}

### Response Fail
```json
{
    "status": "invalid",
    "messages": {
        "item_name.1": "cannot be empty",
        "tax_code_id.0": "record not found"
    }
}