package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

//go:generate sh -c "cd frontend && spago deploy ../dist"
//go:generate cp -Rf frontend/assets dist/
//go:generate cp -f frontend/index.html dist/
//go:generate broccoli -src dist -quality 9

func init() {
	flag.Usage = func() {
		fmt.Printf("Usage: spago-slides [options] MARKDOWN-FILE\n\n")
		flag.PrintDefaults()
	}
}

func main() {
	br.Walk("dist", func(p string, info os.FileInfo, err error) error {
		log.Println(p, err)
		return nil
	})
	addr := ":8080"
	css := ""
	script := ""
	module := ""
	flag.StringVar(&addr, "l", ":8080", "listen address")
	flag.StringVar(&css, "css", "", "custom css file")
	flag.StringVar(&script, "js", "", "custom script file")
	flag.StringVar(&module, "module", "", "custom module file")
	flag.Parse()
	if flag.NArg() < 1 {
		flag.Usage()
		log.Fatal("you need arg a MARKDOWN-FILE")
	}
	mdfile := flag.Arg(0)
	log.Println("target markdown-file:", mdfile)
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	osfs := http.FileServer(http.Dir("."))
	if len(css) > 0 {
		http.HandleFunc("/user.css", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, css)
		})
	}
	if len(script) > 0 {
		http.HandleFunc("/user.js", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, script)
		})
	}
	if len(module) > 0 {
		http.HandleFunc("/module.js", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, module)
		})
	}
	http.HandleFunc("/content.md", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, mdfile)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		p := filepath.Join(pwd, strings.TrimLeft(r.URL.Path, "/"))
		if _, err := os.Stat(p); os.IsExist(err) {
			http.ServeFile(w, r, p)
			return
		}
		p = strings.Join(
			[]string{"dist", strings.TrimLeft(r.URL.Path, "/")},
			"/",
		)
		if strings.HasSuffix(p, "/") {
			p += "index.html"
		}
		log.Println("before:", p)
		if info, _ := br.Stat(p); info != nil {
			log.Println("after:", p)
			br.Serve("dist").ServeHTTP(w, r)
			return
		}
		osfs.ServeHTTP(w, r)
	})
	log.Println("listen:", addr)
	if err := http.Serve(l, nil); err != nil {
		log.Fatal(err)
	}
}
