package users

import (
	"github.com/mdapathy/architecture-2/server/forums"
	"net/http"
)

type Answer struct {
	Id int
	Interests []string
}

type HttpHandlerFunc http.HandlerFunc

func HttpHandler(s *Storage, sf *forums.Storage) HttpHandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {

			var answer Answer;
			var err error;
			answer.Id, answer.Interests, err = handleCreateUser(s, r, rw)

			if err == nil {
				forums.HandleUpdateForums(sf, rw, answer.Id, answer.Interests)
			}

		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}

	}
}
