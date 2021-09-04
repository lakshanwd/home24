package handler

import (
	"encoding/json"
	"home-24/model"
	"home-24/service"
	"home-24/utils"
	"net/http"
)

type CrawlerHandler interface {
	Analyze(rw http.ResponseWriter, r *http.Request)
}

type crawlerHandler struct {
	client utils.HttpClient
}

func NewCrawlerhandler(client utils.HttpClient) CrawlerHandler {
	return &crawlerHandler{client: client}
}

func (handler *crawlerHandler) Analyze(rw http.ResponseWriter, r *http.Request) {
	// decode request body
	var req model.WebCrawlerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"status": false,
			"error":  err.Error(),
		})
		return
	}

	// handle the request
	if result, err := service.Analyze(r.Context(), handler.client, req.Url); err == nil {
		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"status": true,
			"result": result,
		})
	} else {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"status": false,
			"error":  err.Error(),
		})
	}
}
