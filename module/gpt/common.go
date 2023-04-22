package gpt

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var config map[string]interface{}
var gpt_config map[string]interface{}

type Resp struct {
	Code int                    `json:"code"`
	Body map[string]interface{} `json:"data"`
}

type Req struct {
	Model string `json:"model"`
	User  string `json:"user"`
	//MaxTokens int `json:"max_tokens"`
	Messages []map[string]interface{} `json:"messages"`
}

func LoadCnf() {
	log.Println("load site conf")
	path := "configs/config.yml"
	fi, _ := os.Open(path)
	configData, err := ioutil.ReadAll(fi)
	if err != nil {
		log.Fatal(err)
	}

	config = make(map[string]interface{})
	gpt_config = make(map[string]interface{})

	// 执行解析
	err = yaml.Unmarshal(configData, &config)
	log.Printf("config:%v\n", config)

	gpt := config["gpt"].(map[interface{}]interface{})

	for s, _ := range gpt {
		var key = s.(string)
		var value = gpt[s].(string)
		if strings.Contains(value, "$") {
			gpt_config[key] = os.Getenv(strings.ReplaceAll(
				strings.ReplaceAll(value, "${", ""), "}", ""))
		} else {
			gpt_config[key] = gpt[s]
		}
	}
	log.Printf("gpt_config:%v\n", gpt_config)

}

func GetPort() int {
	port, _ := strconv.Atoi(gpt_config["port"].(string))
	return port
}
