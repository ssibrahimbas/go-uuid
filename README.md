<p align="center"><br><img src="https://avatars.githubusercontent.com/u/76786120?v=4" width="128" height="128" style="border-radius: 50px;" /></p>
<h3 align="center">Go-UUID</h3>
<p align="center">
  A uuid project with go. It includes nginx, fiber and postgresql.
</p>

### Installation

Docker must be installed for this project to work. To install Docker:

[click here](https://docs.docker.com/engine/install/)

and install docker compose:

[click here](https://docs.docker.com/compose/install/)

### Bootstrapping

use this command in your bash/terminal

```bash
docker-compose up
```

now send a get request to

```
http://localhost:3050/api/users/v1/all
```

### API

### Users

Operations related to the user are listed here.

### Create User

This endpoint should be used to create a user.

|Params|Type|
|------|----|
|**email**|*string*|
|**password**|*string*|

#### Simple request


```http request
POST http://localhost:3050/api/users/v1/create
Content-Type: application/json

{
    "email": "info@ssibrahimbas.com",
    "password": "12345"
}
```

#### Simple Response

```json
{
  "success": true,
  "message": "User created successfully",
  "data": {
    "id": "4359d6d9-b728-415d-a6ee-f28c0eaf8c1f",
    "email": "info@ssibrahimbas.com"
  }
}
```

### Get All Users

All users can be listed with this endpoint.

#### Simple Request

```http request
GET http://localhost:3050/api/users/v1/all
```

#### Simple Response

```json
{
  "success": true,
  "message": "Users listed successfully",
  "data": [
    {
      "id": "4359d6d9-b728-415d-a6ee-f28c0eaf8c1f",
      "email": "info@ssibrahimbas.com"
    },
    {
      "id": "6f0b2f5a-b416-40ba-b501-121a33b8de5e",
      "email": "info@istanbuljs.org"
    }
  ]
}
```

### Get User Detail

User detail can be listed with this endpoint.

|Params|Type|
|------|----|
|**id**|*uuid*|

#### Simple Request

```http request
GET http://localhost:3050/api/users/v1/detail/4359d6d9-b728-415d-a6ee-f28c0eaf8c1f
```

#### Simple Response

```json
{
  "success": true,
  "message": "User fetched successfully",
  "data": {
    "id": "e6d1dc7d-0ebc-4a11-ae27-18bb517d19ef",
    "email": "info@ssibrahimbas.com"
  }
}
```