package cryption

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
)

func EncriptMD5() {
	data := make(map[string]interface{})
	item := make(map[string]interface{})
	item["sku"] = 100160010
	item["qty"] = 2
	data["items"] = []interface{}{item}
	data["box_type"] = "20x8x8"
	fmt.Println(data)
	res, _ := json.Marshal(data)
	res = append(res, []byte("#test.hasaki.vn")...)
	h := md5.New()
	io.WriteString(h, string(res))
	fmt.Printf("%x\n", h.Sum(nil))
	fmt.Printf("%v\n", string(res))
	//{"box_type":"20x8x8","items":[{"sku":100160010,"qty":2}]}#test.hasaki.vn
	//{"box_type":"20x8x8","items":[{"qty":2,"sku":100160010}]}#test.hasaki.vn
}
