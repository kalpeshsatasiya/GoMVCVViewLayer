package main

import (
	"net/http"
	"html/template"
	"os"
	"strings"
	"bufio"
	"GoMVCVViewLayer/src/viewmodels"
)

func main() {
	templates := populateTemplates()

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		requestedFile := req.URL.Path[1:]

		template := templates.Lookup(requestedFile + ".html")

		var context interface{} = nil
		switch requestedFile {
		case "index":
			context = viewmodels.GetIndex()
		case "about" :
			context = viewmodels.GetAboutUs()
		case "pricing":
			context =viewmodels.GetPricing()
		}

		if template != nil {
			template.Execute(w, context)
		}else {
			w.WriteHeader(404)
		}
	})

	http.HandleFunc("/assets/", serveResource)
	/*http.HandleFunc("/img/", serveResource)
	http.HandleFunc("/js/", serveResource)
	http.HandleFunc("/fonts/", serveResource)
*/
	http.ListenAndServe(":8070", nil)
}
func serveResource(w http.ResponseWriter, req *http.Request) {
	path := "public" + req.URL.Path

	var contentType string
	if (strings.HasSuffix(path, ".css")) {
		contentType = "text/css"
	} else if (strings.HasSuffix(path, ".png")) {
		contentType = "image/png"
	} else if (strings.HasSuffix(path, ".js")) {
		contentType = "application/javascript"
	} else if (strings.HasSuffix(path, ".jpg")) {
		contentType = "image/jpg"
	} else if (strings.HasSuffix(path, ".woff")) {
		contentType = "application/x-font-woff"
	}else if (strings.HasSuffix(path, ".ttf")) {
		contentType = "application/x-font-ttf"
	} else if (strings.HasSuffix(path, ".gif")) {
		contentType = "image/gif"
	} else {
		contentType = "text/plain"
	}

	f, err := os.Open(path)

	if err == nil {
		defer f.Close()
		w.Header().Add("content-type", contentType)

		br := bufio.NewReader(f)
		br.WriteTo(w)
	}else {
		w.WriteHeader(404)
	}

}

func populateTemplates() *template.Template {
	result := template.New("templates")

	basePath := "templates"
	templateFolder, _ := os.Open(basePath)

	defer templateFolder.Close()

	templatePathRaw, _ := templateFolder.Readdir(-1)
	templatePaths := new([] string)

	for _, pathInfo := range templatePathRaw {
		if !pathInfo.IsDir() {
			*templatePaths = append(*templatePaths, basePath + "/" + pathInfo.Name())
		}
	}
	result.ParseFiles(*templatePaths...)

	return result

}