package transaction

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/winzcom/bank-t/account"
	_ "github.com/winzcom/bank-t/account"
	"github.com/winzcom/bank-t/db"
	"github.com/winzcom/bank-t/http"
)

type Transaction struct {
	Amount     float32
	Reference  string
	Account_id string
	Db         db.Storage
}

var AC account.Accounts

var st db.Storage

func init() {
}

func Debit(t Transaction) error {
	/// call third party and save to db
	b := new(bytes.Buffer)

	enc := json.NewEncoder(b)

	t.Db.Save(nil)

	enc.Encode(t)
	// check account to see if there is enough balance
	fmt.Println("acac ", AC)
	bal := AC[t.Account_id]

	if bal <= 0 {
		return errors.New("Insufficient Balance")
	}
	//AC[t.Account_id] -= bal
	st.Save(nil)
	http.PostTransactions(b, "debit")
	return nil
}

func Credit(t Transaction) error {
	var b = new(bytes.Buffer)
	enc := json.NewEncoder(b)
	enc.Encode(t)
	bal := AC[t.Account_id]
	bal += int(t.Amount)
	AC[t.Account_id] = bal
	http.PostTransactions(b, "credit")

	return nil
}
