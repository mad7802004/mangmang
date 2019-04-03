package e

const (
	SUCCESS                 = 0      //成功
	FAIL                    = -1     // 失败
	InvalidParameter        = 400001 // 参数无效
	InvalidAuthorization    = 40100  //授权无效,请登录
	VerificationCodeError   = 400101 //验证码错误
	MobileNumberError       = 400102 // 电话号码错误
	FrequentOperation       = 400103 // 请勿频繁操作
	InconsistentPassword    = 40104  // 密码不一致
	PhoneNumberIsRegistered = 40105  //手机号码被注册
	AccountOrPassWordErr    = 40106  //账户或密码错误
)

type Msg struct {
	M map[string]interface{}
}

func NewMsg(lang string) *Msg {
	m := &Msg{make(map[string]interface{}, 0)}
	if lang == "" {
		m.AddField("lang", "zh_Cn")
	} else {
		m.AddField("lang", lang)
	}
	m.AddField("code", SUCCESS)
	m.AddField("msg", GetMsg(SUCCESS, "zh_Cn"))
	m.AddField("data", nil)
	return m
}

func (m *Msg) AddField(key string, e interface{}) {
	m.M[key] = e
}

// 修改msg内容
func (m *Msg) Update(code int, data interface{}) {
	m.AddField("code", code)
	m.AddField("data", data)
	if lang, ok := m.M["lang"]; ok {
		m.AddField("msg", GetMsg(code, lang.(string)))
	} else {
		m.AddField("msg", GetMsg(code, "zh_Cn"))
	}
}
