package rest

import "net/http"

type ResponseWriter struct {
	http.ResponseWriter
}

func (this *ResponseWriter) WriteHeader(code int) {
	this.Header().Set("Content-Type", "application/json; charset=utf-8")
	this.ResponseWriter.WriteHeader(code)
}

func (this *ResponseWriter) Write(val interface{}) error {
	json, err := Serialize(val)
	if err != nil {
		return err
	}

	_, err = this.write(json)
	if err != nil {
		return err
	}

	return nil
}

func (this *ResponseWriter) write(bytes []byte) (int, error) {
	this.WriteHeader(http.StatusOK)
	return this.ResponseWriter.Write(bytes)
}
