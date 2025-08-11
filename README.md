# Go JWT Authentication with Gin & MongoDB

A backend authentication service built using Golang, Gin Gonic framework, MongoDB, and JWT for secure API access.

Tested with Postman to demonstrate authentication, authorization, and role-based access control.

## Features
- User Signup & Login with hashed passwords (bcrypt).
- JWT Access & Refresh Tokens for secure authentication.
- Role-based Access Control (e.g., ADMIN vs USER).
- Protected Routes using JWT middleware.
- MongoDB as database for user storage.
- Environment variables for secure configuration.

```bash
.
├── main.go                  # Entry point
├── .env                     # Environment variables
├── controllers/             # Business logic
├── database/                # MongoDB connection
├── helpers/                 # Token generation & role helpers
├── middleware/              # JWT authentication middleware
├── models/                  # User model schema
├── routes/                  # API route definitions
└── images/                  # Postman screenshots
```

## Setup & Installation

1. Clone the repository
```bash
git clone https://github.com/0xk4n3ki/golang-jwt.git
```

2. Install dependencies
```bash
go mod tidy
```

3. Configure environment variables
```
PORT=9000
MONGODB_URL=mongodb://localhost:27017/go-auth
SECRET_KEY=<your_secret_key_here>
```

4. Run MongoDB
```bash
sudo systemctl start mongod
```

5. Start the server
```bash
go run main.go
```

The server will start at: 
```
http://localhost:9000
```

## API Endpoints

### Authentication

| Method| Endpoint| Description|
| :------------ | :--------- | ------: |
| POST | /users/signup| Register a new user |
| POST | /users/login | Login and get token |

### User Management (Protected)

| Method| Endpoint| Description| Role Required |
| :------------ | :--------- | :------ | ------: |
| GET | /users| List all users (paginated) | ADMIN |
| GET | /users/:user_id | Get user by ID | USER/ADMIN |

### Sample APIs (Protected)

| Method| Endpoint| Description|
| :------------ | :--------- | ------: |
| GET | /api-1	| Example protected endpoint |
| GET | /api-2 | Example protected endpoint |


## How It Works (Technical Refresher)

### 1. Server Initialization (main.go)
- Load environment variables from .env using joho/godotenv.
```go
godotenv.Load(".env")
port := os.Getenv("PORT")
```
- Initialize a new Gin HTTP engine:
```go
router := gin.New()
router.Use(gin.Logger())
```
- Register public routes (routes.AuthRoutes) for /users/signup and /users/login.
- Register protected routes (routes.UserRoutes) which are wrapped with middleware.Authenticate() for JWT validation.
- Start HTTP server:
```go
router.Run(":" + port)
```

### 2. Database Connection (database/databaseConnection.go)
- Reads MONGODB_URL from .env.
- Creates a new MongoDB client using the official MongoDB Go Driver:
```go
client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))
```
- Connects with a 10-second timeout using context.WithTimeout.
- Stores the *mongo.Client instance in a global variable for reuse across controllers:
```go
var Client *mongo.Client = DBinstance()
```
- Opens specific collections via OpenCollection(client, "user").

### 3. Request Handling in Gin
- Gin routes map HTTP methods and paths to controller handler functions.

    Example:
```go
incomingRoutes.POST("users/signup", controller.Signup())
incomingRoutes.GET("/users", controller.GetUsers())
```
- When a request hits a route:
    - Gin parses JSON body into Go structs using ctx.BindJSON(&model).
    - The controller processes input and interacts with MongoDB via userCollection.
    - A JSON response is sent back with ctx.JSON(statusCode, payload).

### 4. JWT Authentication Flow
- Signup (controllers.Signup)
    - Parse incoming JSON into models.User.
    - Validate fields using go-playground/validator.
    - Hash password with bcrypt.GenerateFromPassword.
    - Populate fields: _id, user_id, timestamps.
    - Generate access & refresh tokens using helpers.GenerateAllTokens().
    - Insert the document into MongoDB:
    ```go
    userCollection.InsertOne(c, user)
    ```

- Login (controllers.Login)
    - Find user by email from MongoDB:
    ```go
    userCollection.FindOne(c, bson.M{"email": user.Email})
    ```
    - Compare password using bcrypt.CompareHashAndPassword.
    - Generate new tokens with GenerateAllTokens() and update them in DB via UpdateAllTokens().
    - Return updated user document in the response.

- Middleware Authentication (middleware.Authenticate())
    - Extract token from HTTP request header:
    ```go
    clientToken := ctx.Request.Header.Get("token")
    ```
    - Validate token using helpers.ValidateToken() which:
    - Parses token using jwt.ParseWithClaims.
    - Verifies signature with SECRET_KEY.
    - Checks ExpiresAt against current time.
    - If valid, store claims in ctx.Set() for access in downstream handlers.

### 5. Role-based Access (helpers/authHelper.go)
- CheckUserType():
    - Compares user_type from JWT claims with the required role (ADMIN or USER).
- MatchUserTypeToUid():
    - Ensures a USER can only access their own data by matching uid from claims with route parameter.
- Used in controllers like GetUsers() and GetUser() to enforce authorization rules.

## Screenshots (Postman Tests)

1. api-1 without token (error)

<img alt="api-1-error" src="/images/api-1-error.png">

2. Signup

<img alt="signup" src="/images/signup.png">

3. Login

<img alt="login" src="/images/login.png">

4. Get Users (ADMIN)

<img alt="get users" src="/images/users.png">

5. api-1 with token (Success)

<img alt="api-1-success" src="/images/api-1-success.png">

## Credits

This project is based on the tutorial series by Akhil Sharma: 

[https://youtube.com/playlist?list=PL5dTjWUk_cPY7Q2VTnMbbl8n-H4YDI5wF&si=bZZ_o5_rszljKIJY](https://youtube.com/playlist?list=PL5dTjWUk_cPY7Q2VTnMbbl8n-H4YDI5wF&si=bZZ_o5_rszljKIJY)