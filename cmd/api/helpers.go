package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

func (app *application) readIDParams(r *http.Request) (int64, error) {
	/*
		When httprouter is parsing a request, any interpolated URL parameters will be stored in the context.
		We can use the ParamsFromContext() function to retrieve a slice containing these parameter names and values.
	*/
	param := r.PathValue("id")

	/*
		We can the use the ByName() method to get the value of the "id" parameter from the slice.
		All movies will have a unique positive integer ID, but the value returned by ByName() is always a string.
		So we try to convert it to a base 10 integer (with a bit size of 64).
	*/
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, headers http.Header) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	js = append(js, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}
