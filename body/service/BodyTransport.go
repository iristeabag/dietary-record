package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func DecodeGetBodyByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req GetBodyByIdRequest
	fmt.Println("-------->>>>into GetById Decoding")
	vars := mux.Vars(r)
	req = GetBodyByIdRequest{
		Id: vars["bodyid"],
	}

	return req, nil
}

func DecodeGetAllBodysRequest(_ context.Context, r *http.Request) (interface{}, error) {
	fmt.Println("-------->>>> Into GetBodys Decoding")
	var req GetAllBodysRequest
	return req, nil
}

func DecodeCreateBodyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req CreateBodyRequest
	fmt.Println("-------->>>>into Decoding")
	if err := json.NewDecoder(r.Body).Decode(&req.body); err != nil {
		return nil, err
	}
	return req, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Println("into Encoding <<<<<<----------------")
	return json.NewEncoder(w).Encode(response)
}

func DecodeDeleteBodyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	fmt.Println("-------->>>> Into Delete Decoding")
	var req DeleteBodyRequest
	vars := mux.Vars(r)
	req = DeleteBodyRequest{
		Id: vars["bodyid"],
	}
	return req, nil
}
