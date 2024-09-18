package models

import (
	"bytes"
	"encoding/json"
	"fitroom-autotests/utils"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	FailedRequestArr  []Requester
	FailedMethodSlice []Requester
)

func Send(r *Request, req *http.Request) error {
	var err error
	if r.Token != nil {
		req.Header.Set("Authorization", *r.Token)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("\033[91mError: can't send a %s request:\033[0m \033[96m%s\033[0m - %s\n", r.Method, r.URL, err.Error())
		r.AddToFailed()
		return err
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("\033[91mError: can't read response body:\033[0m \033[96m%s\033[0m\n", err.Error())
		r.AddToFailed()
		return err
	}
	defer resp.Body.Close()
	// Check status
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		log.Printf("\033[96m%s\033[0m response code status: \033[33m%s\033[0m\n", utils.URLFormat(r.Method, r.URL), resp.Status)

		r.AddToFailed()
		return err
	}
	// Decode response
	if strings.HasPrefix(resp.Header.Get("Content-Type"), "application/json") && len(b) > 0 {
		if err = json.Unmarshal(b, &r.RespData); err != nil {
			log.Printf("\033[91mError: can't Unmarshal response body to RespData:\033[0m \033[96m%s\033[0m\n", utils.URLFormat(r.Method, r.URL))
			r.AddToFailed()
			return err
		}
	}
	// End
	log.Printf("\033[96m%s\033[0m response code status: \033[32m%s\033[0m\n", utils.URLFormat(r.Method, r.URL), resp.Status)
	for _, f := range r.Actions {
		f()
	}
	// Next
	if r.Next != nil {
		switch r.Next.Method {
		case "GET":
			r.Next.GET()
		case "POST":
			r.Next.POST()
		case "PUT":
			r.Next.PUT()
		case "DELETE":
			r.Next.DELETE()
		default:
			log.Printf("\033[96m%s\033[0m method \033[33m%s\033[0m is not supported \n", r.Next.Method, r.Next.URL)
		}
	}
	return nil
}

func (r *Request) GET() error {
	req, err := http.NewRequest("GET", r.URL, nil)
	if err != nil {
		log.Printf("\033[91mError: can't create new request: %s\033[0m\n", err.Error())
		r.AddToFailed()
		return err
	}
	return Send(r, req)
}

func (r *Request) POST() error {
	req, err := http.NewRequest("POST", r.URL, bytes.NewBuffer(r.ReqBody))
	if err != nil {
		log.Printf("\033[91mError: can't create new request: %s\033[0m\n", err.Error())
		r.AddToFailed()
		return err
	}
	return Send(r, req)
}

func (r *Request) PUT() error {
	req, err := http.NewRequest("PUT", r.URL, bytes.NewBuffer(r.ReqBody))
	if err != nil {
		log.Printf("\033[91mError: can't create new request: %s\033[0m\n", err.Error())
		r.AddToFailed()
		return err
	}
	return Send(r, req)
}

func (r *Request) DELETE() error {
	req, err := http.NewRequest("DELETE", r.URL, nil)
	if err != nil {
		log.Printf("\033[91mError: can't create new %s request: %s\033[0m\n", r.Method, err.Error())
		r.AddToFailed()
		return err
	}
	return Send(r, req)
}

func (r *Request) AddToFailed() {
	FailedRequestArr = append(FailedRequestArr, r)
}

func SaveFailedRequests() {
	file, err := json.MarshalIndent(FailedRequestArr, "", " ")
	if err != nil {
		log.Printf("Error: can't MarshalIndent failed requests: %s\n", err.Error())
		return
	}
	// currentTime := time.Now()
	// formattedTime := currentTime.Format("2006-01-02_15-04-05")
	// Dir
	err = os.MkdirAll("./failedRequests", os.ModePerm)
	if err != nil {
		fmt.Println("Ошибка при создании папки:", err)
		return
	}
	// Save
	if err = os.WriteFile("./failedRequests/failed.json", file, 0644); err != nil {
		log.Println("Error: can't save failed requests in file")
		return
	}
}

func LoadFailedRequests() {
	file, err := os.ReadFile("./failedRequests/failed.json")
	if err != nil {
		log.Println("Error: can't read file with failed requests")
		return
	}
	dataArr := []struct {
		URL     string
		Method  string
		ReqBody []byte
		Token   string
	}{}
	if err = json.NewDecoder(bytes.NewBuffer(file)).Decode(&dataArr); err != nil {
		log.Println("Error: can't Decode")
	}
	for _, failedReq := range dataArr {
		newReq := Request{
			URL:     failedReq.URL,
			Method:  failedReq.Method,
			ReqBody: failedReq.ReqBody,
			Token:   &failedReq.Token,
		}
		FailedMethodSlice = append(FailedMethodSlice, &newReq)
	}
}
