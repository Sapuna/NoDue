package Model

import (
	"webapp/Datastore/Postgres"
)

type Due struct {
	ID               int64
	EnrollmentNumber int
	Name             string
	Year             int
	Status           string
	Department       string
}

const queryInsertDue = "INSERT INTO due(enrollment_number, name, year, status, department) VALUES ($1, $2, $3, $4, $5) RETURNING id;"

func (due *Due) Create() error {
	err := Postgres.Db.QueryRow(queryInsertDue, due.EnrollmentNumber, due.Name, due.Year, due.Status, due.Department).Scan(&due.ID)
	return err
}

const queryGetAllDue = "SELECT * FROM due"

func GetAllDue() ([]Due, error) {
	rows, err := Postgres.Db.Query(queryGetAllDue)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	dueList := []Due{}
	for rows.Next() {
		var due Due

		err := rows.Scan(&due.ID, &due.EnrollmentNumber, &due.Name, &due.Year, &due.Status, &due.Department)
		if err != nil {
			return nil, err
		}

		dueList = append(dueList, due)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return dueList, nil
}

const queryUpdateDue = "UPDATE due SET status = $2 WHERE id = $1;"

func (due *Due) UpdateDue() error {
	_, err := Postgres.Db.Exec(queryUpdateDue, due.ID, due.Status)
	return err
}

const queryDeleteDue = "DELETE FROM due WHERE id = $1"

func (due *Due) DeleteDue() error {
	if _, err := Postgres.Db.Exec(queryDeleteDue, due.ID); err != nil {
		return err
	}
	return nil
}
