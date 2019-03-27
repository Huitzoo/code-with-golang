package rest

import (
	"encoding/json"
	"net/http"
	"errors"
	"fmt"
)

/*This struct is the response of frontend server*/
type RestResponse struct {
	Token   string         `json:"token"` // is the token that use for match between frontend and backend
	Message string      `json:"message"`//response some message
	Code    int         `json:"code"` // the http code 
	Payload interface{} `json:"payload"` //data objects send to the requests
}

/*Struct that use for request */
type RestRequest struct {
	Data interface{} `json:"data"` // data is a json that send in a post request.
	Token string `json:"token"`
}

func Response(w http.ResponseWriter, code int, payload interface{}, message string,token string) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("rest: ",payload)
	response := &RestResponse{}
	response.Code = code
	response.Message = message
	response.Payload = payload
	json.NewEncoder(w).Encode(response)
}

func Request(r *http.Request, typeData interface{}) (interface{}, error) {
	request := &RestRequest{Data: typeData}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		return nil, err
	}
	if request.Token == ""{
		return nil, errors.New("Token don't autorization")	
	}
	return request.Data,nil
}
