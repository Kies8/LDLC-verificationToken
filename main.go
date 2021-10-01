package main

import (
	"fmt"
	"strings"
	"time"
)

func VerificationToken() string {
	epoch := int64(time.Now().UnixNano() / 1e6)
	context := map[string]interface{}{"a": epoch, "b": 0, "c": 1, "d": 2, "e": 0}
	epochB16 := fmt.Sprintf("0%x", epoch)
	o, u := "", 0

	for _, char := range strings.Replace(fmt.Sprintf("%v", context), " ", "", -1) {
		s := int(epochB16[u]) - '0'
		u++
		if u >= len(epochB16) {
			u = 0
		}
		xored := byte(char) ^ byte(s)
		o += fmt.Sprintf("%02x", xored)
	}

	return fmt.Sprintf("%s-%s", epochB16, o)
}

func main() {
	token := VerificationToken()
	fmt.Println(token)
}
