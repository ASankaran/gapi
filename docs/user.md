**POST /user/**

Register an user

| Header          | Description                  | Required  |
| --------------- | ---------------------------- | --------- |
| `Content-Type`  | application/json             | Yes       |

Request Body

````
{
  "email": "email@domain.com",
  "password": "pass"
}
````

Response Body

````
{
    "ID": 3,
    "CreatedAt": "2018-03-20T22:21:57.61401963-05:00",
    "UpdatedAt": "2018-03-20T22:21:57.61401963-05:00",
    "DeletedAt": null,
    "email": "email@domain.com"
}
````

**GET /user/**

Gets the logged in user

| Header          | Description                  | Required  |
| --------------- | ---------------------------- | --------- |
| `Content-Type`  | application/json             | Yes       |
| `Authorization` | a valid authentication token | Yes       |

Response Body

Same as response when originally creating the user

**GET /user/{id}**

Gets the logged in user

| Header          | Description                  | Required  |
| --------------- | ---------------------------- | --------- |
| `Content-Type`  | application/json             | Yes       |
| `Authorization` | a valid authentication token | Yes       |

Response Body

Similar for the response for GET /user/ but for the attendee with the given id

**POST /user/login**

Login the user and receive a api token

| Header          | Description                  | Required  |
| --------------- | ---------------------------- | --------- |
| `Content-Type`  | application/json             | Yes       |

Request Body

````
{
  "email": "email@domain.com",
  "password": "pass"
}
````

Response Body

````
{
    "token": "eyJhbGciOiksjdusNiIsInR5cCI6IkpXVCJ9.eyJleHAijsyqnlE4NjE5MTgsImlkIjoxfQ.P5mzvaR-WkQB5QGP5HPDa3e,nhr6jZNXHSEqf4khfoA"
}
````
