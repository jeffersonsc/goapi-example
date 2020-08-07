## Nature API

### Api of example in Go

### System requirements

* Go 1.13x
* Docker 18x
* Docker Compose 1.25x

### Application requirements

* MongoDB lastest version
* Redis lastest version

### System commands

* Start application `$ make run`
* Build docker `$ docker built -t {your name} .`
* Start docker compose `$ docker-compose up -d`

### Envirolment variables

* MONGO_URL - set your mongo database url
* REDIS_URL - set redis url
* PORT - set port when server execute

obs.: In developer mode and makefile has be configurated

### Developer guide

Create a new feature or fix using style guide down:

* feature/{name of feature}
* feature-fix/{name of fix feature}
* bug-fix/{name your case or order ticket}
* hot-fix/{name of case or order ticket case}

Create a new pull request and request reviwers

### Api reference

Allowed methods:

* GET /v1/products - Return all products
* GET /v1/products/:id - Return a product by id
* POST /v1/products - Create a new product
* PUT /v1/products/:id - Update product by id
* DELETE /v1/products/:id - Delele a product by id WARNING this method remove data from database

Response types:

|Status|Description                           |
|------|--------------------------------------|
|200   |Success ok                            |
|400   |Invalid request                       |
|404   |Resource not found or data not found  |
|422   |Used on validate input data invalid   |
|500   |Internal server error                 |

JSON Input:

```JSON
{
    "name": "My Product",
    "description":"My product description",
    "price": 99.99,
    "currency_code": "BRL",
    "images": [
        "https://hbr.org/resources/images/article_assets/2017/10/oct17-11-740519323.jpg",
    ]
}
```

### Especial area, links of references

* [Clean architecture](https://medium.com/@eminetto/clean-architecture-using-golang-b63587aa5e3f)
* [Developer API with validate JWT](https://developer.okta.com/blog/2018/10/23/build-a-single-page-app-with-go-and-vue)
* [MongoDB References](https://www.mongodb.com/blog/post/quick-start-golang--mongodb--modeling-documents-with-go-data-structures)
* [Radix redis driver](https://godoc.org/github.com/mediocregopher/radix)