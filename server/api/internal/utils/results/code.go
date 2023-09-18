package results

const (
	CodeSuccess       int = 200
	CodeInvalidParams int = 201
	CodeNoData        int = 202
	CodeServerBusy    int = 500

	CodeInvalidToken      int = 301
	CodeInvalidAuthFormat int = 302
	CodeNotLogin          int = 303

	CodeTypeError int = 415

	CodeDefaultError int = 1001
)

var _ = map[int]string{
	CodeSuccess:       "success",
	CodeInvalidParams: "请求参数错误",
	CodeServerBusy:    "服务繁忙",
	CodeNoData:        "没有数据",

	CodeInvalidToken:      "无效的Token",
	CodeInvalidAuthFormat: "认证格式有误",
	CodeNotLogin:          "未登录",

	CodeTypeError:    "类型错误",
	CodeDefaultError: "默认错误返回",
}
