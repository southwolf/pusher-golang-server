package Pusher

import (
	"testing"
)

func TestMD5(t *testing.T) {
	body_content := `{"name":"foo","channels":["project-3"],"data":"{\"some\":\"data\"}"}`

	sig := &Signature{
		"278d425bdf160c739803",
		"7ad3773142a6692b25b8",
		"",
		1353088179,
		"1.0",
		[]byte(body_content),
		"/apps/3/events",
	}
	res := sig.md5()
	expected := "ec365a775a4cd0599faeb73354201b6f"

	if res != expected {
		t.Errorf("md5 not equal %s %s", res, expected)
	}
}

func TestSha256(t *testing.T) {
	body_content := `{"name":"foo","channels":["project-3"],"data":"{\"some\":\"data\"}"}`

	sig := &Signature{
		"278d425bdf160c739803",
		"7ad3773142a6692b25b8",
		"",
		1353088179,
		"1.0",
		[]byte(body_content),
		"/apps/3/events",
	}

	s := "POST\n/apps/3/events\nauth_key=278d425bdf160c739803&auth_timestamp=1353088179&auth_version=1.0&body_md5=ec365a775a4cd0599faeb73354201b6f"
	res := sig.sha256([]byte(s))
	expected := "da454824c97ba181a32ccc17a72625ba02771f50b50e1e7430e47a1f3f457e6c"

	if res != expected {
		t.Errorf("sha256 not equal \n%s \n%s", res, expected)
	}
}

func TestSignaturePath(t *testing.T) {
	body_content := `{"name":"foo","channels":["project-3"],"data":"{\"some\":\"data\"}"}`

	sig := &Signature{
		"278d425bdf160c739803",
		"7ad3773142a6692b25b8",
		"",
		1353088179,
		"1.0",
		[]byte(body_content),
		"/apps/3/events",
	}

	res := sig.signature_path()
	expected := "POST\n/apps/3/events\nauth_key=278d425bdf160c739803&auth_timestamp=1353088179&auth_version=1.0&body_md5=ec365a775a4cd0599faeb73354201b6f"

	if res != expected {
		t.Errorf("signarte path not equal \n%s \n%s", res, expected)
	}
}

func TestSign(t *testing.T) {
	body_content := `{"name":"foo","channels":["project-3"],"data":"{\"some\":\"data\"}"}`

	sig := &Signature{
		"278d425bdf160c739803",
		"7ad3773142a6692b25b8",
		"",
		1353088179,
		"1.0",
		[]byte(body_content),
		"/apps/3/events",
	}

	res := sig.sign()
	expected := "da454824c97ba181a32ccc17a72625ba02771f50b50e1e7430e47a1f3f457e6c"

	if res != expected {
		t.Errorf("sign error \n%s\n%s", res, expected)
	}
}

func TestPath(t *testing.T) {
	body_content := `{"name":"foo","channels":["project-3"],"data":"{\"some\":\"data\"}"}`

	sig := &Signature{
		"278d425bdf160c739803",
		"7ad3773142a6692b25b8",
		"",
		1353088179,
		"1.0",
		[]byte(body_content),
		"/apps/3/events",
	}

	res := sig.path()
	expected := "/apps/3/events?auth_key=278d425bdf160c739803&auth_timestamp=1353088179&auth_version=1.0&body_md5=ec365a775a4cd0599faeb73354201b6f&auth_signature=da454824c97ba181a32ccc17a72625ba02771f50b50e1e7430e47a1f3f457e6c"

	if res != expected {
		t.Errorf("path error \n%s\n%s", res, expected)
	}

}
