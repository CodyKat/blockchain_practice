package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/CodyKat/blockchain_practice/blockchain"
)

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	data := homeData{"Home", blockchain.GetBlockChain().AllBlocks()}
	templates.ExecuteTemplate(rw, "home", data)
}

func add(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(rw, "add", "Add Page")
	case "POST":
		r.ParseForm()
		data := r.Form.Get("blockdata")
		blockchain.GetBlockChain().AddBlock(data)
		http.Redirect(rw, r, "/", http.StatusPermanentRedirect)
	}

}

const (
	port        string = ":4000"
	templateDir string = "templates/"
)

var templates *template.Template

func main() {
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.html"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.html"))
	http.HandleFunc("/", home)
	http.HandleFunc("/add", add)
	fmt.Printf("i'm listening http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
