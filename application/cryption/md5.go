package cryption

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
)

func EncriptMD5(str string) string {
	t := make(map[string]string)
	t["key"] = "MyKey"
	res, _ := json.Marshal(t)
	h := md5.New()
	io.WriteString(h, string(res))
	fmt.Printf("%x", h.Sum(nil))
	return ""
}
