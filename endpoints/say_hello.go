package endpoints

import (
	"fmt"
	"net/http"

	"github.com/yosa12978/hjkl/util"
)

// hello handler godoc
//
//	@Summary		Say hello
//	@Description	just returns hello message
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Failure		500
//	@Router			/hello [get]
func SayHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.PathValue("name")
		if name == "" {
			name = "hjkl"
		}
		util.WriteMsg(w, 200, fmt.Sprintf("Hello %s!", name))
	}
}
