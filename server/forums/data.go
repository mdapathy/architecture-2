package forums

import (
	"database/sql"
	"github.com/lib/pq"
	"strings"
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

	rows, err := s.Db.Query(
		"select a.id, a.name, a.topic_keyword, array_agg(a.users_name) as users " +
			" from ( select forums.id, forums.name, forums.topic_keyword, userforum.userid, " +
			"users.users_name  from forums left join userforum on forums.id = userforum.forumid  " +
			"left join users on userforum.userid = users.id order by forums.id) a group by a.id, a.name , " +
			"a.topic_keyword order by a.id;")
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

	str := "'" + strings.Join(topicKeyword, "' , '") + "'"
	println(str)
	rows, err := s.Db.Query("select id from forums where topic_keyword in (" + str + ")")
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {

		var forumid int
		if err = rows.Scan(&forumid); err != nil {
			return err
		}
		_, err = s.Db.Query("insert into userforum values ($1, $2)", forumid, id)

		if err != nil {
			return err
		}

	}


	return nil
}
