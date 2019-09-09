package e

const (
	SUCCESS                  = 0      //成功
	FAIL                     = -1     // 失败
	NewFailed                = 400001 // 新建失败
	UpdateFailed             = 400002 // 更新失败
	FailedToDelete           = 400003 // 删除失败
	AcquisitionFailed        = 400004 // 获取失败
	InvalidParameter         = 400005 // 参数无效
	InvalidAuthorization     = 400100 //授权无效,请登录
	VerificationCodeError    = 400101 //验证码错误
	MobileNumberError        = 400102 // 电话号码错误
	FrequentOperation        = 400103 // 请勿频繁操作
	InconsistentPassword     = 400104 // 密码不一致
	PhoneNumberIsRegistered  = 400105 //手机号码被注册
	AccountOrPassWordErr     = 400106 //账户或密码错误
	OldPasswordError         = 400107 // 旧密码错误
	NoResourcesFound         = 400108 // 没有找到资源
	AccountDoesNotExist      = 400109 //账户不存在
	BusinessCardDoesNotExist = 400110 //名片不存在
	ImageByteIsTooLarge      = 400111 // 图片过大
	ImageFormatIsIncorrect   = 400112 // 图片格式不正确
	RoleDoesNotExist         = 400113 // 角色不存在
	ProjectDoesNotExist      = 400114 // 项目不存在
	ProjectUserExist         = 400115 //用户已存在该项目中，请勿重复添加
	ProjectUserDoesNotExist  = 400116 // 项目用户不存在
	FatherTaskDoesNotExist   = 400117 // 父级任务不存在
	TaskDoesNotExit          = 400118 // 任务不存在
)

type Message struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
	Code int         `json:"code"`
	Lang string      `json:"-"`
}

// 新建消息
func NewMsg(lang string) *Message {
	var msg Message

	if lang == "" {
		msg.Lang = "zh_Cn"
	} else {
		msg.Lang = lang
	}

	msg.Code = SUCCESS
	msg.Msg = GetMsg(SUCCESS, msg.Lang)

	return &msg
}

// 修改msg内容
func (m *Message) Update(code int, data interface{}) {
	m.Msg = GetMsg(code, m.Lang)
	m.Code = code
	m.Data = data
}
