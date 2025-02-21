package explorer

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/CodyKat/blockchain_practice/blockchain"
)

const (
	templateDir string = "explorer/templates/"
)

var port string

var templates *template.Template

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	data := homeData{"Home", nil}
	templates.ExecuteTemplate(rw, "home", data)
}

func add(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(rw, "add", "Add Page")
	case "POST":
		blockchain.BlockChain().AddBlock()
		http.Redirect(rw, r, "/", http.StatusPermanentRedirect)
	}

}

func Start(aPort int) {
	handler := http.NewServeMux()
	port = fmt.Sprintf(":%d", aPort)
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.html"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.html"))
	handler.HandleFunc("/", home)
	handler.HandleFunc("/add", add)
	fmt.Printf("i'm listening http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, handler))
}
