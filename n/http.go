package n

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func PostReq(url string, params string) (string, error) {
	// 1. 创建http客户端实例
	client := &http.Client{}
	// 2. 创建请求实例
	req, err := http.NewRequest("POST", url, strings.NewReader(params))
	if err != nil {
		fmt.Println((err.Error()))
		return "", err
	}
	// 3. 设置请求头，可以设置多个
	req.Header.Set("Content-Type", "application/json")

	// 4. 发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println((err.Error()))
		return "", err
	}
	defer resp.Body.Close()

	// 5. 一次性读取请求到的数据
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println((err.Error()))
		return "", err
	}

	return string(body), err
}
