# structs

Whoo! Data Structures (Only Include Basic Operations)

# Installation
Make sure you have you set your GOPATH (https://github.com/golang/go/wiki/GOPATH).
Type the command below to install the package.

```bash
$ go get github.com/carljoshua/structs
```

# Usage

Import the package and then create the data structure you need. Simple as that.
The data structures that are included in this package are:

* Stack
* Queue
* Linked List
* Heap


```go
package main

import (
    "github.com/carljoshua/structs"
    "reflect"
)

func main(){
  s := structs.NewStack(reflect.Int, 4)
  s.Push(1)
  s.Push(2)
  s.Pop()
}
```
