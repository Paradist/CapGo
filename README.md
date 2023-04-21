# (WIP) CapGo
Golang package for easy and fast integration with [CapSolver API](https://dashboard.capsolver.com/passport/register?inviteCode=GuA_6LuFEqnn)
- An API key it's **required**. [**Get here.**](https://dashboard.capsolver.com/passport/register?inviteCode=GuA_6LuFEqnn)
- [Documentation](https://docs.capsolver.com/guide/getting-started.html), information on what to write in `task`

## Supported CAPTCHA types:
- HCaptcha
- FunCaptcha
- Geetest
- ReCaptchaV2
- ReCaptchav3
- MtCaptcha
- Datadom
- Cloudflare
- Kasada
- Akamai BMP


## Installation

```sh
go get github.com/paradist/capgo
```

## Usage
### Balance
```go
  client := capGo.Client("APIKEY")
  bal, err := client.Balance()
  if (err != nil) {
    log.Println(err)
  }
  log.Println(bal)
```

### Solve captcha
```go
 client := capGo.Client("APIKEY")
 resp, err := client.Solve(map[string]any{
 		"type":       "ReCaptchaV2taskProxyLess",
		"websiteURL": "https://www.google.com/recaptcha/api2/demo",
		"websiteKey": "6Le-wvkSAAAAAPBMRTvw0Q4Muexq9bi0DJwx_mJ-",
  })
 if (err != nil) {
    log.Println(err)
  }
 log.Println(resp.Solution.GRecaptchaResponse)
```
```go
client := capGo.Client("APIKEY")
resp, err := client.Solve(map[string]any{
	"type":       "HCaptchaEnterpriseTaskProxyLess",
	"websiteURL": "https://hcaptcha.com/",
	"websiteKey": "00000000-0000-0000-0000-000000000000",
})
if err != nil {
	log.Fatal(err)
	return
}
fmt.Println(resp)
```

### Recognition
```go
 client := capGo.Client("APIKEY")
 b, err := os.ReadFile("queue-it.jpg")
 if err != nil {
    panic(err)
 }
 resp, err := client.Solve(map[string]any{
		"type": "ImageToTextTask",
		"module": "queueit",
		"body":   base64.StdEncoding.EncodeToString(b),
  })
 if (err != nil) {
    log.Println(err)
  }
 log.Println(resp.Solution.Text)
```

