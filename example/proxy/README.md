# proxy example

This example guides you on how to generate and use proxy code with gen-go-proxy.

## Usage

### implement interface and middleware

```go
package service

// implement target interface
type Foo interface {
  // use @{annotation name} comment if you want to generate proxy code
  // @proxy
  Logic(needEmitErr bool) (string, error)

  // also support multiple annotation
  // proxy middleware runs in order of annotation
  // @custom1
  // @custom2
  Foo() int
}
```

```go
// implement middleware
func Wrapped(next func(c context.Context) error) func(context.Context) error {
  return func(c context.Context) error {
    fmt.Println("[Wrapped] before")

    // run next middleware or target logic
    err := next(c)
    if err != nil {
      fmt.Printf("[Wrapped] err occurred. err : %s\n", err)
    }

    fmt.Println("[Wrapped] after")
    return err
  }
}

func Before(next func(c context.Context) error) func(context.Context) error {
  return func(c context.Context) error {
    fmt.Println("[Before] before")

    // run next middleware or target logic
    return next(c)
  }
}

func After(next func(c context.Context) error) func(context.Context) error {
  return func(c context.Context) error {
    // run next middleware or target logic
    err := next(c)
    if err != nil {
      fmt.Printf("[After] err occurred. err : %s\n", err)
    }

    fmt.Println("[After] after")
    return err
  }
}
```

### generate proxy code

```bash
$ gen-go-proxy -t example/proxy/service
```

### use generated proxy code

```go
// create instence and regist middleware
func main() {
  target := service.NewFoo()

  // middleware by annotation
  // key: annotation name
  // value: middleware list
  // can use middleware helper type
  // or raw type map[string][]func(func(context.Context) error) func(context.Context) error
  //
  //  m := map[string][]func(func(context.Context) error) func(context.Context) error{
  //    "proxy":   {Wrapped, Before, After},
  //    "custom1": {Wrapped},
  //    "custom2": {Before, After},
  //  }
  //
  m := service.FooProxyMiddlewareByAnnotation{
    "proxy":   {Wrapped, Before, After},
    "custom1": {Wrapped},
    "custom2": {Before, After},
  }


  // if use middleware helper type, should call helper.To() when create proxy
  proxy := service.NewFooProxy(target, m.To())

  if val, err := proxy.Logic(false); err != nil {
    fmt.Println("err: ", err)
  } else {
    fmt.Println("val: ", val)
  }

  fmt.Println()

  if val, err := proxy.Logic(true); err != nil {
    fmt.Println("err: ", err)
  } else {
    fmt.Println("val: ", val)
  }

  fmt.Println()

  value := proxy.Foo()
  fmt.Println("value: ", value)
}
```

```bash
$ go run example/proxy/main.go
[Wrapped] before
[Before] before
[Foo] logic
[After] after
[Wrapped] after
val:  foo logic

[Wrapped] before
[Before] before
[Foo] logic
[After] err occurred. err : emit error
[After] after
[Wrapped] err occurred. err : emit error
[Wrapped] after
err:  emit error

[Wrapped] before
[Before] before
[Foo] foo
[After] after
[Wrapped] after
value:  1
```
