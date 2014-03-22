package controllers

import (
	"fmt"
	"net/http"
)

type Welcome struct {
	Res http.ResponseWriter
	Req *http.Request
}

func (a *Welcome) Home(name string) {
	fmt.Fprintf(a.Res, "Welcome %s!\n", name)
}
