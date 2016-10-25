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
	"os"
	"github.com/lfany/go/xx/goweb/web/session"
)

var globalSessions *session.Manager
//然后在init函数中初始化
func init() {
	globalSessions, _ = session.NewManager("memory","gosessionid",3600)
}

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", upload)
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
	sess := globalSessions.SessionStart(w, r)
	r.ParseForm()

	fmt.Println("mthod: ", r.Method) // 请求的方法
	if r.Method == "GET" {
		currTime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(currTime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("login.gtpl")

		w.Header().Set("Content-Type", "text/html")
		t.Execute(w, token)
	} else {
		token := r.Form.Get("token")
		if token != "" {
		//	token存在， 验证token
		}else {
		//	token不存在， 错误
		}

		sess.Set("username", r.Form["username"])
		fmt.Println("sess->username: ", sess.Get("username"))

		fmt.Println("username: ", r.Form["username"])
		fmt.Println("password: ", r.Form["password"])
		//fmt.Println("xxxx: ", r)
		template.HTMLEscape(w, []byte(strings.Join([]string{r.Form.Get("username"), sess.Get("username").(string)}, "##"))) //输出到客户端

		/*for k, v := range r.Form {
			fmt.Println("key:", k)
			fmt.Println("val:", strings.Join(v, ""))
		}*/
	}
}

func upload(w http.ResponseWriter, r *http.Request){
	fmt.Println("methord: ", r.Method)

	if r.Method == "GET" {
		currTime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(currTime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		fmt.Println("Log: ##", currTime, strconv.FormatInt(currTime, 10), token)

		t, _ := template.ParseFiles("upload.gtpl")
		t.Execute(w, token)
	}else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadFile")
		if err != nil {
			fmt.Println("FormFile", err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)

		f, err := os.OpenFile("./test/" + handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil && os.IsNotExist(err){
			os.MkdirAll("./test", 0644)
			ff, err := os.OpenFile("./test/" + handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println("OpenFile", err)
				return
			}else{
				f = ff
			}
			defer ff.Close()
		}
		defer f.Close()
		io.Copy(f, file)
	}

}