# Golang Error Lines and Logging

### Install package
```
go get github.com/yogameleniawan/golang-error-lines
```

## Usage 

#### Import package golang error lines
```go
import errors "github.com/yogameleniawan/golang-error-lines"
```

### Log File
![image](https://github.com/yogameleniawan/golang-error-lines/assets/64576201/63201da5-4f36-45d5-91c7-c29ab4437a4c)
![image](https://github.com/yogameleniawan/golang-error-lines/assets/64576201/faad0ceb-ec8a-4151-b63f-61eee0196748)

##### Error Log
```go
errors.Error("Error Brooo!")
```

##### Info Log
```go
errors.Info("Info Brooo!")
```

##### Warning Log
```go
errors.Warning("Warning Brooo!")
```

#### Error Terminal
![image](https://github.com/yogameleniawan/golang-error-lines/assets/64576201/5fc2205c-1355-44e6-82c7-68c88d1c1182)

##### Error Message Terminal
```go
errors.ErrorLines(err error)
```

## Changelog

Please see [CHANGELOG](CHANGELOG.md) for more information what has changed recently.

## Credits

- [Yoga Meleniawan Pamungkas](https://github.com/yogameleniawan)
- [All Contributors](../../contributors)
