#Simple Go application

- CRUD operation on mongodb

##Routes:
- / (All methods)
- /createProfile (POST)
- /login (POST)
- /app/getProfile (GET)
- /app/updateProfile (PUT)
- /app/deletProfile (DELETE)

##Model:
- **User**
  - Name
  - Username
  - EmailId
  - Age
  - Password

##Config:
- **ServerConfig**
  - Host
  - Port
  - Read timeout
  - Write timeout
  - Idle timeout
- **MongoConfig**
  - Host
  - Port
  - username
  - password
  - DbName

#To run the server
```
cd cmd
go get -u
go run main.go
```
##Example APIs
###Create/Register User
```
POST /createProfile 
{
    "username": "sushant123",
    "password": "123467890",
    "age": 25,
    "email_id": "sushant@gmail.com",
    "name": "Sushant Kumar"
}
```
###Login
```
POST /login 
{
    "username": "sushant123",
    "password": "123467890"
}
```
### Get User Profile
```
GET /app/getProfile/
```
**Note: All the requests to /app/ path will be authenticated using cookie**

### Update the profile

```
PUT /app/updateProfile/
{
    "password": "12346",
    "age": 24
}
```
This will update the age and password for the user

### Delete Profile

```
DELETE /app/deleteProfile
```