package users

import (
	"github.com/mdapathy/softeng/Lab2/server/forums"
	"net/http"
)

type HttpHandlerFunc http.HandlerFunc

func HttpHandler(s *Storage, sf *forums.Storage) HttpHandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {

			id, strings, err := handleCreateUser(s, r, rw)

			if err == nil {
				forums.HandleUpdateForums(sf, rw, id, strings)
			}

		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}

	}
}
