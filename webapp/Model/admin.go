package Model

import (
	"database/sql"
	"webapp/Datastore/Postgres"
)

type Admin struct {
	ID         int64
	Name       string
	Email      string
	UserID     int
	Course     string
	Year       int
	Semester   string
	Department string
	Password   string
	Type       string
}

const queryInsertAdmin = "INSERT INTO users(username, email, userID, course, year, semester, departmentName, password, type) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9);"

func (adm *Admin) Create() error {
	_, err := Postgres.Db.Exec(queryInsertAdmin, adm.Name, adm.Email, adm.UserID, adm.Course, adm.Year, adm.Semester, adm.Department, adm.Password, adm.Type)
	return err
}

const queryGetAdmin = "SELECT id, username, email, userID, course, year, semester, departmentName,type FROM users WHERE userID = $1 AND password = $2"

type UnauthorizedError struct {
	Message string
}

func (e *UnauthorizedError) Error() string {
	return e.Message
}

var ErrUnauthorized = &UnauthorizedError{
	Message: "Unauthorized",
}

func (adm *Admin) Get() error {
	err := Postgres.Db.QueryRow(queryGetAdmin, adm.UserID, adm.Password).Scan(&adm.ID, &adm.Name, &adm.Email, &adm.UserID, &adm.Course, &adm.Year, &adm.Semester, &adm.Department, &adm.Type)
	if err == sql.ErrNoRows {
		return ErrUnauthorized
	}
	return err
}

const queryGetAllusers = "Select * from users"

func GetAllUsers() ([]Admin, error) {
	rows, getErr := Postgres.Db.Query(queryGetAllusers)
	if getErr != nil {
		return nil, getErr
	}
	users := []Admin{}
	for rows.Next() {
		var admin Admin

		dbErr := rows.Scan(&admin.ID, &admin.Name, &admin.Email, &admin.UserID, &admin.Course, &admin.Year, &admin.Semester, &admin.Department, &admin.Password, &admin.Type)
		if dbErr != nil {
			return nil, dbErr
		}

		users = append(users, admin)
	}
	rows.Close()
	return users, nil
}

const queryUpdateUser = "UPDATE users SET username = $2, email = $3, userID = $4, course = $5, year = $6, semester = $7, departmentName = $8, password = $9, type = $10 WHERE id = $1;"

func (user *Admin) UpdateUser(id int64) error {
	_, err := Postgres.Db.Exec(queryUpdateUser, id, user.Name, user.Email, user.UserID, user.Course, user.Year, user.Semester, user.Department, user.Password, user.Type)
	return err
}

const queryGetUser = "SELECT id, username, email, userID, course, year, semester, departmentName, password, type FROM users WHERE id = $1"

func (user *Admin) ReadUser() error {
	// Execute the database query and scan the result into the user struct fields
	return Postgres.Db.QueryRow(queryGetUser, user.ID).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.UserID,
		&user.Course,
		&user.Year,
		&user.Semester,
		&user.Department,
		&user.Password,
		&user.Type,
	)
}

const queryDeleteUser = "DELETE FROM users WHERE id = $1"

func (due *Admin) DeleteUser() error {
	if _, err := Postgres.Db.Exec(queryDeleteUser, due.ID); err != nil {
		return err
	}
	return nil
}

// Forgot password
const queryUpdatepassword = "UPDATE users SET password = $1 WHERE id = $2"

func (user *Admin) UpdatePassword(id int64) error {
	_, err := Postgres.Db.Exec(queryUpdatepassword, user.Password, id)
	return err
}
