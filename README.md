### API Server
##### API server with JWT authentication(RSA) built with GO 
#### Installation Instructions

* Edit config/local.yaml (*Change dbname or you can leave default* )
* `sh run.sh` from root directory (*Generates RSA keys and sqlite3 folder*)
* `go run main.go` run server from root directory 
* Test URL http://localhost:8080/ping

##### API server will run at: http://localhost:8080

##### API Information


| Endpoint | Method | JSON Format | Note
| --- | --- | --- | --- |
| `/register` | POST | ```{ "email": "johndoe@gmail.com","name" : "John Doe","gender" : "Male","country" : "US","age" : 26,"password": "password"}``` | *Email should be unique*
| `/login` | POST | ```{"email" : "johndoe@gmail.com","password" : "password"}```| Gives Access-Token in header after valid credentials
| `/user` | GET | **_Not required_** | This requires authentication, pass access-token value provided after login as bearer
