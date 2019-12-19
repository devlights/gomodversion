# gomodversion
Repository that test golang module versioning.

## try

### no version specified (from v0.0.0 to v1.9.9)

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