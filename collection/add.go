package collection

import (
	"net/http"
	print2 "routing/print"
)

func AddCollection(w http.ResponseWriter, r *http.Request) {

	print2.WebPrint(w, "OK")
}
