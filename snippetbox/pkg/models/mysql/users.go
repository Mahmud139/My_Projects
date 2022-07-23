package mysql

import (
	"database/sql"

	"mahmud139/snippetbox/pkg/models"
)

type UserModel struct {
	DB *sql.DB
}

// We'll use the Insert method to add a new record to the users table.
func (u *UserModel) Insert(name, email, password string) error {
	return nil
}

// We'll use the Authenticate method to verify whether a user exists with 
// the provided email address and password. This will return the relevant 
// user ID if they do.
func (u *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

// We'll use the Get method to fetch details for a specific user based 
// on their user ID.
func (u *UserModel) Get(id int) (*models.User, error) {
	return nil, nil
}

