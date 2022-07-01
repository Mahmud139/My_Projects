package mysql

import (
	"database/sql"
	"mahmud139/snippetbox/pkg/models"

	//"projects/snippetbox/pkg/models"
)
//Define a SnippetModel type which warp a sql.DB connection pool
type SnippetModel struct {
	DB *sql.DB
}

//this will insert a new snippet into the database
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires)
	VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}


//this will return specific snippet based on its ID
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}


//this will return the 10 most recently created snippets.
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}