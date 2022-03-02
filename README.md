# Developing a Rest API with GIN
## Tech stack:
- Go
- GIN
```bash
$ go get -u github.com/gin-gonic/gin
```
- GORM:
```bash
$ go get -u gorm.io/gorm
$ go get gorm.io/driver/postgres # postgres driver
```
- Validator: https://pkg.go.dev/gopkg.in/validator.v2
```bash
go get gopkg.in/validator.v2
```
- Postgres
- PGAdmin
- docker
    - Find Postgres IP:
    ```bash
    $ docker inspect CONTAINER_ID | grep IPAdd
    ```
### Go mod init:
- go mod init github.com/luciormoraes/go-rest-api
- go mod init go-rest-api

## Day 24/100: Developing a Rest API with GIN
- Installing gin
- creating a gin server / first Endpoint(GET)
- creating packages: routes/ controllers/ models

## Day 25/100: Developing a Rest API with GIN
- Installing GORM
- Change controller to create students / new endpoint
- using autoMigrate from gorm: Struct Student as gorm.Model

## Day 27/100: Developing a Rest API with GIN
- return all students from DB
- new endpoints: SearchStudentByID, DeleteStudent, EditStudent, SearchStudentByIRD

## Day 30/100: Developing a Rest API with GIN
- writing validations: installing Validator.v2