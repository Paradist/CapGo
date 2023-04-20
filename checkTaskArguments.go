package capGo

import (
	"fmt"
	"strings"
)

func (c capGoStruct) CheckTaskArguments(args map[string]any) (error) {
	captchaType, ok := args["type"].(string)
	if !(ok) {
		return argError(captchaType, "type")
	}
	captchaType = strings.ToLower(captchaType)
	validType := false
	for _, v := range CAPTCHA_TYPES {
		if strings.ToLower(v) == captchaType {
			validType = true
			break
		}
	}
	if !(validType) {
		return fmt.Errorf("unsupported type `%s`! %s", captchaType, c.FormattedCaptchaTypes())
	}
	if strings.Contains(captchaType, "recaptcha") {
		if captchaType != "recaptchaclassification" {
			return checkNormalCaptcha(args)
		} else {
			return nil
		}
	}
	if strings.Contains(captchaType, "hcaptcha") {
		if captchaType == "hcaptchaclassification" {
			if _, ok = args["queries"]; !ok {
				return argError(captchaType, "queries")
			}
			if _, ok = args["question"]; !ok {
				return argError(captchaType, "question")
			}
			return nil
		} else {
			return checkNormalCaptcha(args)
		}
	}
	if strings.Contains(captchaType, "funcaptcha") {
		if captchaType == "funcaptchaclassification" {
			if _, ok = args["images"]; !ok {
				return argError(captchaType, "images")
			}
			return nil
		} else {
			if _, ok := args["websiteURL"]; !ok {
				return argError(captchaType, "websiteURL")
			}
			if _, ok := args["websitePublicKey"]; !ok {
				return argError(captchaType, "websitePublicKey")
			}
			return nil
		}
	}

	if strings.Contains(captchaType, "mtcaptcha") {
		return checkNormalCaptcha(args)
	}

	if strings.Contains(captchaType, "geetesttask") {
		if _, ok = args["gt"]; !ok {
			return argError(captchaType, "gt")
		}
		if _, ok = args["challenge"]; !ok {
			return argError(captchaType, "challenge")
		}
		return nil
	}
	if strings.Contains(captchaType, "datadom") {
		if _, ok = args["proxy"]; !ok {
			return argError(captchaType, "proxy")
		}
		if _, ok = args["useragent"]; !ok {
			return argError(captchaType, "userAgent")
		}
		return nil
	}
	if strings.Contains(captchaType, "anticloudflare") {
		if _, ok = args["metadata"]; !ok {
			return argError(captchaType, "metadata")
		}
		if _, ok = args["proxy"]; !ok {
			return argError(captchaType, "proxy")
		}
		return nil
	}
	if strings.Contains(captchaType, "antikasada") {
		if _, ok = args["pageURL"]; !ok {
			return argError(captchaType, "pageURL")
		}
		if _, ok = args["proxy"]; !ok {
			return argError(captchaType, "proxy")
		}
		return nil
	}
	if strings.Contains(captchaType, "antiakamaibmp") {
		if _, ok = args["packageName"]; !ok {
			return argError(captchaType, "packageName")
		}
	}
	return nil
}

func checkNormalCaptcha(params map[string]any) (error) {
	captchaType, _ := params["type"].(string)
	if _, ok := params["websiteKey"]; !ok {
		return argError(captchaType, "websiteKey")
	}
	if _, ok := params["websiteURL"]; !ok {
		return argError(captchaType, "websiteURL")
	}
	return nil
}
func argError(captchaType string, arg string) (error) {
	return fmt.Errorf("%s need `%s` argument", captchaType, arg)
}