package core

import (
	"github.com/valyala/fasthttp"
	"time"
)

func Get(url string, timeout int) (status_code int, body string) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.Add("Connection", "close")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.6; rv,2.0.1) Gecko/20100101 Firefox/4.0.1")

	resp := fasthttp.AcquireResponse()

	client := &fasthttp.Client{}
	client.DoTimeout(req, resp, time.Duration(timeout)*time.Second)

	return resp.StatusCode(), string(resp.Body())

}
