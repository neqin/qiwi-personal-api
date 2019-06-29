# QIWI personal API

QIWI Wallet API makes it easy to automate getting info on your account’s state in QIWI Wallet service and making financial operations.

## Installation

`go get github.com/neqin/qiwi-personal-api`

## Examples

### Init and get identification info
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
    log.Println(res)
}
```

### list API method

- [x] Person’s Profile [GetPersonProfile] 
- [x] Payments History [GetPaymentsHistory] (simple)
- [x] Statistics on payments [GetPaymentsStats]
- [ ] Checkout 
- [x] QIWI Wallet Balances [GetWalletBalance]
- [ ] Commission rates
- [ ] Peer-to-Peer QIWI Wallet Transfer
- [ ] Wireless operator check
- [ ] Card transfer
- [ ] Card system check
- [ ] Wire transfer