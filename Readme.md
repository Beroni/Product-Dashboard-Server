# Backend Getty.IO

### Challenge

The Technical Challenge consists of creating an API with Golang using gin+mongoDB that exposes CRUD endpoints of a model of your choice :heavy_check_mark:

The web app should be able to list, create, update and remove the data. :heavy_check_mark:

### Api Requeriments

- The API must guarantee the typing of user inputs. If an input is expected as a string, it can only be received as a string. :heavy_check_mark:
- The structs used with your mongo model should support Marshal/Unmarshal with bson, json and struct :heavy_check_mark:
- The API should contain unit test of methods it uses :heavy_check_mark:

### Extra

- Put some security layer (like JWT) to protect your endpoints :heavy_check_mark:

## Postman

Collection Link

https://www.getpostman.com/collections/0d2a6c09ba025cf2986f

Enviroment Link

## How to run

1. Clone Repository

```bash
git clone https://github.com/Beroni/backend-getty
```

2. Run Project

```bash
  make dev
```

2. Run tests

```bash
  make test
```
