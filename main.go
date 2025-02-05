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
	tmpl := template.Must(template.ParseFiles("templates/pages/home.html"))
	data := homeData{"Home", blockchain.GetBlockChain().AllBlocks()}
	tmpl.Execute(rw, data)
}

const port string = ":4000"

func main() {
	http.HandleFunc("/", home)
	fmt.Printf("i'm listening http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
