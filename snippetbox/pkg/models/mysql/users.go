package mysql

import (
	"database/sql"
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"mahmud139/snippetbox/pkg/models"

	"github.com/go-sql-driver/mysql"
)

type UserModel struct {
	DB *sql.DB
}

// We'll use the Insert method to add a new record to the users table.
func (m *UserModel) Insert(name, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmt := `INSERT INTO users (name, email, hashed_password, created)
	VALUES (?, ?, ?, UTC_TIMESTAMP())`

	_, err = m.DB.Exec(stmt, name, email, hashedPassword)
	if err != nil {
		// If this returns an error, we use the errors.As() function to check 
		// whether the error has the type *mysql.MySQLError. If it does, the 
		// error will be assigned to the mySQLError variable. We can then check 
		// whether or not the error relates to our users_uc_email key by 
		// checking the contents of the message string. If it does, we return 
		// an ErrDuplicateEmail error.
		var mySQLErr *mysql.MySQLError
		if errors.As(err, &mySQLErr) {
			if mySQLErr.Number == 1062 && strings.Contains(mySQLErr.Message, "users_uc_email") {
				return models.ErrDuplicateEmail
			}
		}
		return err
	}

	return nil
}

// We'll use the Authenticate method to verify whether a user exists with 
// the provided email address and password. This will return the relevant 
// user ID if they do.
func (m *UserModel) Authenticate(email, password string) (int, error) {
	// Retrieve the id and hashed password associated with the given email. If no 
	// matching email exists, or the user is not active, we return the 
	// ErrInvalidCredentials error.
	var id int
	var hashedPassword []byte
	stmt := `SELECT id, hashed_password FROM users WHERE email = ? AND active =TRUE`
	row := m.DB.QueryRow(stmt, email)
	err := row.Scan(&id, &hashedPassword)
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, models.ErrInvalidCredentials
		} else {
			return 0, err
		}
	}

	return id, nil
}

// We'll use the Get method to fetch details for a specific user based 
// on their user ID.
func (m *UserModel) Get(id int) (*models.User, error) {
	u := &models.User{}

	stmt := `SELECT id, name, email, created, active FROM users WHERE id = ?`
	err := m.DB.QueryRow(stmt, id).Scan(&u.ID, &u.Name, &u.Email, &u.Created, &u.Active)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return u, nil
}

