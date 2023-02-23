package goQL

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"time"
)

type GoQLClientRunner interface {
	Run(ctx context.Context) interface{}
	RawReq() GoQLClientRunner
	RawRes() GoQLClientRunner
}

type Runner struct {
	method     string
	url        string
	body       []byte
	headers    map[string]string
	rawReq     bool
	rawRes     bool
	timeoutSec int
}

func (r *Runner) RawReq() GoQLClientRunner {
	r.rawReq = true
	return r
}

func (r *Runner) RawRes() GoQLClientRunner {
	r.rawRes = true
	return r
}

func (r *Runner) Run(ctx context.Context) interface{} {
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(r.timeoutSec))
	defer cancel()
	req, err := http.NewRequest(r.method, r.url, bytes.NewBuffer(r.body))
	if err != nil {
		L.Panic("ERROR from NewRequest", err)
	}
	for h, hv := range r.headers {
		req.Header.Set(h, hv)
	}

	req = req.WithContext(ctx)

	if r.rawReq {
		r.printRawRequest(req)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		L.Panic("ERROR from http client", err)
	}
	defer resp.Body.Close()

	if r.rawRes {
		r.printRawResponse(resp)
	}
	//L.Println("RESPONSE STATUS", resp.StatusCode)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		L.Panic("ERROR from ReadAll body", err)
	}
	//L.Println(string(body))

	var response interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		L.Panic("ERROR from Unmarshal response", err)
	}
	//L.Println("\n----Final:----\n", response)
	return response
}

func (r *Runner) printRawRequest(req *http.Request) {
	reqDump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		L.Fatal(err)
	}

	fmt.Printf("----------------:REQUEST:----------------\n%s", string(reqDump))
}

func (r *Runner) printRawResponse(resp *http.Response) {
	respDump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		L.Fatal(err)
	}

	fmt.Printf("\n----------------:RESPONSE:----------------\n%s", string(respDump))
}
