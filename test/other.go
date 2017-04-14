package main

import (
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
)

func main() {

	url := "https://test201503.yobangbang.com/ybbweb/personalServer/incomeRecord.do"

	payload := strings.NewReader("param=%7B%22type%22%3A201703%2C%22logisticsId%22%3A12376350%2C%22pwd%22%3A%22096e07a5eb44b4b63c093f2729cf5019%22%2C%22source%22%3A%22Android%22%2C%22pageNum%22%3A%20%221%22%2C%22pageSize%22%3A%20%2220%22%7D")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("postman-token", "8e554bf4-3ce9-b8b5-cef5-4b265cbaa5ac")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}