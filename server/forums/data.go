package forums

import (
	"database/sql"
	"github.com/lib/pq"
)

type Storage struct {
	Db *sql.DB
}

func NewStore(db *sql.DB) *Storage {
	return &Storage{Db: db}
}

type Forum struct {
	Name    string   `json:"name"`
	Keyword string   `json:"topic_keyword"`
	Users   []string `json:"users"`
}

func (s *Storage) ListForums() ([]*Forum, error) {

	rows, err := s.Db.Query("select forums.id, forums.name, forums.topic_keyword," +
		" array_agg(users.users_name) from forums left join users " +
		"on users.id = any(users) group by forums.id limit 200;")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var forums []*Forum

	for rows.Next() {

		var i int
		var f Forum

		if err = rows.Scan(&i, &f.Name, &f.Keyword, pq.Array(&f.Users)); err != nil {
			return nil, err
		}

		forums = append(forums, &f)
	}
	if forums == nil {
		forums = make([]*Forum, 0)
	}

	return forums, nil
}

func (s *Storage) UpdateForums(id int, topicKeyword []string) error {

	for i := 0; i < len(topicKeyword); i++ {
		_, err := s.Db.Exec("UPDATE forums set users = array_append(users, $2) WHERE  topic_keyword= $1;", topicKeyword[i], id)
		if err != nil {
			return err
		}
	}

	return nil
}
