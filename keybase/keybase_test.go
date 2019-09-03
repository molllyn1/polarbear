package keybase

import (
	"github.com/magiconair/properties/assert"
	assert2 "github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

var testKeyBase = NewDefaultKeyBase("./tmp")

func TestDefaultKeyBase(t *testing.T) {

	name := "default"
	password := "12345678"
	newPassword := "11223344"
	bip39Passphrase := "11111111"
	var mnemonic string
	var account uint32 = 0
	var index uint32 = 0

	mnemonic = testKeyBase.CreateKey(name, password, bip39Passphrase, account, index)
	info := strings.Split(mnemonic, "+")
	assert.Equal(t, len(info), 2)
	mnemonic = info[1]
	assert2.NotNil(t, mnemonic)

	res := testKeyBase.DeleteKey(name, password)
	assert.Equal(t, res, "")

	res = testKeyBase.RecoverKey(name, mnemonic, password, bip39Passphrase, account, index)
	assert.Equal(t, res, info[0])

	//keys := testKeyBase.ListKeys()

	res = testKeyBase.ResetPassword(name, password, newPassword)
	assert.Equal(t, res, "")

	_ = os.RemoveAll("./tmp")
}