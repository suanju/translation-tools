package translation

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	RequestURL  = "https://api.fanyi.baidu.com/api/trans/vip/translate"
	ContentType = "application/x-www-form-urlencoded"
)

type TransResult struct {
	Src string `json:"src"`
	Dst string `json:"dst"`
}
type TranslateResp struct {
	From        string        `json:"from"`
	To          string        `json:"to"`
	TransResult []TransResult `json:"trans_result"`
	ErrorCode   string        `json:"error_code"`
	ErrorMsg    string        `json:"error_msg"`
}

func TranslateByBaidu(text, fromLang, toLang, APPID, KEY string) (resp TranslateResp, err error) {
	salt := time.Now().String()
	sign := generateSign(text, salt, APPID, KEY)
	params := url.Values{}
	params.Set("q", text)
	params.Set("from", fromLang)
	params.Set("to", toLang)
	params.Set("appid", APPID)
	params.Set("salt", salt)
	params.Set("sign", sign)
	translateResp, _ := http.Post(RequestURL, ContentType, strings.NewReader(params.Encode()))
	body, err := io.ReadAll(translateResp.Body)
	if err != nil {
		logx.Errorf("io readAll error :%s", err)
		return resp, err
	}
	resp, err = parseTranslationResponse(body)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func generateSign(text, salt, APPID, KEY string) string {
	signStr := APPID + text + salt + KEY
	signBytes := md5.Sum([]byte(signStr))
	return hex.EncodeToString(signBytes[:])
}

func parseTranslationResponse(response []byte) (resp TranslateResp, err error) {
	resp = TranslateResp{}
	err = json.Unmarshal(response, &resp)
	if err != nil {
		logx.Errorf("json unmarshal error :%s", response)
		return resp, err
	}
	return resp, err
}
