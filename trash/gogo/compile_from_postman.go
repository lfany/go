package main

import (
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
	"compress/gzip"
)

func main() {

	url := "http://www.runoob.com/api/compile.php"

	payload := strings.NewReader("code=%2F*%20%E6%88%91%E7%9A%84%E7%AC%AC%E4%B8%80%E4%B8%AA%20Swift%20%E7%A8%8B%E5%BA%8F%20*%2F%0Avar%20myString%20%3D%20%22Hello%2C%20World!%22%0A%0A%0Aprint(myString)&language=16")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("origin", "http://www.runoob.com")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("accept", "*/*")
	req.Header.Add("x-devtools-emulate-network-conditions-client-id", "4b64ff28-0ea6-45d7-b18a-142fc46cd195")
	req.Header.Add("x-requested-with", "XMLHttpRequest")
	req.Header.Add("x-devtools-request-id", "7774.356")
	req.Header.Add("save-data", "on")
	req.Header.Add("dnt", "1")
	req.Header.Add("referer", "http://www.runoob.com/try/runcode.php?filename=HelloWorld&type=swift")
	req.Header.Add("accept-encoding", "gzip, deflate")
	req.Header.Add("accept-language", "zh,en-US;q=0.8,en;q=0.6,zh-CN;q=0.4")
	req.Header.Add("cookie", "UM_distinctid=15c30485541458-0275d1dc057e85-30637509-1fa400-15c30485542604; CNZZDATA5578006=cnzz_eid%3D785875425-1495455155-%26ntime%3D1495460555; Hm_lvt_8e2a116daf0104a78d601f40a45c75b4=1495458676; Hm_lpvt_8e2a116daf0104a78d601f40a45c75b4=1495460611")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("postman-token", "8b36f471-a7c0-c7ff-f433-a93abfcfd93b")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	var body []byte
	if res.Header.Get("Content-Encoding") == "gzip" {
		reader, _ := gzip.NewReader(res.Body)
		body, _ = ioutil.ReadAll(reader)
	} else {
		body, _ = ioutil.ReadAll(res.Body)
	}

	fmt.Println(res)
	fmt.Println(string(body))

}
