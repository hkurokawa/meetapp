package views

import (
	"log"
	"net/http"

	"github.com/guregu/kami"
	"github.com/kyokomi/goroku"
	"golang.org/x/net/context"
)

func init() {
	kami.Get("/error", Error)
}

func Error(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	executeError(ctx, w, r, nil)
}

func executeError(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
	preload := NewHeader(ctx, "Error", "", "", false)

	if err != nil {
		log.Println("ERROR!", err)
		sendAirbrake(ctx, err, r)
		preload.SubTitle = err.Error()
	}

	if err := FromContextTemplate(ctx, "error").Execute(w, preload); err != nil {
		log.Println("ERROR!", err)
		renderer.JSON(w, 400, err.Error())
		return
	}
}

func sendAirbrake(ctx context.Context, err error, r *http.Request) {
	air, ok := goroku.Airbrake(ctx)
	if ok {
		go air.Notify(err, r)
	}
}
