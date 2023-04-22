package gpt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Api(req *Req) (*Resp, error) {
	client := &http.Client{
		Transport: tr,
	}

	reqs, _ := json.Marshal(req)
	bodyCopy := ioutil.NopCloser(bytes.NewBuffer(reqs))

	log.Printf("api body: %v\n", string(reqs))
	//提交请求
	request, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bodyCopy)
	if err != nil {
		log.Println(err)
	}

	//增加header选项
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", gpt_config["authorization"].(string))

	resp := &Resp{Code: 200}

	//处理返回结果
	response, e := client.Do(request)
	if response == nil {
		resp.Code = 500
		log.Printf("response nil: %v\n", e)
		return resp, nil
	}
	if response.StatusCode != 200 {
		resp.Code = response.StatusCode
		log.Printf("status code error: %d %s\n", response.StatusCode, response.Status)
		return resp, nil
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err.Error())
	}

	var data map[string]interface{}
	json.Unmarshal(body, &data)
	fmt.Printf("Results: %v\n", data)

	resp.Code = 200
	resp.Body = data

	return resp, nil

}
