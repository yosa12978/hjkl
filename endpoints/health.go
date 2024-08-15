package endpoints

import (
	"net/http"

	"github.com/yosa12978/hjkl/util"
)

// Health endpoint godoc
//
//	@Summary		Healthcheck
//	@Description	Healthcheck
//	@Produce		json
//	@Success		200	{object}	types.Message
//	@Failure		500	{object}	types.Message
//	@Router			/health [get]
func Health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		util.WriteMsg(w, 200, "healthy")
	}
}
