# go-api-ecommerce
Ini merupakan dokumentasi final project untuk divisi backend Intern AI di JATIS MOBILE. 
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
Berfungsi untuk create user. Endpoint ini hanya berlaku untuk `user`.

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
    "message" : "CREATE_USER_SUCCESS"
}
```
Jika gagal, maka akan menghasilkan response :
```json
{
    "status" : 400, # Bad request, or others...
    "message" : "CREATE_USER_FAILED",
    "error" : "BAD_REQUEST"
}
```

#### `PUT` /users/profile
Berfungsi untuk mengubah data user yang sedang login. Endpoint ini hanya berlaku untuk `user`.

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
    "message" : "UPDATE_USER_FAILED",
    "error" : "BAD_REQUEST"
}
```

#### `GET` /users/profile
Akan menampilkan detail profile dari user yang sedang login. Endpoint ini hanya berlaku untuk `user`.

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
Berfungsi untuk create data alamat user. Endpoint ini hanya berlaku untuk `user`.

**Request Headers**
```bash
Authorization : Bearer <token>
```

**Request Body**
```json
{
    "street": "string",
    "city_id": "string",
    "province_id": "string",
    "address_tag": "string"
}
```
**Notes** : Pada proses ini, akan diberikan default `"activate" : "y"`, jika tidak ada alamat terdaftar pada akun user.

**Response Body**
```json
{
    "status" : 201,
    "message" : "CREATE_USER_ADDRESS_SUCCESS"
}
```
Jika gagal, maka akan menghasilkan response :
```json
{
    "status" : 400, # Bad request, or others...
    "message" : "CREATE_USER_ADDRESS_FAILED",
    "error" : "BAD_REQUEST"
}
```

#### `PUT` /users/address/id/:id
Berfungsi untuk mengubah data alamat user. Endpoint ini hanya berlaku untuk `user`.

**Params**
- id : `int` | required

**Request Headers**
```bash
Authorization : Bearer <token>
```

**Request Body**
```json
{
    "street": "string",
    "city_id": "string",
    "province_id": "string",
    "address_tag": "string"
}
```

**Response Body**
```json
{
    "status" : 200,
    "message" : "UPDATE_USER_ADDRESS_SUCCESS"
}
```
Jika gagal, maka akan menghasilkan response :
```json
{
    "status" : 400, # Bad request, or others...
    "message" : "UPDATE_USER_ADDRESS_FAILED",
    "error" : "BAD_REQUEST"
}
```

#### `DELETE` /users/address/id/:id
Berfungsi untuk menghapus data alamat user. Endpoint ini hanya berlaku untuk `user`

**Params**
- id : `int` | required

**Request Headers**
```bash
Authorization : Bearer <token>
```

```json
{
    "status" : 200,
    "message" : "DELETE_USER_ADDRESS_SUCCESS"
}
```
**Notes** : Pada proses ini, data alamat user yang `activate` tidak akan bisa dihapus.

Jika gagal, maka akan menghasilkan response :
```json
{
    "status" : 403, # Bad request, or others...
    "message" : "DELETE_USER_ADDRESS_FAILED",
    "error" : "FORBIDDEN_ACCESS"
}
```

#### `GET` /users/address/id/:id
Akan menampilkan detail data alamat user. Endpoint ini hanya berlaku untuk `user`.

**Params**
- id : `int` | required
  
**Request Headers**
```bash
Authorization : Bearer <token>
```

**Response Body**
```json
{
    "status": 200,
    "message": "GET_USER_ADDRESS_SUCCESS",
    "payload": {
        "id": 1,
        "user": {
            "id": 1,
            "name": "string",
            "phone_number": "08xxxxxxxxxx"
        },
        "address": {
            "city_id": "string",
            "province_id": "string",
            "street": "string"
        },
        "address_tag": "string",
        "activate": "n",
        "created_at": "datetime",
        "updated_at": "timestamp",
        "deleted_at": "datetime"
    }
}
```
Jika gagal, maka akan menghasilkan response :
```json
{
    "status" : 401, # Unauthorized, Not Found, or others...
    "message" : "GET_USER_ADDRESS_FAILED",
    "error" : "UNAUTHORIZED"
}
```

#### `GET` /users/profile/address
Akan menampilkan seluruh data alamat user yang sedang login. Endpoint ini hanya berlaku untuk `user`.

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
    "status": 200,
    "message": "GET_ALL_USER_ADDRESSES_SUCCESS",
    "payload": [
        {
            "id": 1,
            "user": {
                "id": 1,
                "name": "string",
                "phone_number": "08xxxxxxxxxx"
            },
            "address": {
                "city_id": "string",
                "province_id": "string",
                "street": "string"
            },
            "address_tag": "string",
            "activate": "y",
            "created_at": "datetime",
            "updated_at": "timestamp",
            "deleted_at": "datetime"
        }
    ],
    "query": {
        "limit": 25,
        "page": 1,
        "total": 1
    }
}
```

Jika gagal, maka akan menghasilkan response :
```json
{
    "status" : 403, # Forbidden, Not Found, or others...
    "message" : "GET_ALL_USER_ADDRESSES_FAILED",
    "error" : "FORBIDDEN_ACCESS"
}
```

#### `PUT` /users/address/id/:id/activate
Ini berfungsi untuk mengubah status data alamat user yang aktif digunakan. Endpoint ini hanya bisa di akses oleh `user`

**Params**
- id : `int` | required
  
**Request Headers**
```bash
Authorization : Bearer <token>
```

**Request Body**
```json
{
    "activate": "y"
}
```
**Notes** : Pada proses ini, data alamat yang memiliki status activate aktif akan dinon-aktifkan. Jadi tiap user hanya memiliki satu data alamat aktif.

List Aktifasi:
- `"y"` : aktif
- `"n"` : non-aktif

**Response Body**
```json
{
    "status": 202,
    "message": "UPDATE_ACTIVATE_ADDRESS_SUCCESS"
}
```

### Merchant
#### `POST` /merchants/profile
Berfungsi untuk create merchant. Endpoint ini hanya berlaku untuk `merchant`.

**Request Headers**
```bash
Authorization : Bearer <token>
```

**Request Body**
```json
{
    "name": "string",
    "phone_number": "08xxxxxxxxxx",
    "street": "string",
    "city_id": "string",
    "province_id": "string",
    "pict_url": "string"
}
```

**Response Body**
```json
{
    "status" : 201,
    "message" : "CREATE_MERCHANT_SUCCESS"
}
```
Jika gagal, maka akan menghasilkan response :
```json
{
    "status" : 400, # Bad request, or others...
    "message" : "CREATE_MERCHANT_FAILED",
    "error" : "BAD_REQUEST"
}
```

#### `PUT` /merchants/profile
Berfungsi untuk mengubah data merchant yang sedang login. Endpoint ini hanya berlaku untuk `merchant`.

**Request Headers**
```bash
Authorization : Bearer <token>
```

**Request Body**
```json
{
    "name": "string",
    "phone_number": "08xxxxxxxxxx",
    "street": "string",
    "city_id": "string",
    "province_id": "string",
    "pict_url": "string"
}
```

**Response Body**
```json
{
    "status" : 200,
    "message" : "UPDATE_MERCHANT_SUCCESS"
}
```
Jika gagal, maka akan menghasilkan response :
```json
{
    "status" : 400, # Bad request, or others...
    "message" : "UPDATE_MERCHANT_FAILED",
    "error" : "BAD_REQUEST"
}
```

#### `GET` /merchants 
Akan menampilkan seluruh data merchants. Endpoint ini hanya berlaku untuk `admin`.

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
    "status": 200,
    "message": "GET_ALL_MERCHANTS_SUCCESS",
    "payload": [
        {
            "id": 2,
            "name": "string",
            "phone_number": "08xxxxxxxxxxx",
            "pict_url": "string",
            "address": {
                "city_id": "string",
                "province_id": "string",
                "street": "string"
            },
            "auth": {
                "email": "test1@test.com"
            }
        }
    ],
    "query": {
        "limit": 25,
        "page": 1,
        "total": 1
    }
}
```

Jika gagal, maka akan menghasilkan response :
```json
{
    "status" : 403, # Forbidden, Not Found, or others...
    "message" : "GET_ALL_MERCHANTS_FAILED",
    "error" : "FORBIDDEN_ACCESS"
}
```

#### `GET` /merchants/profile
Akan menampilkan detail profile dari merchant yang sedang login. Endpoint ini hanya berlaku untuk `merchant`.

**Request Headers**
```bash
Authorization : Bearer <token>
```

**Response Body**
```json
{
    "status": 200,
    "message": "GET_MERCHANT_PROFILE_SUCCESS",
    "payload": {
        "id": 2,
        "name": "string",
        "phone_number": "08xxxxxxxxxx",
        "pict_url": "string",
        "address": {
            "city_id": "string",
            "province_id": "string",
            "street": "string"
        },
        "auth": {
            "email": "test1@test.com"
        }
    }
}
```
Jika gagal, maka akan menghasilkan response :
```json
{
    "status" : 401, # Unauthorized, Not Found, or others...
    "message" : "GET_MERCHANT_PROFILE_FAILED",
    "error" : "UNAUTHORIZED"
}
```

###	Product
#### `POST` /products
Berfungsi untuk create product dari merchant yang sedang login. Endpoint ini hanya berlaku untuk `merchant`.

**Request Headers**
```bash
Authorization : Bearer <token>
```

**Request Body**
```json
{
    "name" : "string",
    "category" : "string",
    "desc" : "string",
    "price" : 10000000, #int 
    "stock" : 10, #int
    "weight" : 1000, #int
    "img_url" : "string" 
}
```

**Response Body**
```json
{
    "status" : 201,
    "message" : "CREATE_PRODUCT_SUCCESS"
}
```
Jika gagal, maka akan menghasilkan response :
```json
{
    "status" : 400, # Bad request, or others...
    "message" : "CREATE_PRODUCT_FAILED",
    "error" : "BAD_REQUEST"
}
```

#### `PUT` /products/id/:id
Berfungsi untuk mengubah data product. Endpoint ini hanya berlaku untuk `merchant`.

**Params**
- id : `int` | required
  
**Request Headers**
```bash
Authorization : Bearer <token>
```

**Request Body**
```json
{
    "name" : "string",
    "category" : "string",
    "desc" : "string",
    "price" : 10000000, #int 
    "stock" : 10, #int
    "weight" : 1000, #int
    "img_url" : "string" 
}
```

**Response Body**
```json
{
    "status" : 200,
    "message" : "UPDATE_PRODUCT_SUCCESS"
}
```
Jika gagal, maka akan menghasilkan response :
```json
{
    "status" : 400, # Bad request, or others...
    "message" : "UPDATE_PRODUCT_FAILED",
    "error" : "BAD_REQUEST"
}
```

#### `DELETE` /products/id/:id
Berfungsi untuk menghapus data product. Endpoint ini hanya berlaku untuk `merchant`

**Params**
- id : `int` | required

**Request Headers**
```bash
Authorization : Bearer <token>
```

```json
{
    "status" : 200,
    "message" : "DELETE_PRODUCT_SUCCESS"
}
```

Jika gagal, maka akan menghasilkan response :
```json
{
    "status" : 403, # Bad request, or others...
    "message" : "DELETE_PRODUCT_FAILED",
    "error" : "FORBIDDEN_ACCESS"
}
```

#### `GET` /products
Akan menampilkan seluruh data products. Endpoint ini bisa diakses seluruh pengguna tanpa login.

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
    "status": 200,
    "message": "GET_ALL_PRODUCTS_SUCCESS",
    "payload": [
        {
            "id": 1,
            "merchant": {
                "id": 2,
                "name": "string"
            },
            "category": "string",
            "name": "string",
            "desc": "string",
            "price": 10000000,
            "stock": 10,
            "weight": 1000,
            "img_url": "string"
        }
    ],
    "query": {
        "limit": 25,
        "page": 1,
        "total": 1
    }
}
```

Jika gagal, maka akan menghasilkan response :
```json
{
    "status" : 403, # Forbidden, Not Found, or others...
    "message" : "GET_ALL_PRODUCTS_FAILED",
    "error" : "FORBIDDEN_ACCESS"
}
```

#### `GET` /products/merchant/:id
Akan menampilkan seluruh data products. Endpoint ini bisa diakses seluruh pengguna tanpa login.

**Params**
- id : `int` | required
  
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
    "status": 200,
    "message": "GET_ALL_MERCHANT_PRODUCTS_SUCCESS",
    "payload": [
        {
            "id": 1,
            "merchant": {
                "id": 2,
                "name": "string"
            },
            "category": "string",
            "name": "string",
            "desc": "string",
            "price": 10000000,
            "stock": 10,
            "weight": 1000,
            "img_url": "string"
        }
    ],
    "query": {
        "limit": 25,
        "page": 1,
        "total": 1
    }
}
```

Jika gagal, maka akan menghasilkan response :
```json
{
    "status" : 403, # Forbidden, Not Found, or others...
    "message" : "GET_ALL_MERCHANT_PRODUCTS_FAILED",
    "error" : "FORBIDDEN_ACCESS"
}
```

#### `GET` /products/id/:id
Akan menampilkan detail product yang dipilih. Endpoint ini bisa diakses seluruh pengguna tanpa login.

**Params**
- id : `int` | required
  
**Request Headers**
```bash
Authorization : Bearer <token>
```

**Response Body**
```json
{
    "status": 200,
    "message": "GET_PRODUCT_SUCCESS",
    "payload": {
        "id": 1,
        "merchant": {
            "id": 2,
            "name": "string"
        },
        "category": "string",
        "name": "string",
        "desc": "string",
        "price": 10000000,
        "stock": 10,
        "weight": 1000,
        "img_url": "string"
    }
}
```
Jika gagal, maka akan menghasilkan response :
```json
{
    "status" : 401, # Unauthorized, Not Found, or others...
    "message" : "GET_PRODUCT_FAILED",
    "error" : "UNAUTHORIZED"
}
```

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