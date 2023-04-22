package gpt

import (
	"net"
	"net/http"
	"time"
)

var tr *http.Transport

func init() {
	//proxyUrl,_ := url.Parse("http://127.0.0.1:7890")
	//proxy := http.ProxyURL(proxyUrl)
	timeout := time.Second * 600
	tr = &http.Transport{
		// 设置代理
		//Proxy: proxy,
		MaxIdleConns: 100,
		Dial: func(netw, addr string) (net.Conn, error) {
			conn, err := net.DialTimeout(netw, addr, timeout) //设置建立连接超时
			if err != nil {
				return nil, err
			}
			err = conn.SetDeadline(time.Now().Add(timeout)) //设置发送接受数据超时
			if err != nil {
				return nil, err
			}
			return conn, nil
		},
	}
}
