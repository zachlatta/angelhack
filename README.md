# AngelHack LA

This is the backend server for our AngelHack LA hack.

## API Docs

### Users

#### Create account

    POST /users

##### Request

```json
{
  "firstName": "Zaphod",
  "lastName": "Beeblebrox",
  "email": "zaphod@beeblebrox.com",
  "password": "Betelgeuse123"
}
```

##### Response

```json
HTTP/1.1 200 OK
Content-Length: 173
Content-Type: application/json
Date: Sun, 18 May 2014 07:14:55 GMT

{
  "created": "2014-05-18T07:14:55.324614747Z",
  "email": "zaphod@beeblebrox.com",
  "firstName": "Zaphod",
  "id": 14,
  "lastName": "Beeblebrox",
  "updated": "2014-05-18T07:14:55.324614861Z"
}
```

#### Authenticate

##### Request

```json
{
  "email": "zaphod@beeblebrox.com",
  "password": "Betelgeuse123"
}
```

##### Response

```json
HTTP/1.1 200 OK
Content-Length: 129
Content-Type: application/json
Date: Sun, 18 May 2014 07:16:08 GMT

{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0MDA2NTY1NjgsImlkIjoxNH0.hmfpEmTidzQ5kEzJ3iZ_dMmhg-ohetW5rEyBx2Jl1TM"
}
```

`token` is a JSON Web Token (JWT). It should be included with every
authenticated request in the `Authorization` header.

Example of a valid `Authorization` header.

    Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0MDA2NTY1NjgsImlkIjoxNH0.hmfpEmTidzQ5kEzJ3iZ_dMmhg-ohetW5rEyBx2Jl1TM

### Journals

#### Create

    POST /journals

* Must be authenticated

##### Request

```json
{
  "name": "Personal"
}
```

##### Response

```json
HTTP/1.1 200 OK
Content-Length: 126
Content-Type: application/json
Date: Sun, 18 May 2014 07:26:43 GMT

{
  "created": "2014-05-18T07:26:43.204760935Z",
  "id": 14,
  "name": "Personal",
  "updated": "2014-05-18T07:26:43.204760993Z",
  "userID": 14
}
```

#### List All Journals

    GET /journals

###### Response

```json
HTTP/1.1 200 OK
Content-Length: 210
Content-Type: application/json
Date: Sun, 18 May 2014 07:31:16 GMT

[
  {
    "created": "2014-05-18T00:00:00Z",
    "id": 14,
    "name": "Personal",
    "updated": "2014-05-18T00:00:00Z",
    "userID": 14
  },
  {
    "created": "2014-05-18T00:00:00Z",
    "id": 15,
    "name": "Work",
    "updated": "2014-05-18T00:00:00Z",
    "userID": 14
  }
]
```

#### Get Journal

    GET /journals/:id

##### Response

```json
HTTP/1.1 200 OK
Content-Length: 106
Content-Type: application/json
Date: Sun, 18 May 2014 07:32:42 GMT

{
    "created": "2014-05-18T00:00:00Z",
    "id": 14,
    "name": "Personal",
    "updated": "2014-05-18T00:00:00Z",
    "userID": 14
}
```

#### Create New Entry in Journal

    POST /journals/:id/entries

##### Request

```json
{
  "rating": 4, // must be between 1 and 5
  "message": "I'm having a great time at this hackathon, but just ran into a bug."
}
```

##### Response

```json
HTTP/1.1 200 OK
Content-Length: 201
Content-Type: application/json
Date: Sun, 18 May 2014 07:35:06 GMT

{
  "created": "2014-05-18T07:35:06.517898461Z",
  "id": 10,
  "journalID": 14,
  "message": "I'm having a great time at this hackathon, but just ran into a bug",
  "rating": 4,
  "updated": "2014-05-18T07:35:06.517898522Z"
}
```
