package service

import (
	"context"
	"encoding/json"
	"fmt"
	endpoint "go-kit-demo/eat/endpoint/http"
	svc "go-kit-demo/eat/service"
	"net/http"

	"github.com/gorilla/mux"
)

func DecodeGetEatByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.GetEatByIdRequest
	fmt.Println("-------->>>>into GetById Decoding")
	vars := mux.Vars(r)
	req = endpoint.GetEatByIdRequest{
		Id: vars["eatid"],
	}

	return req, nil
}

func DecodeGetTotalEatByDateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.GetTotalEatByDateRequest
	fmt.Println("-------->>>>into GetById Decoding")
	vars := mux.Vars(r)
	req = endpoint.GetTotalEatByDateRequest{
		Date: vars["date"],
	}

	return req, nil
}

func DecodeGetAllEatsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	if r.URL.Query().Get("date") != "" {
		fmt.Println("-------->>>> Into GetEatsByDate Decoding")
		date := r.URL.Query().Get("date")
		return endpoint.GetAllEatsRequest{
			Date: date,
		}, nil
	}
	fmt.Println("-------->>>> Into GetEats Decoding")
	var req endpoint.GetAllEatsRequest
	return req, nil
}

func DecodeCreateEatRequest(_ context.Context, r *http.Request) (interface{}, error) {
	// var req endpoint.CreateEatRequest
	var req svc.Eat
	fmt.Println("-------->>>>into Decoding")
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func DecodeUpdateEatRequest(_ context.Context, r *http.Request) (interface{}, error) {
	fmt.Println("-------->>>> Into Update Decoding")
	// var req endpoint.UpdateEatRequest
	var req svc.Eat
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	vars := mux.Vars(r)
	req.Id = vars["eatid"]

	return req, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Println("into Encoding <<<<<<----------------")
	return json.NewEncoder(w).Encode(response)
}

func DecodeDeleteEatRequest(_ context.Context, r *http.Request) (interface{}, error) {
	fmt.Println("-------->>>> Into Delete Decoding")
	var req endpoint.DeleteEatRequest
	vars := mux.Vars(r)
	req = endpoint.DeleteEatRequest{
		Id: vars["eatid"],
	}
	return req, nil
}
