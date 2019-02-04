package http

import (
	"encoding/json"
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/sqmmm/event-service/adapters/handlers/http/data"
	"github.com/sqmmm/event-service/usecases/finish"
	"github.com/sqmmm/event-service/usecases/start"
	"github.com/valyala/fasthttp"
	"log"
)

type handler struct {
	startUseCase  start.UseCase
	finishUseCase finish.UseCase
}

func NewHandler(start start.UseCase, finish finish.UseCase) *handler {
	return &handler{
		startUseCase:  start,
		finishUseCase: finish,
	}
}

func (h handler) Serve(addr string) {
	router := fasthttprouter.New()

	router.POST("/v1/start", h.start)
	router.POST("/v1/finish", h.finish)

	log.Fatal(fasthttp.ListenAndServe(addr, router.Handler))
}

func (h handler) start(ctx *fasthttp.RequestCtx) {
	request, err := getData(ctx)
	if err != nil {
		log.Println(err)
		ctx.Response.SetStatusCode(500)
		return
	}

	err = h.startUseCase.Start(request.Type)
	if err != nil {
		log.Println(err)
		ctx.Response.SetStatusCode(500)
	}
}

func (h handler) finish(ctx *fasthttp.RequestCtx) {
	request, err := getData(ctx)
	if err != nil {
		log.Println(err)
		ctx.Response.SetStatusCode(500)
		return
	}

	err = h.finishUseCase.Finish(request.Type)
	if err != nil {
		ctx.Response.SetStatusCode(500)
		log.Println(err)
	}
}

func getData(ctx *fasthttp.RequestCtx) (data.RequestData, error) {
	var request data.RequestData
	body := ctx.Request.Body()

	err := json.Unmarshal(body, &request)
	if err != nil {
		return request, fmt.Errorf("failed to unmarshal request: %s", err)
	}

	return request, nil
}
