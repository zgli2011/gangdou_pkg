package gangdoupkg

import (
	"bytes"
	"context"
	"net/http"
)

const (
	POST    = "POST"
	GET     = "GET"
	HEAD    = "HEAD"
	PUT     = "PUT"
	DELETE  = "DELETE"
	PATCH   = "PATCH"
	OPTIONS = "OPTIONS"
)

type Request struct {
	Head           http.Header
	client         http.Client
	method         string
	queueURLString string
}

func (req *Request) Header(head map[string]string) *Request {
	for key, value := range head {
		req.Head.Add(key, value)
	}
	return req
}

func (req *Request) Timeout(time int) *Request {

	return req
}

func (req *Request) JSON() *Request {

	return req
}

func (req *Request) QueueParam(queueParam *map[string]string) *Request {
	var queueString bytes.Buffer
	if queueParam != nil && len(*queueParam) > 0 {
		queueString.WriteString("?")
		for key, value := range *queueParam {
			queueString.WriteString(key)
			queueString.WriteString("=")
			queueString.WriteString(value)
			queueString.WriteString("&")
		}
		s := queueString.String()
		req.queueURLString = s[0 : len(s)-1]
	}
	req.queueURLString = ""
	return req
}

////////////////////////////////////////////////
type RequestParam struct {
}

////////////////////////////////////////////////
func (req *Request) POST(url string) *Request {
	req.method = POST
	return req
}

func (req *Request) GET(url string) *Request {
	req.method = GET
	return req
}

func (req *Request) DELETE(url string) *Request {
	req.method = DELETE
	return req
}
func (req *Request) PUT(url string) *Request {
	req.method = PUT
	return req
}

func (req *Request) HEAD(url string) *Request {
	req.method = HEAD
	return req
}

func (req *Request) PATCH(url string) *Request {
	req.method = PATCH
	return req
}

func (req *Request) OPTIONS(url string) *Request {
	req.method = OPTIONS
	return req
}

func (req *Request) Send(ctx context.Context) {
	switch req.method {
	case POST:
		return s.Post(targetUrl)
	case GET:
		return s.Get(targetUrl)
	case HEAD:
		return s.Head(targetUrl)
	case PUT:
		return s.Put(targetUrl)
	case DELETE:
		return s.Delete(targetUrl)
	case PATCH:
		return s.Patch(targetUrl)
	case OPTIONS:
		return s.Options(targetUrl)
	default:
	}
}
