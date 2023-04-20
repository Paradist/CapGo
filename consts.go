package capGo

var (
	CAPTCHA_TYPES = []string{

		"HCaptchaTask",
		"HCaptchaTaskProxyLess",
		"HCaptchaEnterpriseTask",
		"HCaptchaEnterpriseTaskProxyLess",
		"HCaptchaTurboTask",

		"FunCaptchaTask",
		"FunCaptchaTaskProxyLess",

		"GeeTestTask",
		"GeeTestTaskProxyLess",

		"ReCaptchaV2Task",
		"ReCaptchaV2TaskProxyLess",

		"ReCaptchaV2EnterpriseTaskProxyLess",
		"ReCaptchaV2EnterpriseTask",

		"ReCaptchaV3Task",
		"ReCaptchaV3TaskProxyLess",

		"ReCaptchaV3EnterpriseTask",
		"ReCaptchaV3EnterpriseTaskProxyLess",

		"MtCaptchaTask",
		"MtCaptchaTaskProxyLess",

		"DataDomeSliderTask",

		"AntiCloudflareTask",

		"AntiKasadaTask",

		"AntiAkamaiBMPTask",

		"ImageToTextTask",

		"HCaptchaClassification",

		"FunCaptchaClassification",

		"AwsWafClassification",
	}
)

const (
	STATUS_READY    = "ready"
	CREATE_TASK_PATH = "/createTask"
	GET_TASK_PATH   = "/getTaskResult"
	BALANCE_PATH     = "/getBalance"
	TASK_TIMEOUT    = 45
	API_URL			= "https://api.capsolver.com"
)