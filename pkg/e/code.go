package e

const (
	SUCCESS          = 0     //成功
	FAIL             = -1    // 失败
	InvalidAuthorization = 40001 //授权无效,请登录
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
