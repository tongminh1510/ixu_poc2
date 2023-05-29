package client

import (
	"PRS/entity"
	"bytes"
	"encoding/json"
	"net/http"
)

type OrderClient struct {
	BaseURL string
}

func NewOrderClient(baseURL string) *OrderClient {
	return &OrderClient{
		BaseURL: baseURL,
	}
}

func (oc *OrderClient) DoOrder(request entity.BillRequest, w http.ResponseWriter) (*http.Response, error) {
	requestBody, _ := json.Marshal(request)
	URL := NewOrderClient(oc.BaseURL)
	resp, err := http.Post(URL.BaseURL+"/payment/deduct", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, err
	}
	defer resp.Body.Close()
	return resp, err
}
