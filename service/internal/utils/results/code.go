package results

const (
	CodeTypeError         int = 415
	CodeNoPlatformCode    int = 10001
	CodeErrorPlatformCode int = 10002
	CodeNoFillConfig      int = 10003
)

var Message = map[int]string{
	CodeTypeError:         "类型错误",
	CodeNoPlatformCode:    "未选择翻译平台",
	CodeErrorPlatformCode: "错误的翻译平台",
	CodeNoFillConfig:      "未填写翻译配置",
}
