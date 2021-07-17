# HTTP2

http2 is a implementation of HTTP/2 protocol for [fasthttp](https://github.com/valyala/fasthttp).

# Download

```bash
go get github.com/ip-rw/http2@v0.0.7
```

# Help

If you need any help setting up, contributing or understanding this repo, you can contact me on [gofiber's Discord](https://gofiber.io/discord).

# How to use the server?

The server can only be used if your server supports TLS.
Then, you can call [ConfigureServer](https://pkg.go.dev/github.com/ip-rw/http2@v0.0.3/fasthttp2#ConfigureServer).

```go
import (
	"github.com/valyala/fasthttp"
	"github.com/ip-rw/http2/fasthttp2"
)

func main() {
    s := &fasthttp.Server{
        Handler: yourHandler,
        Name:    "HTTP2 test",
    }

    fasthttp2.ConfigureServer(s)
    
    s.ListenAndServeTLS(...)
}
```

# How to use the client?

The HTTP/2 client only works with the HostClient.

```go
package main

import (
        "fmt"
        "log"

        "github.com/ip-rw/http2/fasthttp2"
        "github.com/valyala/fasthttp"
)

func main() {
        hc := &fasthttp.HostClient{
                Addr:  "api.binance.com:443",
                IsTLS: true,
        }

        if err := fasthttp2.ConfigureClient(hc); err != nil {
                log.Printf("%s doesn't support http/2\n", hc.Addr)
        }

        statusCode, body, err := hc.Get(nil, "https://api.binance.com/api/v3/time")
        if err != nil {
                log.Fatalln(err)
        }

        fmt.Printf("%d: %s\n", statusCode, body)
}
```
