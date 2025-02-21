# Go: contexttime

## USAGE

```go
package main

import (
 "context"
 "fmt"
 "time"

 "github.com/hidori/go-contexttime"
)

func main() {
 production()
 test()
 production()
 test()
}

func production() {
 ctx := context.Background()
 ctx = contexttime.SetDefaultTime(ctx)

 fmt.Println(contexttime.GetTime(ctx).Now())
}

var _ = contexttime.Time((*MockTime)(nil))

type MockTime struct{}

func (t *MockTime) Now() time.Time {
 return time.Date(2025, 2, 21, 0, 0, 0, 0, time.UTC)
}

func test() {
 ctx := context.Background()
 ctx = contexttime.SetTime(ctx, &MockTime{})

 fmt.Println(contexttime.GetTime(ctx).Now())
}
```

OUTPUT:

```text
2025-02-21 19:17:26.403650012 +0900 JST m=+0.000040009
2025-02-21 00:00:00 +0000 UTC
2025-02-21 19:17:26.403748377 +0900 JST m=+0.000138383
2025-02-21 00:00:00 +0000 UTC
```
