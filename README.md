
# Mini-Online-Shop

Mini olshop Service


## API Reference


#### Register

```http
  POST /user/register
```
| Body | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `email` | `string` | **Required**. Your email|
| `password` | `string` | **Required**. Your password |
| `phone` | `string` | **Required**. Your phone |
| `email` | `string` | **Required**. Your email |
| `address` | `string` | **Required**. Your address |

#### Login

```http
  POST /user/login
```
| Body | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `email` | `string` | **Required**. Your email|
| `password` | `string` | **Required**. Your password |

-------

#### Get all items

```http
  GET /product/all-products
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `api_key` | `string` | **Required**. Your API key |

#### Get All Item by Category

```http
  GET /product/all-products/:category
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `category`      | `string` | **Required**. Id of item to fetch |

---------

#### add item to cart

```http
  POST /cart/add-to-cart
```
| Headers | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `token` | `string` | **Required**. token from login|

| Body | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `product_uuid` | `string` | **Required**. product_uuid|
| `qty` | `string` | **Required**. qty of product |

#### get list cart 

```http
  GET /cart/list-cart 
```
| Headers | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `token` | `string` | **Required**. token from login|

#### Remove cart 

```http
  DELETE /cart/remove-product/:cart_uuid
```
| Headers | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `token` | `string` | **Required**. token from login|

| param | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `cart_uuid` | `string` | **Required**. uuid from list cart|

----


#### Buy 

```http
  POST /buy
```
| Headers | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `token` | `string` | **Required**. token from login|

| param | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `product_uuid` | `string` | **Required**. uuid from list product|
| `price` | `string` | **Required**. price from list product|
----
## ERD

![ERD](![Untitled Diagram](https://github.com/andryansyhh/mini-online-store/assets/80319764/8b8f32ac-db1d-44ec-8204-082c32f3001b)
)

