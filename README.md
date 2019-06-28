# QIWI personal api

QIWI Wallet API makes it easy to automate getting info on your account’s state in QIWI Wallet service and making financial operations.

## Installation

`go get github.com/neqin/qiwi-personal-api`

## Examples

### Create new API

```go
package main

import (
	"log"

	qiwi "github.com/neqin/qiwi-personal-api"
)

func main() {
    qiwi.NewQiwiPersonalApi("5fa740ea1daf00665aa312...")
}
```

### Get Identification
```go
package main

import (
	"log"

	qiwi "github.com/neqin/qiwi-personal-api"
)

func main() {
    q := qiwi.NewQiwiPersonalApi("5fa740ea1daf00665aa312...")
    res, err := q.GetIdentification("79264810000")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(res)
}
```
