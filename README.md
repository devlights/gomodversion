# gomodversion

Repository that test golang module versioning.

## Try

### no version specified (from v0.0.0 to v1.9.9)

```sh
$ go list -versions -m github.com/devlights/gomodversion
github.com/devlights/gomodversion v0.0.1 v0.1.0 v1.0.0 v1.0.1 v1.0.2 v1.1.0 v1.1.1 v1.1.2
```

```go
package main

import (
        "fmt"

        ver "github.com/devlights/gomodversion"
)

func main() {
        fmt.Println(ver.Hello()) // OUTPUT: HELLO
}
```

### v2

```sh
$ go list -versions -m github.com/devlights/gomodversion/v2
github.com/devlights/gomodversion/v2 v2.0.0 v2.0.1 v2.0.2
```

```go
package main

import (
        "fmt"

        ver "github.com/devlights/gomodversion/v2"
)

func main() {
        fmt.Println(ver.Hello()) // OUTPUT: --**HELLO**--
}
```

### v3

```sh
$ go list -versions -m github.com/devlights/gomodversion/v3
github.com/devlights/gomodversion/v3 v3.0.0 v3.0.1
```

```go
package main

import (
        "fmt"

        ver "github.com/devlights/gomodversion/v3"
)

func main() {
        fmt.Println(ver.Hello()) // OUTPUT: |--**HELLO WORLD**--|
}
```

### v4

```sh
$ go list -versions -m github.com/devlights/gomodversion/v4
github.com/devlights/gomodversion/v4 v4.0.0 v4.0.1
```

```go
package main

import (
        "fmt"

        ver "github.com/devlights/gomodversion/v4"
)

func main() {
        fmt.Println(ver.Hello()) // OUTPUT: ^^|--**HELLO WORLD**--|^^
}
```
