package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("say:", r.Method)
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!")
}

func web_index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("web_index method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("index.html")
		t.Execute(w, nil)
		fmt.Fprintf(w, "Hello Paul!")

	} else {
		fmt.Fprint(w, "Error Request")
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login method:", r.Method)
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("login.html")
		t.Execute(w, token)
	} else {
		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
			fmt.Fprint(w, "already have token")
		} else {
			fmt.Fprint(w, "not yet have token")
		}
		if len(r.Form["username"][0]) == 0 {
			fmt.Fprint(w, "Please enter username.")
			return
		} else if len(r.Form["password"][0]) == 0 {
			fmt.Fprint(w, "Please enter your password.")
			return
		}
		fmt.Fprintln(w, "username:", r.Form["username"])
		fmt.Fprintln(w, "password:", r.Form["password"])
	}
}

func login2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login2 method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login2.html")
		t.Execute(w, nil)
	} else {
		fmt.Fprintln(w, "e")
		r.ParseForm()
		if len(r.Form["username"][0]) == 0 {
			fmt.Fprint(w, "Please enter username.")
			return
		} else if len(r.Form["password"][0]) == 0 {
			fmt.Fprint(w, "Please enter your password.")
			return
		}
		fmt.Fprintln(w, "username:", r.Form["username"])
		fmt.Fprintln(w, "password:", r.Form["password"])
	}
}

func main() {
	http.HandleFunc("/", web_index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/login2", login2)
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
