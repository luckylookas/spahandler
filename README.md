![build](https://github.com/luckylookas/spahandler/workflows/build/badge.svg?branch=master)
![codecov](https://codecov.io/gh/luckylookas/spahandler/branch/master/graph/badge.svg)

# spahandler

provides a simple notFoundHandler for serving an api alongside a single plage application

### configuration

if the defaults do not suite your case, these are your options

```go
type SpaOptions struct {
	ContentRoot string
	FailureHandler  http.HandlerFunc
}
```

| key             | default                       | use                                                                                                           |
|-----------------|-------------------------------|---------------------------------------------------------------------------------------------------------------|
| ContentRoot | "./webapp"                  | the name of your static resources directory                                                                             |
| FailureHandler       | default writes statuscode 404) | if the handler cannot find the requested resource, it propagates to this "actual" NotFoundHandler

