package explorer

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/hansuh/kapicoin/blockchain"
)

const (
	port        string = ":4000"
	templateDir string = "explorer/templates/"
)

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

var templates *template.Template

func homeHandler(w http.ResponseWriter, r *http.Request) {
	allBlocks := blockchain.GetBlockChain().AllBlocks()
	data := homeData{PageTitle: "Home", Blocks: allBlocks}
	templates.ExecuteTemplate(w, "home", data)
}

func addHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(w, "add", nil)
	case "POST":
		r.ParseForm()
		data := r.FormValue("blockData")
		blockchain.GetBlockChain().AddBlock(data)
		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
	}

}

func Start(aPort int) {

	handler := mux.NewRouter()
	port := fmt.Sprintf(":%d", aPort)

	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))

	handler.HandleFunc("/", homeHandler)
	handler.HandleFunc("/add", addHandler)

	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, handler))

}
