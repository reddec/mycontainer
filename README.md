# mycontainer


[![Go Reference](https://pkg.go.dev/badge/github.com/reddec/mycontainer.svg)](https://pkg.go.dev/github.com/reddec/mycontainer)


Detect self container ID using various methods such as: `cpuset` or `mountpoint`.

```go
package main

import (
	"fmt"

	"github.com/reddec/mycontainer"
)

func main() {
	containerID, err := mycontainer.ContainerID()
	if err != nil {
		panic(err)
	}

	fmt.Println("Container ID:", containerID)
}

```

If you have [ko](https://ko.build) installed, you may test it:

    $ docker run  $(ko build -L ./example)
    ....
    Container ID: 9b3d8a8c029bb30827f28fa09598693162cb58edb2389eb59e6c5309047643b1