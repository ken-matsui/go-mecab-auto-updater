package main

import (
    "fmt"
    "os/exec"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/gorilla/context"
    "cloud.google.com/go/storage"
    "google.golang.org/appengine/urlfetch"
)

const script = "./mecab-ipadic-neologd/bin/install-mecab-ipadic-neologd"
const nOp = "--newest"
const aOp = "--install_all_seed_files"
const yOp = "--forceyes"
const uOp = "--asuser"
const pOp = "--prefix "
// Get absolute path from relative path
const targetPath, _ = filepath.Abs("./")

// AppEngineが初回実行する関数
func init() {
    // MeCab Command
    err := exec.Command(script, nOp, aOp, yOp, uOp, pOp + targetPath).Run()

    // Routing
    r := mux.NewRouter()
    r.HandleFunc("/update", Handler).Methods("GET")
    // r.Host("localhost") // localhost以外からのアクセスを制限
    http.Handle("/", r)
}

// GETリクエスト時に呼び出される関数
func Handler(w http.ResponseWriter, r *http.Request) {
    err := exec.Command(script, nOp, aOp, yOp, uOp, pOp + targetPath).Run()
    fmt.Fprint(w, "done.") // wに"done."を書き込む
}