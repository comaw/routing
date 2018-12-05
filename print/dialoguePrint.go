package print

import (
	"fmt"
	"net/http"
	"routing/errors"
)

func WebPrint(w http.ResponseWriter, text string) bool {
	_, err := fmt.Fprintf(w, "OK")
	errors.ErrorPanic(err)

	return true
}
