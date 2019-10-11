package forums

import (
	"github.com/mdapathy/softeng/Lab2/server/tools"
	"log"
	"net/http"
)

type HttpHandlerFunc http.HandlerFunc

func HttpHandler(s *Storage) HttpHandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handleListForums(s, rw)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func handleListForums(s *Storage, rw http.ResponseWriter) {

	res, err := s.ListForums()
	if err != nil {
		log.Printf("Error making query to the db: %s", err)
		tools.WriteJsonInternalError(rw)
		return
	}

	tools.WriteJsonOk(rw, res)

}

func HandleUpdateForums(s *Storage, rw http.ResponseWriter, id int, topics []string) {

	err := s.UpdateForums(id, topics)

	if err != nil {
		log.Printf("Error making query to the db: %s", err)
		tools.WriteJsonInternalError(rw)
		return
	}

}
