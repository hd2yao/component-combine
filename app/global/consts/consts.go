package consts

const (
    // 进程被结束
    ProcessKilled string = "收到信号，进程被结束"
    // 表单验证器前缀
    ValidatorPrefix              string = "Form_Validator_"
    ValidatorParamsCheckFailCode int    = -400300
    ValidatorParamsCheckFailMsg  string = "参数校验失败"

    //服务器代码发生错误
    ServerOccurredErrorCode int    = -500100
    ServerOccurredErrorMsg  string = "服务器内部发生代码执行错误,请联系开发者排查错误日志"

    // token相关
    JwtTokenOK            int    = 200100                                  //token有效
    JwtTokenInvalid       int    = -400100                                 //无效的token
    JwtTokenExpired       int    = -400101                                 //过期的token
    JwtTokenFormatErrCode int    = -400102                                 //提交的 token 格式错误
    JwtTokenFormatErrMsg  string = "提交的 token 格式错误"                 //提交的 token 格式错误
    JwtTokenMustValid     string = "token为必填项,请在请求header部分提交!" //提交的 token 格式错误
    GinSetTrustProxyError string = "Gin 设置信任代理服务器出错"

    // CURD 常用业务状态码
    CurdStatusOkCode         int    = 200
    CurdStatusOkMsg          string = "成功"
    CurdCreatFailCode        int    = -400200
    CurdCreatFailMsg         string = "新增失败"
    CurdUpdateFailCode       int    = -400201
    CurdUpdateFailMsg        string = "更新失败"
    CurdDeleteFailCode       int    = -400202
    CurdDeleteFailMsg        string = "删除失败"
    CurdSelectFailCode       int    = -400203
    CurdSelectFailMsg        string = "查询无数据"
    CurdRegisterFailCode     int    = -400204
    CurdRegisterFailMsg      string = "注册失败"
    CurdLoginFailCode        int    = -400205
    CurdLoginFailMsg         string = "登录失败"
    CurdRefreshTokenFailCode int    = -400206
    CurdRefreshTokenFailMsg  string = "刷新Token失败"
    CurdTokenFailCode        int    = -400207
    CurdTokenFailMsg         string = "token错误"
    CurdMacFailCode          int    = -400208
    CurdMacFailMsg           string = "已经被绑定"
    CurdOrderFailCode        int    = -400209
    CurdOrderFailMsg         string = "只有预约中的订单才能取消"

    // 图片保存地址
    Filepath      = "E:/DesktopFile/Bycoin/ComponentsCombinations_1209_female_weapon"
    FileShu       = Filepath + "/Shu"
    FileWei       = Filepath + "/Wei"
    FileWu        = Filepath + "/Wu"
    FileNeutral   = Filepath + "/Neutral"
    FileOfRarityN = "/Rarity_N"
    FileOfRarityR = "/Rarity_R"

    ClothesN = FileOfRarityN + "/clothes/"
    ClothesR = FileOfRarityR + "/clothes/"

    ShuNBgFile         = "resource/Shu/N/1-background/"
    ShuNBodyFile       = "resource/Shu/N/2-body/"
    ShuNFaceFile       = "resource/Shu/N/3-face/"
    ShuNPosthairFile   = "resource/Shu/N/4-posthair/"
    ShuNEarFile        = "resource/Shu/N/5-ear/"
    ShuNEyeFile        = "resource/Shu/N/6-eye/"
    ShuNClothesFile    = "resource/Shu/N/7-clothes/"
    ShuNMouthFile      = "resource/Shu/N/8-mouth/"
    ShuNFormerhairFile = "resource/Shu/N/9-formerhair/"
    ShuNHeadFile       = "resource/Shu/N/10-head/"
    ShuNWeaponFile     = "resource/Shu/N/11-weapon/"

    ShuRBgFile         = "resource/Shu/R/1-background/"
    ShuRFroWeaponFile  = "resource/Shu/R/2-frotools/"
    ShuRBodyFile       = "resource/Shu/R/3-body/"
    ShuRFaceFile       = "resource/Shu/R/4-face/"
    ShuRPosthairFile   = "resource/Shu/R/5-posthair/"
    ShuREarFile        = "resource/Shu/R/6-ear/"
    ShuREyeFile        = "resource/Shu/R/7-eye/"
    ShuRClothesFile    = "resource/Shu/R/8-clothes/"
    ShuRMouthFile      = "resource/Shu/R/9-mouth/"
    ShuRFormerhairFile = "resource/Shu/R/10-formerhair/"
    ShuRHeadFile       = "resource/Shu/R/11-head/"
    ShuRBehWeaponFile  = "resource/Shu/R/12-behtools/"

    WeiNBgFile         = "resource/Wei/N/1-background/"
    WeiNBodyFile       = "resource/Wei/N/2-body/"
    WeiNFaceFile       = "resource/Wei/N/3-face/"
    WeiNPosthairFile   = "resource/Wei/N/4-posthair/"
    WeiNEarFile        = "resource/Wei/N/5-ear/"
    WeiNEyeFile        = "resource/Wei/N/6-eye/"
    WeiNClothesFile    = "resource/Wei/N/7-clothes/"
    WeiNMouthFile      = "resource/Wei/N/8-mouth/"
    WeiNFormerhairFile = "resource/Wei/N/9-formerhair/"
    WeiNHeadFile       = "resource/Wei/N/10-head/"
    WeiNWeaponFile     = "resource/Wei/N/11-weapon/"

    WeiRBgFile         = "resource/Wei/R/1-background/"
    WeiRFroWeaponFile  = "resource/Wei/R/2-frotools/"
    WeiRBodyFile       = "resource/Wei/R/3-body/"
    WeiRFaceFile       = "resource/Wei/R/4-face/"
    WeiRPosthairFile   = "resource/Wei/R/5-posthair/"
    WeiREarFile        = "resource/Wei/R/6-ear/"
    WeiREyeFile        = "resource/Wei/R/7-eye/"
    WeiRClothesFile    = "resource/Wei/R/8-clothes/"
    WeiRMouthFile      = "resource/Wei/R/9-mouth/"
    WeiRFormerhairFile = "resource/Wei/R/10-formerhair/"
    WeiRHeadFile       = "resource/Wei/R/11-head/"
    WeiRBehWeaponFile  = "resource/Wei/R/12-behtools/"

    WuNBgFile         = "resource/Wu/N/1-background/"
    WuNBodyFile       = "resource/Wu/N/2-body/"
    WuNFaceFile       = "resource/Wu/N/3-face/"
    WuNPosthairFile   = "resource/Wu/N/4-posthair/"
    WuNEarFile        = "resource/Wu/N/5-ear/"
    WuNEyeFile        = "resource/Wu/N/6-eye/"
    WuNClothesFile    = "resource/Wu/N/7-clothes/"
    WuNMouthFile      = "resource/Wu/N/8-mouth/"
    WuNFormerhairFile = "resource/Wu/N/9-formerhair/"
    WuNHeadFile       = "resource/Wu/N/10-head/"
    WuNWeaponFile     = "resource/Wu/N/11-weapon/"

    WuRBgFile         = "resource/Wu/R/1-background/"
    WuRFroWeaponFile  = "resource/Wu/R/2-frotools/"
    WuRBodyFile       = "resource/Wu/R/3-body/"
    WuRFaceFile       = "resource/Wu/R/4-face/"
    WuRPosthairFile   = "resource/Wu/R/5-posthair/"
    WuREarFile        = "resource/Wu/R/6-ear/"
    WuREyeFile        = "resource/Wu/R/7-eye/"
    WuRClothesFile    = "resource/Wu/R/8-clothes/"
    WuRMouthFile      = "resource/Wu/R/9-mouth/"
    WuRFormerhairFile = "resource/Wu/R/10-formerhair/"
    WuRHeadFile       = "resource/Wu/R/11-head/"
    WuRBehWeaponFile  = "resource/Wu/R/12-behtools/"

    NeutralNBgFile      = "resource/Neutral/N/1-background/"
    NeutralNBodyFile    = "resource/Neutral/N/2-body/"
    NeutralNIneyeFile   = "resource/Neutral/N/3-eye/"
    NeutralNClothesFile = "resource/Neutral/N/6-clothes/"
    NeutralNHeadFile    = "resource/Neutral/N/5-posthair/"
    NeutralNEarFile     = "resource/Neutral/N/4-ear/"
    NeutralNMouthFile   = "resource/Neutral/N/7-mouth/"
    NeutralNOuteyeFile  = "resource/Neutral/N/8-formerhair/"

    NeutralRBgFile      = "resource/Neutral/R/1-background/"
    NeutralRBodyFile    = "resource/Neutral/R/2-body/"
    NeutralRIneyeFile   = "resource/Neutral/R/3-eye/"
    NeutralRClothesFile = "resource/Neutral/R/6-clothes/"
    NeutralRHeadFile    = "resource/Neutral/R/5-posthair/"
    NeutralREarFile     = "resource/Neutral/R/4-ear/"
    NeutralRMouthFile   = "resource/Neutral/R/7-mouth/"
    NeutralROuteyeFile  = "resource/Neutral/R/8-formerhair/"
)
