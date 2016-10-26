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
	_ "github.com/lfany/go/xx/goweb/web/session/providers"
)

var globalSessions *session.Manager
//然后在init函数中初始化
func init() {
	var err error

	globalSessions, err = session.NewManager("memory","gosessionid",3600)
	if(err != nil) {
		panic(err)
	}
	go globalSessions.GC()
}

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/count", count)
	http.HandleFunc("/count2", count2)
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

		err := sess.Set("username", strings.Join(r.Form["username"], "#"))
		if err != nil {
			panic(err)
		}
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

func count(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)
	createtime := sess.Get("createtime")
	if createtime == nil {
		sess.Set("createtime", time.Now().Unix())
	} else if (createtime.(int64) + 360) < (time.Now().Unix()) {
		globalSessions.SessionDestroy(w, r)
		sess = globalSessions.SessionStart(w, r)
	}
	ct := sess.Get("countnum")
	if ct == nil {
		sess.Set("countnum", 1)
	} else {
		sess.Set("countnum", (ct.(int) + 1))
	}
	t, _ := template.ParseFiles("count.gtpl")
	w.Header().Set("Content-Type", "text/html")
	t.Execute(w, sess.Get("countnum"))
}

func count2(w http.ResponseWriter, r *http.Request) {



	sess := globalSessions.SessionStart(w, r)
	ct := sess.Get("countnum")
	if ct == nil {
		sess.Set("countnum", 1)
	} else {
		sess.Set("countnum", (ct.(int) + 1))
	}

	h := md5.New()
	salt:="hello%^7&8888"
	io.WriteString(h,salt+time.Now().String())
	token:=fmt.Sprintf("%x",h.Sum(nil))
	if strings.Join(r.Form["token"], "")!=token{
		//提示登录
		w.Write([]byte("<script>alert(\"you are not login!\")</script>"))
		return
	}else{
		sess.Set("token",token)
		
		t, _ := template.ParseFiles("count.gtpl")
		w.Header().Set("Content-Type", "text/html")
		t.Execute(w, sess.Get("countnum"))
	}

}