package iridium

import "testing"

var k = "^LSU!$aatnuF5f8u"

func TestAESEncryptWithCFB(t *testing.T) {
	s := "hello world"
	c, err := AESEncryptWithCFB(s, k)
	if err != nil {
		t.Log(err)
	}
	t.Log(c)
}

func TestAESDecryptWithCFB(t *testing.T) {
	c := "4390fc2b261a103fd662e5cf01b5af15827b1ff2627e77df67585f"
	s, err := AESDecryptWithCFB(c, k)
	if err != nil {
		t.Log(err)
	}
	t.Log(s)
}
