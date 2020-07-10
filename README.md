![build](https://github.com/luckylookas/spahandler/workflows/build/badge.svg?branch=master)
![codecov](https://codecov.io/gh/luckylookas/spahandler/branch/master/graph/badge.svg)
#spahandler

provides a simple notFoundHandler for serving an api alongside a single plage application


## api

```go
your_router.NotFoundHandler = spahandler.NewDefaultSpaHandler()
```

should do for most cases

### configuration

if the defaults do not suite your case, these are your options

```go
type SpaOptions struct {
	IgnorePrefix    string
	DefaultResource string
	ContentProvider ContentProvider
	Propagate http.Handler
}
```

| key             | default                       | use                                                                                                           |
|-----------------|-------------------------------|---------------------------------------------------------------------------------------------------------------|
| IgnorePrefix    | "/api"                        | setting this to "api" means that even if a file exists unter /api/index.html, the handler will never serve it |
| DefaultResource | "index.html"                  | the name of your spa root file                                                                                |
| ContentProvider | serving files from "./webapp" | this can be configured to eg. transparently serve from eg. a backing fileserver if need be                    |
| propagate       | default write statuscode 404) | in case not founds should be handled in a special way (eg. logged)        

for further details, see godoc