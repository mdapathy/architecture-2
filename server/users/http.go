package users

import (
	"encoding/json"
	"github.com/mdapathy/architecture-2/server/tools"
	"log"
	"net/http"
)

func handleCreateUser(s *Storage, r *http.Request, rw http.ResponseWriter) (int, []string, error) {

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return 0, nil, err
	}
	id, err := s.CreateUser(user.Name, user.Interests)

	if err == nil {
		tools.WriteJsonOk(rw, &user)
		return id, user.Interests, err
	} else {
		log.Printf("Error inserting record: %s", err)
		tools.WriteJsonInternalError(rw)
		return 0, nil, err
	}

}
