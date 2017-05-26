package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"io/ioutil"
	"os"
	"fmt"
	"github.com/russross/blackfriday"
	"github.com/fairlyblank/md2min"
)

const (
	emtyStr = ""

	ReadmeMd    = "./README.md"
	ReadmeMDUrl = "https://raw.githubusercontent.com/avelino/awesome-go/master/README.md"
	indexHtml   = "./index.html"

	bfHTMLRendererOpts = 0 |
		blackfriday.HTML_USE_XHTML |
		blackfriday.HTML_USE_SMARTYPANTS |
		blackfriday.HTML_SMARTYPANTS_FRACTIONS |
		blackfriday.HTML_SMARTYPANTS_DASHES |
		blackfriday.HTML_SMARTYPANTS_LATEX_DASHES

	bfMDOpts = 0 |
		blackfriday.EXTENSION_NO_INTRA_EMPHASIS |
		blackfriday.EXTENSION_TABLES |
		blackfriday.EXTENSION_FENCED_CODE |
		blackfriday.EXTENSION_AUTOLINK |
		blackfriday.EXTENSION_STRIKETHROUGH |
		blackfriday.EXTENSION_SPACE_HEADERS |
		blackfriday.EXTENSION_HEADER_IDS |
		blackfriday.EXTENSION_BACKSLASH_LINE_BREAK |
		blackfriday.EXTENSION_DEFINITION_LISTS |
		blackfriday.EXTENSION_AUTO_HEADER_IDS
)

func doWithHtml(ch chan string) {
	_, err := os.Stat(ReadmeMd)
	if err == nil {
		//	file existes
		fmt.Println("start deleting ", ReadmeMd)
		if err := os.Remove(ReadmeMd); err != nil {
			fmt.Println("delete ", ReadmeMd, " failed")
			panic(err)
		} else {
			fmt.Println("delete ", ReadmeMd, " success")
		}
	}
	//	file not exist

	readme, err := http.Get(ReadmeMDUrl)
	if err != nil {
		fmt.Println("download ", ReadmeMDUrl, " failed")
		panic(err)
	} else {
		all, err := ioutil.ReadAll(readme.Body)
		if err == nil {
			fmt.Println("starting create ", indexHtml)
			f, _ := os.Create(indexHtml)
			f.Write(all)
			ch <- string(all);
		} else {
			fmt.Println(" create ", indexHtml)
			panic(err)
		}
	}

}

func serverHandler(w http.ResponseWriter, r *http.Request) {
	ch := make(chan string)
	go doWithHtml(ch)

	//result := <- ch
	//w.Write([]byte(<- ch))
	body :=
		blackfriday.Markdown(
			[]byte(<-ch),
			blackfriday.HtmlRenderer(
				bfHTMLRendererOpts,
				emtyStr,
				emtyStr,
			),
			bfMDOpts,
		)

	w.Write(body)

	//file, err := ioutil.ReadFile(indexHtml)
	//if err != nil {
	//	fmt.Println("read ", indexHtml, " failed")
	//	panic(err)
	//}
	//fmt.Println("here!")
	//w.Write(file)
}

func md2htmlHandler(w http.ResponseWriter, r *http.Request) {
	ch := make(chan string)
	go doWithHtml(ch)

	md := md2min.New("none")
	err := md.Parse([]byte(<-ch), w)
	if err != nil {
		panic(err)
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", serverHandler)
	router.HandleFunc("/{*}", md2htmlHandler)
	err := http.ListenAndServe(":9000", router)
	if err != nil {
		panic(err)
	}
}
