package helper

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

func GetMD5(str string) string {
	h := md5.New()
	salt := "vp"
	salt1 := "tine"
	io.WriteString(h, salt+str+salt1)
	urlmd5 := fmt.Sprintf("%x", h.Sum(nil))
	return urlmd5
}

func StrToTime(date string) int64 {
	parsed, err := time.Parse("2006-01-02 15:04:05", date)
	if err == nil {
		parsed := parsed.Unix()
		return parsed
	}
	return 0
}

func Upload(w http.ResponseWriter, r *http.Request) (string, bool) {

	r.ParseMultipartForm(32 << 20) //ParseMultipartForm parses a request body as multipart/form-data.
	//The whole request body is parsed and up to a total of maxMemory bytes of its file parts are stored in memory,
	//with the remainder stored on disk in temporary files.
	//ParseMultipartForm calls ParseForm if necessary. After one call to ParseMultipartForm, subsequent calls have no effect.
	file, handler, err := r.FormFile("image")
	if err != nil {
		fmt.Println(err)
		return "", false
	}
	defer file.Close()
	beego.Debug("Uploading-------------------------")
	//handler.Filename = randSeq(5)
	//fmt.Fprintf(w, "%v", handler.Header)
	savePath := "./static/upload/" + handler.Filename
	f, err := os.OpenFile(savePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return "", false
	}
	defer f.Close()
	io.Copy(f, file)
	savePath = strings.Trim(savePath, ".")
	return savePath, true
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func Substr(str string, start, length int, symbol string) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end]) + symbol
}

func WriteXlsx(data []string, path string) bool {

	return true
}

func RenderPartial(path string, data map[interface{}]interface{}) (template.HTML, error) {
	pf, _ := template.ParseFiles(path)
	ibytes := bytes.NewBufferString("")

	pf.ExecuteTemplate(ibytes, "msg", data)

	icontent, _ := ioutil.ReadAll(ibytes)
	html := template.HTML(string(icontent))

	return html, nil
}
