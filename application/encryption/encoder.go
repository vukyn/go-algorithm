package encryption

import "time"

/*
	Rules for encoding:
	- Current timestamp
	- Execution time
	- Secret question
	- Issuer token
	- Key for encode
*/

var secretQuestion = map[int]string{
	1: "Concurrent",
	2: "Golang",
}

type Encoder struct {
	Key           string
	Issuer        string
	EncryptedCode string
}

func (e *Encoder) Encode(content string, issuer string) string {
	encryptedCode := make([]byte, 10)
	now := time.Now().Unix()
	encryptedCode[0] = byte(now)
	return ""
}
