package Pusher

import (
	"testing"
)

func TestMD5(t *testing.T) {

	s := `{"name":"foo","channels":["project-3"],"data":"{\"some\":\"data\"}"}`
	expected := "ec365a775a4cd0599faeb73354201b6f"
	sig := &Signature{
		"key",
	}
	res := sig.md5([]byte(s))
	if res != expected {
		t.Errorf("md5 not equal %s %s", res, expected)
	}
}
