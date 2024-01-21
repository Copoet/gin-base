package util

const (
	// 参数类
	PublicError                = 1000
	PublicSuccess              = 2000
	PublicParamsNull           = 2001
	PublicLoginError           = 2002
	PublicExportEmpty          = 2003
	PublicParamsIllegal        = 2004
	PublicParamsAlreadyExist   = 2005
	PublicParamsLack           = 2006
	PublicUsernameAlreadyExist = 2007
	PublicPhoneAlreadyExist    = 2008
	PublicEmailAlreadyExist    = 2009
	PublicLogoutSuccess        = 2010
	// 短信类
	MessageTagError            = 3001
	MessageFormatError         = 3002
	MessageTemplatesError      = 3003
	MessageTemplatesParamError = 3004
	MessageTimestampError      = 3005
	// 权限类
	PublicAuthError        = 4000
	PublicAuthAlreadyExist = 4001
	// 身份类
	PublicTokenError = 5000
	PublicSignError  = 5001
	// 请求类
	PublicOtherHttpError = 6001
)

var Config = map[int]string{
	// 参数类
	PublicSuccess:              "处理成功",
	PublicError:                "处理失败",
	PublicParamsNull:           "操作参数为空",
	PublicLoginError:           "登陆失败，账号或密码错误",
	PublicExportEmpty:          "导出为空",
	PublicLogoutSuccess:        "退出成功",
	PublicParamsIllegal:        "参数非法",
	PublicParamsAlreadyExist:   "数据已存在",
	PublicParamsLack:           "缺少参数",
	PublicUsernameAlreadyExist: "用户名已存在",
	PublicPhoneAlreadyExist:    "手机号已存在",
	PublicEmailAlreadyExist:    "邮箱已存在",
	// 短信类
	MessageTagError:            "消息错误",
	MessageFormatError:         "格式错误",
	MessageTemplatesError:      "模板id错误",
	MessageTemplatesParamError: "缺少模板参数",
	MessageTimestampError:      "时间过期",
	// 权限类
	PublicAuthError:        "权限错误",
	PublicAuthAlreadyExist: "权限已存在",
	// 身份类
	PublicSignError:  "sign错误",
	PublicTokenError: "请求token错误",
	// http请求类
	PublicOtherHttpError: "系统内部请求网络错误", // 第三方请求网络
}

func GetMsg(code int) string {
	msg, ok := Config[code]
	if ok {
		return msg
	}

	return Config[PublicError]
}
