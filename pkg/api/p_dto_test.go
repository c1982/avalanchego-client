package api

import (
	"testing"
)

func TestPDTOKey(t *testing.T) {

	user, txID, changeAddr := "", "", ""

	p := P{}
	p["user"] = "user"
	p["txID"] = "txID"
	p["changeAddr"] = "changeAddr"

	err := p.OutStr("user", &user).OutStr("txID", &txID).OutStr("changeAddr", &changeAddr).Error()
	if err != nil {
		t.Error(err)
	}

	if user != "user" {
		t.Errorf("user value error. expect: user, got: %s", user)
	}

	if txID != "txID" {
		t.Errorf("txID value error. expect: txID, got: %s", txID)
	}

	if changeAddr != "changeAddr" {
		t.Errorf("changeAddr value error. expect: changeAddr, got: %s", changeAddr)
	}

	t.Run("key_not_found_error", func(t *testing.T) {
		var muser string
		err := p.OutStr("muser", &muser).Error()
		if err == nil {
			t.Error(err)
		}
	})
}
