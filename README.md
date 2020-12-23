## `go-gin-starter`

![build](https://github.com/udaya2899/go-gin-starter/workflows/Go/badge.svg) [![Go Reference](https://pkg.go.dev/badge/github.com/udaya2899/go-gin-starter.svg)](https://pkg.go.dev/github.com/udaya2899/go-gin-starter) [![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)

[![license](https://img.shields.io/github/license/udaya2899/go-gin-starter)](https://github.com/udaya2899/go-gin-starter/blob/master/LICENSE)

[![stars](https://img.shields.io/github/stars/udaya2899/go-gin-starter?style=social)](https://github.com/udaya2899/go-gin-starter)

An opinionated starter for Go Backend projects using:
* `gin-gonic/gin` as the REST framework
* `logrus` for logging
* `viper` for configs
* `Docker` for containerization
  
To be added:
* `sqlc` for type-safe SQL Go code generation
* `golang-migrate/migrate` for migration
* `jwt` authentication
* `casbin` authorization
* `prometheus` monitoring

### Docker support

Run `make docker-build-run` to build the command and run the container. Note that the port is set as 8080 by default.


Verify by checking `localhost:8080/ping`. You should receive a `pong`

### Contributing

Feel free to raise a PR with one of the features from "To be added" section.

If you want a new feature PR, raise an issue to discuss about it.
