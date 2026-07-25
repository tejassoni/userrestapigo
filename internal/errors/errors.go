package errors

import (
	"errors"

	"github.com/go-sql-driver/mysql"
)

// IsDuplicateEmailError reports whether err is the MySQL unique-key violation
// raised for the users.email constraint.
func IsDuplicateEmailError(err error) bool {
	var mysqlErr *mysql.MySQLError
	return errors.As(err, &mysqlErr) && mysqlErr.Number == 1062
}
