package capGo

import (
	"fmt"
)

func (c capGoStruct) FormattedCaptchaTypes() (string) {
	resp := "The current list of captcha types:"

	for i, v := range CAPTCHA_TYPES {
		resp += fmt.Sprintf("\n%d) %s", i, v)
	}
	return resp
}