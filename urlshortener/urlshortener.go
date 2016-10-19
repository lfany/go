package main

import (
	"fmt"
	"net/http"
	"text/template"

	"google.golang.org/api/urlshortener/v1"
)

const key string = "AIzaSyCsGKjNxb965WdGJB2YLSH-9P0F_nB2hmE";

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/short", short)
	http.HandleFunc("/long", long)
	http.HandleFunc("/hello", hello)

	http.ListenAndServe("0.0.0.0:88", nil)
}

var rootHtmlTmpl = template.Must(template.New("rootHtml").Parse(`
<html><body>
<h1>URL SHORTENER</h1>
{{if .}}{{.}}<br /><br />{{end}}
<form action="/short" type="POST">
Shorten this: <input type="text" name="longUrl" />
<input type="submit" value="Give me the short URL" />
</form>
<br />
<form action="/long" type="POST">
Expand this: http://goo.gl/<input type="text" name="shortUrl" />
<input type="submit" value="Give me the long URL" />
</form>
</body></html>
`))

func root(w http.ResponseWriter, r *http.Request)  {
	rootHtmlTmpl.Execute(w, nil)
}

func short(w http.ResponseWriter, r *http.Request)  {
	longUrl := r.FormValue("longUrl")
	urlshortenerSvc, _ := urlshortener.New(http.DefaultClient)
	urlshortenerSvc.UserAgent
	url, err := urlshortenerSvc.Url.Insert(&urlshortener.Url{LongUrl: longUrl}).Do()
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	rootHtmlTmpl.Execute(w, fmt.Sprintf("Shortened version of %s is: %s", longUrl,
	url.Id))
}

func long(w http.ResponseWriter, r *http.Request) {
	shortUrl := "http://goo.gl/" + r.FormValue("shortUrl")
	urlshortenerSvc, _ := urlshortener.New(http.DefaultClient)
	url, err := urlshortenerSvc.Url.Get(shortUrl).Do()
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	rootHtmlTmpl.Execute(w, fmt.Sprintf("Longer version of %s is: %s",
		shortUrl, url.LongUrl))
}

func hello(w http.ResponseWriter, r *http.Request) {
	//fmt.Printf("%v\n%v", w, *r)
	w.WriteHeader(204)
	fmt.Printf("%v",w.Header())
	w.Write([]byte("lalla啦啦啦"))
}