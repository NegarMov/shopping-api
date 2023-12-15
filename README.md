<h1 align="center">Shopping API</h1>
<h6 align="center">Fall-2023 Internet Engineering Course Midterm Exam at Amirkabir University of Tech.</h6>


## Introduction
This is a simple shopping server which uses echo, JWT authentication, and PostgreSQL as its database. You can sign in as a user and manage your shopping basket.
The API endpoints are as follows:
### /user
  * `[POST]` **/signup** - Is used to sign up a new user. The input should provide user's username and password in JSON format.
  * `[POST]` **/signin** - Users can use this endpoint to sign in. It return a JWT token if the credentials are OK.
### /basket
  * `[GET]` **/** - Get all of the items in the user's basket. User authorization is done by the JWT token provided.
  * `[GET]` **/:id** - Get the item with the specified `id`. User authorization is done by the JWT token provided.
  * `[POST]` **/** - Add a new item to the user's basket. User authentication is done by the JWT token provided.
  * `[PATCH]` **/:id** - Update an item with the `PENDING` status. User authorization is done by the JWT token provided.
  * `[DELETE]` **/:id** - Delete the item with the specified `id`. User authorization is done by the JWT token provided.

**NOTE:** for using the current version, also add `/v1` before each of the paths specified.

## Running the Server

Build and run the server simply by using the following commands:

```bash
go build .
./shopping-api
```
Don't forget to change the config file according to your system.
