package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func DecodeGetEatByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req GetEatByIdRequest
	fmt.Println("-------->>>>into GetById Decoding")
	vars := mux.Vars(r)
	req = GetEatByIdRequest{
		Id: vars["eatid"],
	}

	return req, nil
}

func DecodeGetTotalEatByDateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req GetTotalEatByDateRequest
	fmt.Println("-------->>>>into GetById Decoding")
	vars := mux.Vars(r)
	req = GetTotalEatByDateRequest{
		Date: vars["date"],
	}

	return req, nil
}

func DecodeGetAllEatsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	if r.URL.Query().Get("date") != "" {
		fmt.Println("-------->>>> Into GetEatsByDate Decoding")
		date := r.URL.Query().Get("date")
		return GetAllEatsRequest{
			Date: date,
		}, nil
	}
	fmt.Println("-------->>>> Into GetEats Decoding")
	var req GetAllEatsRequest
	return req, nil
}

func DecodeCreateEatRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req CreateEatRequest
	fmt.Println("-------->>>>into Decoding")
	if err := json.NewDecoder(r.Body).Decode(&req.eat); err != nil {
		return nil, err
	}
	return req, nil
}

func DecodeUpdateEatRequest(_ context.Context, r *http.Request) (interface{}, error) {
	fmt.Println("-------->>>> Into Update Decoding")
	var req UpdateEatRequest
	if err := json.NewDecoder(r.Body).Decode(&req.eat); err != nil {
		return nil, err
	}
	return req, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Println("into Encoding <<<<<<----------------")
	return json.NewEncoder(w).Encode(response)
}

func DecodeDeleteEatRequest(_ context.Context, r *http.Request) (interface{}, error) {
	fmt.Println("-------->>>> Into Delete Decoding")
	var req DeleteEatRequest
	vars := mux.Vars(r)
	req = DeleteEatRequest{
		Id: vars["eatid"],
	}
	return req, nil
}
