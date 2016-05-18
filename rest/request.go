package rest

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type Request struct {
	*http.Request

	PathParams map[string]string
}

func NewRequest(super *http.Request) *Request {
	return &Request{
		super,
		make(map[string]string),
	}
}

func (this Request) PathParam(name string) string {
	return this.PathParams[name]
}

func (this *Request) Deserialize(val interface{}) error {
	body, err := ioutil.ReadAll(this.Body)
	this.Body.Close()

	if err != nil {
		return err
	}

	if len(body) < 1 {
		return errors.New("No payload")
	}

	if err = json.Unmarshal(body, val); err != nil {
		return err
	}

	return nil
}
