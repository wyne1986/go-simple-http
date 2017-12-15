package main
 
import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
)

var rootpath string = ".";

var mime = map[string]string{
	"js":"text/javascript",
	"css":"text/css",
	"html":"text/html",
	"htm":"text/html",
	"png":"image/png",
	"jpg":"image/jpeg",
	"jpeg":"image/jpeg",
	"gif":"image/gif",
	"woff":"font/woff",
	"svg":"image/svg+xml"}

func main() {
	fmt.Println("created by github.com/wyne1986, for web developer simple httpServer");
	fmt.Println("server start on port :88");
    http.HandleFunc("/", Handler)
    http.ListenAndServe(":88", nil)
}

func Handler(res http.ResponseWriter, req *http.Request){

	res.Header().Set("Content-Type","application/x-ms-application")
	arr := strings.Split(req.URL.Path,".");
	var lenth int = len(arr);
	ctype, ok := mime[arr[lenth-1]]
	if ok {
		res.Header().Set("Content-Type",ctype)
	}
	dat, err := ioutil.ReadFile(rootpath+req.URL.Path)
	if(err!=nil){
		fmt.Println(err);
		http.Error(res,err.Error(),404);
	}
	res.Write([]byte(string(dat)));
}