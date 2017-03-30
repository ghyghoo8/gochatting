package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"./chat"
	"fmt"
	"flag"
	"strings"
	"github.com/hoisie/mustache"
	"time"
	"strconv"
)

func main() {
	var dir string
	// usage 是说明参数，没有什么功能作用。。。
	flag.StringVar(&dir, "dir", "./public/assets/", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()

	r := mux.NewRouter()
	r.HandleFunc("/chat", chat.BuildConnection)
	//静态文件
	// html
	r.HandleFunc("/public/", pageHandler)
	//assets
	r.PathPrefix("/public/assets/").Handler(http.StripPrefix("/public/assets/", http.FileServer(http.Dir(dir))))

	fmt.Println(dir)

	// 开启路由===
	http.Handle("/", r)

	go chat.InitChatRoom()

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

//路由 页面html
func pageHandler(w http.ResponseWriter, r *http.Request) {
	_path := r.URL.Path[len("/"):]
	//fmt.Println(r.URL,_path)
	if strings.IndexAny(_path, "html") < 4 {
		_path += "index.html"
	}
	content := mustache.RenderFile(_path, map[string]string{"name": "cc哈哈", "t": getTimeLineStr()})
	//fmt.Println(content, _path)
	//http.ServeFile(w, r, _path)
	//http 输出html
	w.Write([]byte(content))
}

/**
 * 获取时间错
 */
func getTimeLineStr() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}
