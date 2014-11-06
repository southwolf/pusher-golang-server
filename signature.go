package Pusher

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Signature struct {
	auth_key       string
	auth_secret    string
	auth_signature string
	auth_timestamp int
	auth_version   string
	body           []byte
	request_path   string
}

type ParamItem struct {
	key, value string
}

func (s *Signature) md5() string {
	hash := md5.New()
	hash.Write(s.body)
	return hex.EncodeToString(hash.Sum(nil))
}

func (s *Signature) sha256(content []byte) string {
	mac := hmac.New(sha256.New, []byte(s.auth_secret))
	mac.Write(content)
	return hex.EncodeToString(mac.Sum(nil))
}

func (s *Signature) signature_path() string {
	format := "POST\n%s\nauth_key=%s&auth_timestamp=%d&auth_version=%s&body_md5=%s"
	return fmt.Sprintf(format, s.request_path, s.auth_key, s.auth_timestamp, s.auth_version, s.md5())
}

func (s *Signature) sign() string {
	return s.sha256([]byte(s.signature_path()))
}

func (s *Signature) path() string {
	format := "%s?auth_key=%s&auth_timestamp=%d&auth_version=%s&body_md5=%s&auth_signature=%s"
	return fmt.Sprintf(format, s.request_path, s.auth_key, s.auth_timestamp, s.auth_version, s.md5(), s.sign())
}
