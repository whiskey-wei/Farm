package e

var MsgFlags = map[int]string{
	SUCCESS:                        "ok",
	ERROR:                          "fail",
	INVALID_PARAMS:                 "请求参数错误",
	ERROR_EXIST:                    "该数据已经存在",
	ERROR_NOT_EXIST:                "该数据不存在",
	ERROR_TIME_FORMAT:              "时间格式出错",
	ERROR_USER_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_USER_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_USER_TOKEN:               "Token生成失败",
	ERROR_USER:                     "Token错误",
	ERROR_PERMISSION_DENY:          "用户权限不够",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
