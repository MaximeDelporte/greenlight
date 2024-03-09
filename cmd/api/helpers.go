package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) readIDParams(r *http.Request) (int64, error) {
	/*
		When httprouter is parsing a request, any interpolated URL parameters will be stored in the context.
		We can use the ParamsFromContext() function to retrieve a slice containing these parameter names and values.
	*/
	params := httprouter.ParamsFromContext(r.Context())

	/*
		We can the use the ByName() method to get the value of the "id" parameter from the slice.
		All movies will have a unique positive integer ID, but the value returned by ByName() is always a string.
		So we try to convert it to a base 10 integer (with a bit size of 64).
	*/
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil
}
