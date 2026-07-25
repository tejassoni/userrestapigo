package repository

import (
	"errors"
	"userrestapigo/internal/config"
	"userrestapigo/internal/models"
)

/*
* CreateUser inserts a new user into the database.
* It takes a User struct as input and returns an error if the operation fails.
@param user models.User - The user data to be inserted into the database.
@return error - Returns an error if the database operation fails, otherwise returns nil.
*/
func CreateUser(user models.User) (int, error) {
	// prepare the SQL query for inserting a new user into the users table
	query := `
		INSERT INTO users
		(name, email, is_active, birthdate, gender, password)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	// execute the SQL query with the provided user data
	result, err := config.DB.Exec(
		query,
		user.Name,
		user.Email,
		user.IsActive,
		user.Birthdate,
		user.Gender,
		user.Password,
	)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	user.ID = int(lastInsertID)
	return user.ID, nil
}

/*
* GetUsers retrieves all users from the database.
* It returns a slice of User structs and an error if the operation fails.
@return []models.User - A slice containing all users retrieved from the database.
@return error - Returns an error if the database operation fails, otherwise returns nil.
*/
func GetUsers() ([]models.User, error) {
	// prepare the SQL query for selecting all users from the users table
	query := `
		SELECT id, name, email, is_active, birthdate, gender, created_at, updated_at
		FROM users
		ORDER BY created_at DESC, id DESC
	`
	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.IsActive, &user.Birthdate, &user.Gender, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	// check for any errors encountered during iteration
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

/*
* GetUserByID retrieves a user from the database by their ID.
* It takes an integer ID as input and returns a User struct and an error if the operation fails.
@param id int - The ID of the user to be retrieved from the database.
@return models.User - The user data retrieved from the database.
@return error - Returns an error if the database operation fails, otherwise returns nil.
*/
func GetUserByID(id int) (models.User, error) {
	var user models.User
	query := `
		SELECT id, name, email, is_active, birthdate, gender, created_at, updated_at
		FROM users
		WHERE id = ?
	`
	err := config.DB.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.IsActive, &user.Birthdate, &user.Gender, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return user, err
	}
	return user, nil
}

/*
* UpdateUser updates an existing user in the database.
* It takes a User struct as input and returns an error if the operation fails.
@param user models.User - The user data to be updated in the database.
@return error - Returns an error if the database operation fails, otherwise returns nil.
*/
func UpdateUser(user models.User) error {
	// Validate input
	if user.ID <= 0 || user.Email == "" {
		return errors.New("invalid user data")
	}

	query := `
		UPDATE users
		SET name = ?, email = ?, is_active = ?, birthdate = ?, gender = ?, password = ?, updated_at = NOW()
		WHERE id = ?
	`
	result, err := config.DB.Exec(
		query,
		user.Name,
		user.Email,
		user.IsActive,
		user.Birthdate,
		user.Gender,
		user.Password, // should be hashed
		user.ID,
	)
	if err != nil {
		return err
	}

	// Verify row was actually updated
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("user not found")
	}

	return nil
}

/*
* DeleteUser removes a user from the database by their ID.
* It takes an integer ID as input and returns an error if the operation fails.
@param id int - The ID of the user to be deleted from the database.
@return bool - Returns true when a user was deleted, false when the ID does not exist.
@return error - Returns an error if the database operation fails, otherwise returns nil.
*/
func DeleteUser(id int) (bool, error) {
	query := `
		DELETE FROM users
		WHERE id = ?
	`
	result, err := config.DB.Exec(query, id)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}

/*
* UserEmailExists checks if a user with the given email already exists in the database.
* It takes a string email as input and returns a boolean indicating existence and an error if the operation fails.
@param email string - The email address to check for existence in the database.
@return bool - Returns true if the email exists, false otherwise.
@return error - Returns an error if the database operation fails, otherwise returns nil.
*/
func UserEmailExists(email string) (bool, error) {
	var count int

	query := `SELECT COUNT(*) FROM users WHERE email = ?`

	err := config.DB.QueryRow(query, email).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

/*
* UserIDExists checks if a user with the given ID already exists in the database.
* It takes an integer ID as input and returns a boolean indicating existence and an error if the operation fails.
@param id int - The ID of the user to check for existence in the database.
@return bool - Returns true if the user exists, false otherwise.
@return error - Returns an error if the database operation fails, otherwise returns nil.
*/
func UserIDExists(id int) (bool, error) {
	var count int

	query := `SELECT COUNT(*) FROM users WHERE id = ?`

	err := config.DB.QueryRow(query, id).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
