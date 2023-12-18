package translation

import (
	"context"
	"fmt"
	"github.com/elliotchance/orderedmap"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"regexp"
	"strings"
	"translation/internal/svc"
	"translation/internal/types"
	"translation/internal/utils/fun"
	tr "translation/internal/utils/translation"
)

type JsonTranslationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJsonTranslationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JsonTranslationLogic {
	return &JsonTranslationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JsonTranslationLogic) JsonTranslation(req *types.TranslationJsonReq) (resp *types.TranslationJsonResp, err error) {
	m := orderedmap.NewOrderedMap()
	// 使用正则表达式提取键值对
	re := regexp.MustCompile(`"([^"]+)":\s*("[^"]+"|'[^']+')`)
	matches := re.FindAllStringSubmatch(req.Json, -1)
	//得到一个有序map
	for _, match := range matches {
		key := match[1]
		value := match[2][1 : len(match[2])-1]
		m.Set(key, value)
	}

	//判断翻译平台
	switch l.ctx.Value("platform_code") {
	case "baidu":
		return TranslateByBaidu(l, req, m)
	case "deepl":
		return TranslateByDeepl(l, req, m)
	default:
		return nil, err
	}
}

// TranslateByBaidu 百度翻译
func TranslateByBaidu(l *JsonTranslationLogic, req *types.TranslationJsonReq, m *orderedmap.OrderedMap) (resp *types.TranslationJsonResp, err error) {
	resp = &types.TranslationJsonResp{}
	var trText strings.Builder
	for _, k := range m.Keys() {
		trText.WriteString(m.GetElement(k).Value.(string))
		trText.WriteString("\n")
	}
	text := trText.String()
	//进行分割字符限制长度
	textArr := fun.SplitTextByMeasure(text, 1500)
	resultArr := make([]tr.TransBaiduResult, 0)
	for _, v := range textArr {
		result, err := tr.TranslateByBaidu(v, req.Original, req.Results, l.ctx.Value("app_id").(string), l.ctx.Value("app_key").(string))
		if err != nil {
			logc.Errorf(l.ctx, "TranslateByBaidu error %v", err)
			return nil, err
		}
		if len(result.ErrorCode) != 0 {
			logc.Errorf(l.ctx, "TranslateByBaidu error_code : %s , err_msg : %s ", result.ErrorCode, result.ErrorMsg)
			return nil, fmt.Errorf("配置异常 : err : %s", result.ErrorMsg)
		}
		resultArr = append(resultArr, result.TransResult...)
	}
	//响应结果生成
	logx.Info(resultArr)
	for _, v := range resultArr {
		for _, vv := range m.Keys() {
			if strings.TrimSpace(m.GetElement(vv).Value.(string)) == strings.TrimSpace(v.Src) {
				m.Set(vv, strings.TrimSpace(v.Dst))
			}
		}
	}
	resp.Json, err = fun.OrderedMapToJSON(m)
	return
}

// TranslateByDeepl deepl翻译
func TranslateByDeepl(l *JsonTranslationLogic, req *types.TranslationJsonReq, m *orderedmap.OrderedMap) (resp *types.TranslationJsonResp, err error) {
	resp = &types.TranslationJsonResp{}
	//进行分割字符限制长度
	text := make([]string, 0)
	for _, v := range m.Keys() {
		text = append(text, m.GetElement(v).Value.(string))
	}
	var textArr = make([][]string, 1)
	textArr = fun.SplitTextByStringArr(text, textArr, 2000)

	//fmt.Println(textArr)
	//fmt.Println(len(textArr))
	//fmt.Println(fun.GetStringLengthSum(textArr[0]))

	resultArr := make([]tr.Translation, 0)
	for _, v := range textArr {
		result, err := tr.TranslateByDeepl(l.ctx.Value("app_key").(string), req.Original, req.Results, v)
		if err != nil {
			logc.Errorf(l.ctx, "TranslateByDeepl error %v", err)
			return nil, err
		}
		resultArr = append(resultArr, result.Translations...)
	}

	for k, v := range resultArr {
		m.Set(m.Keys()[k], strings.TrimSpace(v.Text))
	}
	resp.Json, err = fun.OrderedMapToJSON(m)
	return
}
