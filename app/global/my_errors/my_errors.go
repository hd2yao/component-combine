package my_errors

const (
	//系统部分
	ErrorsContainerKeyAlreadyExists string = "该键已经注册在容器中了"
	ErrorsConfigInitFail            string = "初始化配置文件发生错误"
	ErrorsBasePath                  string = "初始化项目根目录失败"

	ErrorsTokenBaseInfo    string = "token最基本的格式错误,请提供一个有效的token!"
	ErrorsNoAuthorization  string = "token鉴权未通过，请通过token授权接口重新获取token,"
	ErrorsRefreshTokenFail string = "token不符合刷新条件,请通过登陆接口重新获取token!"

	// 表单参数验证器未通过时的错误
	ErrorsValidatorNotExists      string = "不存在的验证器"
	ErrorsValidatorTransInitFail  string = "validator的翻译器初始化错误"
	ErrorNotAllParamsIsBlank      string = "该接口不允许所有参数都为空,请按照接口要求提交必填参数"
	ErrorsValidatorBindParamsFail string = "验证器绑定参数失败"
)
