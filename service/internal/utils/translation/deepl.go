package translation

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type TransDeeplRequest struct {
	Text       []string `json:"text"`
	SourceLang string   `json:"source_lang,omitempty"`
	TargetLang string   `json:"target_lang"`
}

type TransDeeplResponse struct {
	Translations []Translation `json:"translations"`
}

type Translation struct {
	Text string `json:"text"`
}

func TranslateByDeepl(apiKey, sourceLang, targetLang string, text []string) (resp TransDeeplResponse, err error) {
	req := TransDeeplRequest{
		Text:       text,
		SourceLang: sourceLang,
		TargetLang: targetLang,
	}
	requestBody, err := json.Marshal(req)

	if err != nil {
		return resp, fmt.Errorf("error encoding request body: %v", err)
	}

	responseBody, err := sendTranslationRequest(apiKey, requestBody)
	if err != nil {
		return resp, err
	}

	// 解析翻译响应
	translatedText, err := parseTranslationResponseByDeepl(responseBody)
	if err != nil {
		return resp, err
	}
	return translatedText, nil
}

// sendTranslationRequest 封装了发送翻译请求的函数
func sendTranslationRequest(key string, requestBody []byte) ([]byte, error) {
	url := "https://api-free.deepl.com/v2/translate"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Authorization", "DeepL-Auth-Key "+key)
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	// 读取响应
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	return responseBody, nil
}

func parseTranslationResponseByDeepl(responseBody []byte) (resp TransDeeplResponse, err error) {
	var translationResponse TransDeeplResponse
	err = json.Unmarshal(responseBody, &translationResponse)
	if err != nil {
		return resp, fmt.Errorf("error decoding response body: %v", err)
	}
	// 检查是否有翻译结果
	if len(translationResponse.Translations) > 0 {
		return translationResponse, nil
	}

	fmt.Println(translationResponse)
	return resp, fmt.Errorf("translation not available")
}
