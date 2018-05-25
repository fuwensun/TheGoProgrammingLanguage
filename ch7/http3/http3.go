package main

import (
	"fmt"
	"net/http"
	"log"
)

type dollars float32

func (d dollars)String() string {
	return fmt.Sprintf("%.2f", d)
}

/* 
type Hndler interface {
	ServeHTTP(ResponseWriter, *Request)
}

type ServeMux struct {
	mu    sync.RWMutex
	m     map[string]muxEntry
	hosts bool // whether any patterns contain hostnames
}

type muxEntry struct {
	h       Handler
	pattern string
}
func NewServeMux() *ServeMux { return new(ServeMux) }

func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request) {
}

func (mux *ServeMux) Handle(pattern string, handler Handler) {

	mux.m[pattern] = muxEntry{h: handler, pattern: pattern}

}
func ListenAndServe(addr string, handler Handler) error {
	server := &Server{Addr: addr, Handler: handler}
	return server.ListenAndServe()
}
*/
func main() {
	db := database{"shoes":50, "socks":5}
	mux := http.NewServeMux()
	mux.Handle("/list",http.HandlerFunc(db.list))
	mux.Handle("/price",http.HandlerFunc(db.price))
	log.Fatal(http.ListenAndServe("localhost:8000",mux))
}

type database map[string]dollars

func (db database)list(w http.ResponseWriter, req *http.Request){
	for item, price := range db{
		fmt.Fprintf(w, "%s: %s\n",item, price)
	}
}

func (db database)price(w http.ResponseWriter, req *http.Request){
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok{
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w,"no such item: %q\n",item)
		return
	}
	fmt.Fprintf(w,"%s\n",price)
}