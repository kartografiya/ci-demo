package functions

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
)

func CalcHash(w http.ResponseWriter, r *http.Request) {
	var d struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		http.Error(w, "bad request!", http.StatusBadRequest)
		return
	}
	if d.Name == "" {
		fmt.Fprint(w, "Hello, %name%!")
		return
	}
	fmt.Fprintf(w, "Hello, %s!", html.EscapeString(d.Name))
}
