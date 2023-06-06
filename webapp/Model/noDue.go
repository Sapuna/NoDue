package Model

import (
	"webapp/Datastore/Postgres"
)

type NoDue struct {
	ID               int64
	EnrollmentNumber int
	Name             string
	Year             int
	Status           string
}

const queryInsertNoDue = "INSERT INTO nodue(enrollment_number, name, year, status) VALUES ($1, $2, $3, $4) RETURNING id;"

func (due *NoDue) Create() error {
	err := Postgres.Db.QueryRow(queryInsertNoDue, due.EnrollmentNumber, due.Name, due.Year, due.Status).Scan(&due.ID)
	return err
}

const queryGetAllNoDue = "SELECT * FROM nodue"

func GetAllNoDue() ([]NoDue, error) {
	rows, err := Postgres.Db.Query(queryGetAllNoDue)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	dueList := []NoDue{}
	for rows.Next() {
		var nodue NoDue

		err := rows.Scan(&nodue.ID, &nodue.EnrollmentNumber, &nodue.Name, &nodue.Year, &nodue.Status)
		if err != nil {
			return nil, err
		}

		dueList = append(dueList, nodue)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return dueList, nil
}

const queryDeleteNoDue = "DELETE FROM nodue WHERE id = $1"

func (due *NoDue) DeleteNoDue() error {
	if _, err := Postgres.Db.Exec(queryDeleteNoDue, due.ID); err != nil {
		return err
	}
	return nil
}
