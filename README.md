# Simple go dependency injection library

## How to inject

```go
svc := service.InitService()

// Inject the service
// make sure you pass pointer to the object
err := gi.Inject(&svc) 
if err != nil {
    slog.Error("failed to inject service", "error", err)
    return
}
```

## How to invoke
```go
// using generics to find the exact type of the object passed
// make sure you use pointer
// service.IService is service interface
svc, err := gi.Invoke[*service.IService]()
if err != nil {
    slog.Error("failed to invoke service", "error", err)
    return
}

// call your method here
(*svc).InterfaceMethod()
```

## Improvements to be done

- Allow multiple type injections of same type
- Support named injections