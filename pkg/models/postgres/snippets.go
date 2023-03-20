package postgres

import (
	"database/sql"

	"github.com/sahidrahman404/snippetbox/pkg/models"
)

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	stmt := `insert into snippets (title, content, expires)
	values($1, $2, current_date + cast($3 as interval)) returning id`

	var id int
	err := m.DB.QueryRow(stmt, title, content, expires).Scan(&id)
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	// stmt := `select * from snippets where id = ?`

	return nil, nil
}

func (m *SnippetModel) Latest(id int) ([]*models.Snippet, error) {
	return nil, nil
}
