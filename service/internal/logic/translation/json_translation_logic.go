package translation

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logc"
	"strings"

	"translation/internal/svc"
	"translation/internal/types"
	tr "translation/internal/utils/translation"

	"github.com/zeromicro/go-zero/core/logx"
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
	resp = &types.TranslationJsonResp{}
	appid := "20230818001785620"
	key := "hF2iXa1a_mfxZIsrSP4v"
	var text string
	for _, v := range req.Json {
		if text == "" {
			text = fmt.Sprintf("%s", v)
		} else {
			text = fmt.Sprintf("%s \n %s", text, v)
		}
	}

	info, err := tr.TranslateByBaidu(text, req.Original, req.Results, appid, key)
	logc.Info(l.ctx, info.TransResult)
	if err != nil {
		logc.Errorf(l.ctx, "TranslateByBaidu error %v", err)
		return nil, err
	}
	fmt.Println(info.TransResult)
	//响应结果生成
	resp.Json = req.Json
	for _, v := range info.TransResult {
		for kk, vv := range resp.Json {
			fmt.Printf("循环 vv == v.Src %s %s %v \n", vv, v.Src, strings.TrimSpace(vv) == strings.TrimSpace(v.Src))
			if strings.TrimSpace(vv) == strings.TrimSpace(v.Src) {
				fmt.Printf("匹配成功 vv %s \n", vv)
				resp.Json[kk] = v.Dst
			}
		}
	}
	fmt.Println(resp.Json)

	return resp, nil
}
