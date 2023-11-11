# sushi_api

This is a simple sushi API for study purposes.

### Routes

| Method      | Path        | Type|
| ----------- | ----------- |-----|
| GET      | [/](https://sushiapi-production.up.railway.app/)      | String|
| GET   | [/sushi](https://sushiapi-production.up.railway.app/sushi)       | Sushi |

### Types

- Sushi - example
```json
{
  "_id": "asbd2432",
  "price": 12,
  "discount": 2,
  "image": "https://image",
  "ingredients": [
    {
    "ingredient_name": "nmae",
    "calories": 100
    },
  ],
  "rating": 4
},
```
