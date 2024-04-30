package transaction

import (
	"strings"
	"testing"
)

func TestDebit(t *testing.T) {
	t.Run("it should fail with insufficient balance error", func(t *testing.T) {
		err := Debit(Transaction{
			Amount:     1,
			Account_id: "hello",
		})
		if err == nil {
			t.Log("Err should not be nil")
			t.FailNow()
		}
		if !strings.Contains(strings.ToLower(err.Error()), "insufficient balance") {
			t.Log("Error message should contain insufficient balance")
			t.FailNow()
		}
	})
}
