package main

import (
	"fmt"
	"net/http"
	"net/url"
	"io/ioutil"
	"bytes"
	"compress/gzip"
	"encoding/json"
)

func main() {

	reqUrl := "http://www.runoob.com/api/compile.php"

	values := url.Values{
		"code":     {"/* 我的第一个 Swift 程序 */\nvar myString = \"Hello, World!\"\nprint(myString)"},
		"language": {"16"},
	}

	reqBody := bytes.NewBufferString(values.Encode())
	req, _ := http.NewRequest("POST", reqUrl, reqBody)

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
	//req.Header.Add("cookie", "UM_distinctid=15c30485541458-0275d1dc057e85-30637509-1fa400-15c30485542604; CNZZDATA5578006=cnzz_eid%3D785875425-1495455155-%26ntime%3D1495460555; Hm_lvt_8e2a116daf0104a78d601f40a45c75b4=1495458676; Hm_lpvt_8e2a116daf0104a78d601f40a45c75b4=1495460611")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("postman-token", "badee672-b34a-40b8-6e3f-acf53c68293a")

	//fmt.Println("---------request-------\n", req, "----------------\n")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("--------err-req--------\n", err, "----------------\n")
	}

	//fmt.Println("--------res-Header--------\n", res.Header, "----------------\n")

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	//fmt.Println("--------response--------\n", res, "----------------\n")

	if res.Header.Get("Content-Encoding") == "gzip" {
		reader, err := gzip.NewReader(bytes.NewReader(body))
		if err != nil {
			fmt.Println("------err-ungzip----------\n", err, "----------------\n")
		}
		//fmt.Println("------res-Header2----------\n", res.Header, "----------------\n")
		body, _ = ioutil.ReadAll(reader)
	}

	fmt.Println("--------response-body--------\n", string(body), "----------------\n")
	err = ioutil.WriteFile("e:\\outfile", body, 0)
	if err != nil {
		println("--------err-read--------\n", err, "----------------\n")
	}

	var response SwiftResponse

	json.Unmarshal(body, &response)

	fmt.Println(response)

	fmt.Println(response.Output)

	var vv interface{}
	json.Unmarshal(body, &vv)
	m := vv.(map[string]interface{})

	for k, v := range m {
		switch vt := v.(type) {
		case string:
			fmt.Println(k, vt, v)
		case int:
			fmt.Println(k, vt, v)
		case float64:
			fmt.Println(k, vt, v)
		case []interface{}:
			fmt.Println(k, vt, v)
			for i, u := range vt {
				fmt.Println(i, u)
			}

		}
	}

	fmt.Println(vv)

}

type SwiftResponse struct {
	Output string        `json:"output, string"`
	Langid int        `json:"langid, int"`
	Code   string        `json:"code, string"`
	Errors string        `json:"errors, string"`
	Time   string        `json:"time, string"`
}
