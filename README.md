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
- Postgres
- PGAdmin
- docker
### Go mod init:
- go mod init github.com/luciormoraes/go-rest-api
- go mod init go-rest-api

## Day 24/100: Developing a Rest API with GIN
- Installing gin
- creating a gin server / first Endpoint(GET)
- creating packages: routes/ controllers/ models