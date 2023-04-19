package transport

import (
	"context"
	"encoding/json"
	"fmt"
	endpoint "go-kit-demo/body/endpoint/http"
	svc "go-kit-demo/body/service"
	"net/http"

	"github.com/gorilla/mux"
)

func DecodeGetBodyByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.GetBodyByIdRequest
	fmt.Println("-------->>>>into GetById Decoding")
	vars := mux.Vars(r)
	req = endpoint.GetBodyByIdRequest{
		Id: vars["bodyid"],
	}

	return req, nil
}

func DecodeGetAllBodysRequest(_ context.Context, r *http.Request) (interface{}, error) {
	fmt.Println("-------->>>> Into GetBodys Decoding")
	var req endpoint.GetAllBodysRequest
	return req, nil
}

func DecodeCreateBodyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req svc.Body
	fmt.Println("-------->>>>into Decoding")
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	// req := endpoint.CreateBodyRequest{}
	// req.body = body

	return req, nil
}

func DecodeUpdateBodyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	fmt.Println("-------->>>> Into Update Decoding")
	vars := mux.Vars(r)
	var req svc.Body

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	req.Id = vars["bodyid"]
	return req, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Println("into Encoding <<<<<<----------------")
	return json.NewEncoder(w).Encode(response)
}

func DecodeDeleteBodyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	fmt.Println("-------->>>> Into Delete Decoding")
	var req endpoint.DeleteBodyRequest
	vars := mux.Vars(r)
	req = endpoint.DeleteBodyRequest{
		Id: vars["bodyid"],
	}
	return req, nil
}
