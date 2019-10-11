package users

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
)

type Storage struct {
	Db *sql.DB
}

func NewStore(db *sql.DB) *Storage {
	return &Storage{Db: db}
}

type User struct {
	Name      string   `json:"name"`
	Interests []string `json:"interests"`
}

func (s *Storage) CreateUser(Name string, Interests []string) (int, error) {
	if len(Name) < 0 || len(Interests) < 0 {
		return 0, fmt.Errorf("users info is not provided")
	}

	res, err := s.Db.Query("INSERT INTO users (users_name, users_interests) VALUES ($1, $2) RETURNING ID;", Name, pq.Array(Interests))

	var id int
	if err == nil && res.Next() {
		scanErr := res.Scan(&id)
		if scanErr != nil {
			return 0, scanErr
		}
	}
	return id, err
}
