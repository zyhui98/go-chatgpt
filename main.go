package main

import (
	"encoding/json"
	"fmt"
	"go-chatgpt/module/gpt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("html")))
	http.HandleFunc("/health", health)
	http.HandleFunc("/s", s)
	gpt.LoadCnf()

	go func() {
		log.Println("go in")
		defer func() {
			if err := recover(); err != nil {
				log.Println("go err:", err)
			}
		}()
		log.Println("go out")
	}()
	//handle定义请求访问该服务器里的/health路径，就有下面health去处理，health一般为健康检查
	err := http.ListenAndServe(fmt.Sprintf(":%v", gpt.GetPort()), nil)
	if err != nil {
		log.Fatal(err)
	}
}

//定义handle处理函数，只要该health被调用，就会写入ok
func health(w http.ResponseWriter, request *http.Request) {
	log.Println(request.URL)
	_ = request.ParseForm()
	log.Println(request.Form.Get("user"))
	_, _ = io.WriteString(w, "ok")
}

//定义handle处理函数，只要该health被调用，就会写入ok
func s(w http.ResponseWriter, request *http.Request) {
	log.Println(request.URL)
	//{
	//	"model": "gpt-3.5-turbo",
	//	"messages": [{"role": "user", "content": "Hello!"}]
	//}
	content := request.FormValue("content")
	assistant := request.FormValue("assistant")
	user := request.FormValue("user")
	messages := make([]map[string]interface{}, 3)
	msOne := make(map[string]interface{})
	msOne["role"] = "user"
	msOne["content"] = content
	msTwo := make(map[string]interface{})
	msTwo["role"] = "assistant"
	msTwo["content"] = assistant
	msThree := make(map[string]interface{})
	msThree["role"] = "system"
	msThree["content"] = "你是一个中文智能助手，用中文回答问题"
	messages[0] = msOne
	messages[1] = msTwo
	messages[2] = msThree

	req := &gpt.Req{Model: "gpt-3.5-turbo",
		//MaxTokens: 5000,
		Messages: messages, User: user}
	resp, _ := gpt.Api(req)
	respData, _ := json.Marshal(resp)
	_, _ = w.Write(respData)
}
