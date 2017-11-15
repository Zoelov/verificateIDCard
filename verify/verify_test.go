package verify

import (
	"fmt"
	"testing"
)

func TestVerifyID(t *testing.T) {
	id := ""

	ok, err := VerificateID(id)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ok)
}
