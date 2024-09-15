package handles

import "net/http"

type Handlesfuncs interface{}

type Handles struct{}

func (Handles) Getbyall(http.ResponseWriter, *http.Request)
func (Handles) CreateUser(http.ResponseWriter, *http.Request)
func (Handles) Update(http.ResponseWriter, *http.Request)
