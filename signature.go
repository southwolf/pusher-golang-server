package Pusher

import (
	"crypto/md5"
	"encoding/hex"
)

type Signature struct {
	key string
}

type ParamItem struct {
	key, value string
}

func (s *Signature) md5(content []byte) string {
	hash := md5.New()
	hash.Write(content)
	return hex.EncodeToString(hash.Sum(nil))
}
