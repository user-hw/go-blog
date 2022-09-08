package views

import (
	"go/common"

	"go/service"
	"net/http"
)

func (*HTMLApi) Pigeonhole(w http.ResponseWriter, r *http.Request) {
	pigeonhole := common.Template.Pigeonhole

	pigeonholeRes := service.FindPostPigeonhole

	pigeonhole.WriteData(w, pigeonholeRes)
}
