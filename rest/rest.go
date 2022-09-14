package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hansuh/kapicoin/blockchain"
	"github.com/hansuh/kapicoin/utils"
)

var port string

type URL string

func (u URL) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

type urlDescription struct {
	URL         URL    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	PayLoad     string `json:"payload,omitempty"`
}

type addBlockBody struct {
	Message string
}

type errorResponse struct {
	ErrorMessage string `json:"errorMessage"`
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []urlDescription{
		{
			URL:         URL("/"),
			Method:      "GET",
			Description: "See Document",
		},
		{
			URL:         URL("/blocks"),
			Method:      "GET",
			Description: "See All Blocks",
		},
		{
			URL:         URL("/blocks"),
			Method:      "POST",
			Description: "Add A Block",
			PayLoad:     "data:string",
		},
		{
			URL:         URL("/blocks/{height}"),
			Method:      "GET",
			Description: "See A Block",
		},
	}
	// b, err := json.Marshal(data)
	// utils.HandleError(err)
	// fmt.Fprintf(rw, "%s", b)
	json.NewEncoder(rw).Encode(data)
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		json.NewEncoder(rw).Encode(blockchain.GetBlockChain().AllBlocks())
	case "POST":
		var addBlockBody addBlockBody
		utils.HandleError(json.NewDecoder(r.Body).Decode(&addBlockBody))
		blockchain.GetBlockChain().AddBlock(addBlockBody.Message)
		http.Redirect(rw, r, "/", http.StatusCreated)
	}
}

func block(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["height"])
	utils.HandleError(err)
	block, err := blockchain.GetBlockChain().GetBlock(id)
	encoder := json.NewEncoder(rw)
	if err == blockchain.ErrorNotFound {
		encoder.Encode(errorResponse{fmt.Sprint(err)})
	} else {
		encoder.Encode(block)
	}
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}

func Start(aPort int) {
	port = fmt.Sprintf(":%d", aPort)
	handler := mux.NewRouter()
	handler.Use(jsonContentTypeMiddleware)
	handler.HandleFunc("/", documentation).Methods("GET")
	handler.HandleFunc("/blocks", blocks).Methods("GET", "POST")
	handler.HandleFunc("/blocks/{height}", block).Methods("GET")

	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, handler))
}
