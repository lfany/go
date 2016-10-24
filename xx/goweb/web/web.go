package main

import (
	"net/http"
	"log"
	"fmt"
	"strings"
	"html/template"
	"time"
	"crypto/md5"
	"io"
	"strconv"
)

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":88", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("val:", strings.Join(v, " "))
	}
	fmt.Fprintf(w, "hello asssss!")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("mthod: ", r.Method) // 请求的方法
	if r.Method == "GET" {
		currTime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(currTime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
		//	token存在， 验证token
		}else {
		//	token不存在， 错误
		}

		fmt.Println("username: ", r.Form["username"])
		fmt.Println("password: ", r.Form["password"])
		//fmt.Println("xxxx: ", r)
		template.HTMLEscape(w, []byte(r.Form.Get("username"))) //输出到客户端

		/*for k, v := range r.Form {
			fmt.Println("key:", k)
			fmt.Println("val:", strings.Join(v, ""))
		}*/
	}
}