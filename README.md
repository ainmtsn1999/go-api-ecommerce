# go-api-ecommerce
Ini merupakan dokumentasi final project untuk divisi backend Intern AI di JATIS MOBILE 
by: [Muhammad Inayatulloh](https:#www.linkedin.com/in/muhammad-inayatulloh/) 

## Skema Database
Skema database yang digunakan dalam project ini bisa dilihat detailnya di: [dbdiagram](https://dbdiagram.io/d/634d696347094101957b7381)

![screencapture-dbdiagram-io-d-634d696347094101957b7381-2022-11-30-08_40_17](https://user-images.githubusercontent.com/37493831/204686651-5b8799ab-7c62-447e-9469-c154b18bd6c5.png)

> referensi: gofood, shopee

## Architecture
- Database yang digunakan adalah `MySQL`
- Terdapat 6 services, yaitu :
  - Auth
  - User
  - Merchant
  - Product
  - Order
  - Review
- Package yg dibutuhkan :

| Nama Package | Fungsi | 
| --- | --- |
| `github.com/labstack/echo` | Framework API |
| `golang.org/x/crypto` | Hash dan Verify Password |
| `github.com/golang-jwt/jwt` | Generate dan Verify Token |
| `gorm.io/gorm` | ORM di gGolang|
| `gorm.io/drivers/mysql` | Driver untuk ORM|
| `github.com/go-playground/validator/v10` | Validator Input User |
| `github.com/joho/godotenv` | Env di Golang|

## API Spec
### Auth 
#### `POST` /auth/register
Berfungsi untuk melakukan registrasi. Fitur ini akan melakukan registrasi untuk `admin`, `user`, atau `merchant`.

**Request Body** :
```json
{
    "email" : "string", # must be a valid email
    "password" : "string",
	"role" : "string" #must be a valid role
}
```
**Notes** : Pada proses ini, `password` akan di hash menggunakan library `crypto`

**Response Body** :
```json
{
    "status" : 201, # created
    "message" : "REGISTER_SUCCESS"
}
```
Jika gagal, maka akan menghasilkan response :
```json
{
    "status" : 400, # Bad request, or others...
    "message" : "REGISTER_FAILED",
    "error" : "BAD_REQUEST"
}
```

#### `POST` /auth/login
Berfungsi untuk melakukan login seluruh pengguna

**Request Body** :
```json
{
    "email" : "string", # must be a valid email
    "password" : "string"
}
```
**Notes** : Pada proses ini, `password` akan di verify menggunakan library `crypto`

**Response Body** :
```json
{
    "status" : 200, 
    "message" : "LOGIN_SUCCESS",
    "payload" : {
        "token" : "string" # JWT 
    }
}
```
**Notes** : Pada proses ini, akan me-generate token menggunakan `JWT`

Jika gagal, maka akan menghasilkan response :
```json
{
    "status" : 400, # Bad request, or others...
    "message" : "LOGIN_FAIL",
    "error" : "BAD_REQUEST"
}
```
### User
#### `POST` /users/profile
Berfungsi untuk create user

**Request Headers**
```bash
Authorization : Bearer <token>
```

**Request Body**
```json
{
    "name": "string",
    "gender": "string",
    "phone_number": "string", #Phone Number
    "pict_url": "string" #Pict Url
}
```

**Response Body**
```json
{
    "status" : 201,
    "message" : "CREATED_USER_SUCCESS"
}
```
Jika gagal, maka akan menghasilkan response :
```json
{
    "status" : 400, # Bad request, or others...
    "message" : "CREATED_USER_FAILED",
    "error" : "BAD_REQUEST"
}
```

#### `PUT` /users/profile
Berfungsi untuk mengubah data user yang sedang login

**Request Headers**
```bash
Authorization : Bearer <token>
```

**Request Body**
```json
{
    "name": "string",
    "gender": "string",
    "phone_number": "08xxxxxxxxxxx", #Phone Number
    "pict_url": "string" #Pict Url
}
```

**Response Body**
```json
{
    "status" : 200,
    "message" : "UPDATE_USER_SUCCESS"
}
```
Jika gagal, maka akan menghasilkan response :
```json
{
    "status" : 400, # Bad request, or others...
    "message" : "UPDATE_USER_FAIL",
    "error" : "BAD_REQUEST"
}
```

#### `GET` /users/profile
Akan menampilkan detail profile dari user yang sedang login.

**Request Headers**
```bash
Authorization : Bearer <token>
```

**Response Body**
```json
{
    "status" : 200,
    "message" : "GET_USER_PROFILE_SUCCESS",
    "payload": {
        "id": 1,
        "name": "string",
        "gender": "string",
        "phone_number": "08xxxxxxxxxxx",
        "pict_url": "string",
        "address": {
            "address_tag": "string",
            "city_id": "string",
            "province_id": "string",
            "street": "string"
        },
        "auth": {
            "email": "test@test.com"
        }
    }
}
```
Jika gagal, maka akan menghasilkan response :
```json
{
    "status" : 401, # Unauthorized, Not Found, or others...
    "message" : "GET_USER_PROFILE_FAILED",
    "error" : "UNAUTHORIZED"
}
```

#### `GET` /users
Akan menampilkan seluruh data users. Endpoint ini hanya berlaku untuk `admin`.

**Request Headers**
```bash
Authorization : Bearer <token>
```

**Query String**
- limit : `int` with default is 25  | optional
- page  : `int` with default is 1   | optional

**Response Body**
```json
{
    "status" : 200,
    "message" : "GET_ALL_USERS_SUCCESS",
    "payload" : [
        {
			"id": 1,
			"name": "string",
			"gender": "string",
			"phone_number": "08xxxxxxxxxxx",
			"pict_url": "string",
			"address": {
				"address_tag": "string",
				"city_id": "string",
				"province_id": "string",
				"street": "string"
			},
			"auth": {
				"email": "test@test.com"
			}
		}
    ],
    "query" : {
        "limit" : 25, # default is 25
        "page" : 1, # default is 1
        "total" : 1
    }
}
```

Jika gagal, maka akan menghasilkan response :
```json
{
    "status" : 403, # Forbidden, Not Found, or others...
    "message" : "GET_ALL_USERS_FAILED",
    "error" : "FORBIDDEN_ACCESS"
}
```

#### `POST` /users/profile/address
#### `PUT` /users/address/id/:id
#### `DELETE` /users/address/id/:id
#### `GET` /users/address/id/:id
#### `GET` /users/profile/address
#### `PUT` /users/address/id/:id/activate

### Merchant
#### `GET` /merchants 
#### `GET` /merchants/profile
#### `POST` /merchants/profile
#### `PUT` /merchants/profile

###	Product
#### `POST` /products
#### `PUT` /products/id/:id
#### `DELETE` /products/id/:id
#### `GET` /products
#### `GET` /products/merchant/:id
#### `GET` /products/id/:id

###	Order
#### `POST` /orders/inquire
#### `POST` /orders/confirm
#### `PUT` /orders/id/:id/status
#### `GET` /orders/id/:id
#### `GET` /orders/histories/me
#### `GET` /orders/histories/list
#### `PUT` /orders/id/:orderId/product/:productId/status

###	Review
#### `POST` /reviews
#### `GET` /reviews
#### `GET` /reviews/id/:id
#### `GET` /reviews/product/:id
#### `GET` /reviews/order/:id
#### `GET` /reviews/user/:id