package handlers

import (
	"net/http"
	"log"
	"fmt"
	"io/ioutil"
)
type Home struct {
	l *log.Logger
}

func NewHome(l *log.Logger) *Home {
	return &Home{l}
}
func (h *Home) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("HOME")
	d, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(rw, "Data: %s \n", d)
	fmt.Fprintf(rw, "METHOD %s", r.Method)
	fmt.Fprintf(rw, " %s -- %s", r.URL, r.Form) 
	
}