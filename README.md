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
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0MDA2NTY1NjgsImlkIjoxNH0.hmfpEmTidzQ5kEzJ3iZ_dMmhg-ohetW5rEyBx2Jl1TM"
}
```

`token` is a JSON Web Token (JWT). It should be included with every
authenticated request in the `Authorization` header.

Example of a valid `Authorization` header.

    Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0MDA2NTY1NjgsImlkIjoxNH0.hmfpEmTidzQ5kEzJ3iZ_dMmhg-ohetW5rEyBx2Jl1TM
