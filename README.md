<!--
 * @Author: your name
 * @Date: 2024-01-02 18:40:01
 * @LastEditTime: 2024-01-03 16:57:12
 * @LastEditors: your name
 * @Description: 
 * @FilePath: \mfn-translation\README.md
 * 
-->


**项目环境 go 1.21 node 18.17.1**

项目预览

![image](https://github.com/suanju/translation-tools/blob/main/data/cover.png?raw=true)

项目目前支持接口 百度翻译 deepl 目前仅支持处理json格式翻译

🖊 TODO：
- 
- 对接谷歌翻译 微软翻译
- 支持多语言翻译结果
- 支持更多格式处理




**Setup**

###### service使用go-zero 启动命令
```
go run .\api.go -f .\etc\api.yaml

go build
```

其他项目结构参考 [go-zero](https://go-zero.dev/)


###### web基于Nuxt3 启动 建议使用 yarn

前端启动 : 修改.ENV 文件执行命令
```
npm install

npm run dev

npm run build
```

其他项目结构参考 [Nuxt3](https://nuxt.com/)