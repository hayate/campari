package main

import (
	"net/http"
	"reflect"
	"strings"

	"github.com/hayate/campari/app/controllers"
	"github.com/hayate/campari/app/helpers"
)

func handler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")

	segs := strings.Split(strings.Trim(req.URL.Path, "/\\"), "/")
	if len(segs) < 2 {
		http.NotFound(res, req)
		return
	}
	params := []reflect.Value{}

	if len(segs) > 2 {
		for _, param := range segs[2:3] {
			params = append(params, reflect.ValueOf(param))
		}
	} else {
		params = append(params, reflect.ValueOf("Campari"))
	}

	switch segs[0] {
	case "welcome":
		controller := controllers.Welcome{res, req}
		action := helpers.Capitalize(segs[1])
		method := reflect.ValueOf(&controller).MethodByName(action)
		if method.IsValid() {
			method.Call(params)
		} else {
			http.NotFound(res, req)
		}
		break
	default:
		http.NotFound(res, req)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
