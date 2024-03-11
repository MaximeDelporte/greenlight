package main

/*
	Because I'm not using httprouter library, I couldn't implement custom error message that we can found page 71.
	It's not the best way to do but, at least, it's working.
	Solution here: https://stackoverflow.com/questions/35528330/overriding-responsewriter-interface-to-catch-http-errors
*/

import (
	"encoding/json"
	"fmt"
	"log"
)
import "net/http"

type Interceptor struct {
	writer     http.ResponseWriter
	request    *http.Request
	overridden bool
}

func (i *Interceptor) WriteHeader(statusCode int) {
	switch statusCode {
	case 500:
		writeJSON("Custom 500 message / content", statusCode, i)
	case 403:
		writeJSON("You can't access this resource.", statusCode, i)
	case 404:
		writeJSON("The requested resource could not be found.", statusCode, i)
	case 405:
		message := fmt.Sprintf("The %s method is not supported for this resource.", i.request.Method)
		writeJSON(message, statusCode, i)
	default:
		i.writer.WriteHeader(statusCode)
		return
	}

	// if the default case didn't execute (and return) we must have overridden the output
	i.overridden = true
	log.Println(i.overridden)
}

func writeJSON(message string, statusCode int, i *Interceptor) {
	js, err := json.Marshal(message)
	if err != nil {
		return
	}

	i.writer.Header().Set("Content-Type", "application/json")
	i.writer.WriteHeader(statusCode)
	i.writer.Write(js)
}

func (i *Interceptor) Write(b []byte) (int, error) {
	if !i.overridden {
		return i.writer.Write(b)
	}

	// Return nothing if we've overriden the response.
	return 0, nil
}

func (i *Interceptor) Header() http.Header {
	return i.writer.Header()
}
