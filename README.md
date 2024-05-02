# go-fileprobe

Probe util with file for golang

## Install

```shell
go get github.com/i14t/go-fileprobe
```

## Example

```go
package main

import (
 "fmt"
 "log"
 "time"

 "github.com/i14t/go-fileprobe"
)

func main() {
 fp := fileprobe.NewHandler()
 if err := fp.Create(); err != nil {
  log.Fatal(err)
 }
 fmt.Printf("File probe is exists: %v\n", fp.Exists())

 duration := 5 * time.Second
 for remaining := duration; remaining >= 0; remaining -= time.Second {
  fmt.Printf("Time remaining: %d seconds\n", remaining/time.Second)
  time.Sleep(time.Second)
 }

 if err := fp.Remove(); err != nil {
  log.Fatal(err)
 }
 fmt.Printf("File probe is exists: %v\n", fp.Exists())
}
```

## Reference

- [dollarsignteam/go-fileprobe][1]

## License

Licensed under the MIT License - see the [LICENSE][2] file for details.

[1]: https://github.com/dollarsignteam/go-fileprobe
[2]: https://github.com/i14t/go-fileprobe/blob/main/LICENSE