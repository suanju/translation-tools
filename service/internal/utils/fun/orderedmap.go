package fun

import (
	"encoding/json"
	"fmt"
	"github.com/elliotchance/orderedmap"
	"strings"
)

// OrderedMapToJSON 将有序列表转化为Json
func OrderedMapToJSON(omap *orderedmap.OrderedMap) (string, error) {
	var jsonBuilder strings.Builder
	jsonBuilder.WriteString("{")
	first := true
	for el := omap.Front(); el != nil; el = el.Next() {
		if !first {
			jsonBuilder.WriteString(",")
		}
		first = false
		keyStr, _ := json.Marshal(el.Key)
		valueStr, _ := json.Marshal(el.Value)

		jsonBuilder.WriteString(fmt.Sprintf("%s:%s", keyStr, valueStr))
	}
	jsonBuilder.WriteString("}")

	return jsonBuilder.String(), nil
}

// SplitTextByMeasure 将字符字符串进行分割 按\n 进行
func SplitTextByMeasure(text string, maxLength int) []string {
	lines := strings.Split(text, "\n")
	var result []string
	currentSlice := ""
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)

		// 检查添加的当前行是否超过maxLength
		if len(currentSlice)+len(trimmedLine) > maxLength {
			result = append(result, currentSlice)
			currentSlice = trimmedLine
		} else {
			//将当前行添加到当前切片
			if currentSlice == "" {
				currentSlice = trimmedLine
			} else {
				currentSlice += "\n" + trimmedLine
			}
		}
	}
	//最后一个切片
	if currentSlice != "" {
		result = append(result, currentSlice)
	}
	return result
}

// GetStringLengthSum 获取字符串切片的长度的合
func GetStringLengthSum(strings []string) int {
	sum := 0
	for _, str := range strings {
		sum += len(str)
	}
	return sum
}

//// SplitTextByStringArr 将字符字切片进行分割
//func SplitTextByStringArr(text []string, textArr [][]string, maxLength int) [][]string {
//	cur := 0
//	for _, v := range text {
//		if GetStringLengthSum(textArr[cur])+len(text) > maxLength {
//			cur++
//			textArr = append(textArr, make([]string, 0))
//			continue
//		} else {
//			textArr[cur] = append(textArr[cur], v)
//		}
//	}
//	return textArr
//}

// SplitTextByStringArr 将字符字切片进行分割
func SplitTextByStringArr(text []string, textArr [][]string, maxLength int) [][]string {
	cur := 0
	for _, v := range text {
		if GetStringLengthSum(textArr[cur])+len(v) > maxLength {
			cur++
			textArr = append(textArr, make([]string, 0))
		}
		textArr[cur] = append(textArr[cur], v)
	}
	return textArr
}
