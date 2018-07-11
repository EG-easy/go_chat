package main

import (
	"html/template"
	"net/http"
	"log"
	"path/filepath"
	"sync"
	"flag"
	"os"
	"go_programing/trace"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/facebook"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/objx"
)


//templは一つのテンプレートを表します
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

//serveHTTPはHTTPリクエストを返します
func (t *templateHandler)ServeHTTP(w http.ResponseWriter, r *http.Request){
	t.once.Do(func() {
		t.templ =
			template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})

	data := map[string]interface{}{
		"Host" : r.Host,
	}
	if authCookie, err := r.Cookie("auth"); err == nil{
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
	}
	t.templ.Execute(w, data)
}

func main() {
	var addr = flag.String("addr", ":8080", "アプリケーションのアドレス")
	flag.Parse() //フラグを解釈します
	//gomniauthのセットアップ
	gomniauth.SetSecurityKey("セキュリティキー")
	gomniauth.WithProviders(
		facebook.New("クライアントID", "秘密の値", "http://localhost:8080/auth/callback/facebook"),
		github.New("クライアントID", "秘密の値", "http://localhost:8080/auth/callback/github"),
		google.New("346358804589-7r228bs4i6g4hkvagpcl0bphau7i2grh.apps.googleusercontent.com", "mC9iZcU7rAsKAjN6pvLTNsmG", "http://localhost:8080/auth/callback/google"),
	)
	r := newRoom()
	r.tracer = trace.New(os.Stdout)
	//ルート
	http.Handle("/chat", MustAuth(&templateHandler{filename:"chat.html"}))
	http.Handle("/login", &templateHandler{filename:"login.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/room", r)
	//チャットを開始します
	go r.run()
	// サーバーを開始します
	log.Println("webサーバーを起動します。ポート：", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}






