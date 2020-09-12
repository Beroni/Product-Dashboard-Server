## Backend Getty.IO

### Challenge

The Technical Challenge consists of creating an API with Golang using gin+mongoDB that exposes CRUD endpoints of a model of your choice

The web app should be able to list, create, update and remove the data.

### Api Requeriments

- The API must guarantee the typing of user inputs. If an input is expected as a string, it can only be received as a string. :heavy_check_mark:
- The structs used with your mongo model should support Marshal/Unmarshal with bson, json and struct
- The API should contain unit test of methods it uses

### Extra

- Put some security layer (like JWT) to protect your endpoints :heavy_check_mark:
