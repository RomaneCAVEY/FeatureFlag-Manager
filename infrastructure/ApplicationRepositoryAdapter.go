package infrastructure

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/RomaneCAVEY/FeatureFlag-Manager/domain/entities"
)

type ApplicationRepository struct {
	Collection *sql.DB
}

var errDuplicationKey = errors.New("Label already used")
var errNoId = errors.New("no Application with this id")
var errNoLabel = errors.New("no Application with this label")

const name_table_Application string = "Applications"

func (r *ApplicationRepository) Save(s entities.Application) (entities.Application, error) {
	query := fmt.Sprintf("INSERT INTO %s (label,description, createdat, updatedat) VALUES ($1, $2, $3, $4);", name_table_Application)

	_, err := r.Collection.Exec(query, s.Label, s.Description, s.CreatedAt, s.CreatedAt)
	if err != nil {
		log.Println(errDuplicationKey)
		return entities.Application{}, errDuplicationKey
	}
	queryResponse := fmt.Sprintf("SELECT * FROM %s  WHERE label = $1 AND description = $2 ;", name_table_Application)

	rows, err := r.Collection.Query(queryResponse, s.Label, s.Description)
	if err != nil {
		log.Println(err)
		return entities.Application{}, err
	}
	if rows.Next() {
		return rowsToApplication(rows)
	}
	return entities.Application{}, nil
}

func (r *ApplicationRepository) FindAll(start int, end int) ([]entities.Application, int, error) {
	limit := end - start
	query := fmt.Sprintf("SELECT * FROM %s OFFSET $1 LIMIT $2;", name_table_Application)
	rows, err := r.Collection.Query(query, start, limit)

	var listApplications = []entities.Application{}
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}
	for rows.Next() {
		s, err := rowsToApplication(rows)
		if err != nil {
			log.Println(err)
			return listApplications, 0, err
		}
		listApplications = append(listApplications, s)
	}
	var count int
	query_count := fmt.Sprintf("SELECT COUNT (*) FROM %s;", name_table_Application)
	row_count, err := r.Collection.Query(query_count)
	if row_count.Next() {
		row_count.Scan(&count)
	}
	return listApplications, count, nil
}

func (r *ApplicationRepository) FindById(id uint32) (entities.Application, error) {
	queryResponse := fmt.Sprintf("SELECT * FROM %s  WHERE id = $1 ;", name_table_Application)
	rows, err := r.Collection.Query(queryResponse, id)
	if err != nil {
		log.Println(err)
		return entities.Application{}, err
	}
	if rows.Next() {
		return rowsToApplication(rows)
	}
	return entities.Application{}, errNoId
}

func (r *ApplicationRepository) FindByLabel(label string) (entities.Application, error) {
	queryResponse := fmt.Sprintf("SELECT * FROM %s  WHERE label = $1 ;", name_table_Application)
	rows, err := r.Collection.Query(queryResponse, label)
	if err != nil {
		log.Println(err)
		return entities.Application{}, err
	}
	if rows.Next() {
		return rowsToApplication(rows)
	}
	return entities.Application{}, errNoLabel
}

func (r *ApplicationRepository) UpdateApplication(id uint32, label string, description string) (entities.Application, error) {
	queryCheck := fmt.Sprintf("UPDATE %s  SET label = $1, description = $2 , updatedat= $3 WHERE id = $4;", name_table_Application)
	rows, err := r.Collection.Query(queryCheck, label, description, time.Now(), id)
	if err != nil {
		log.Println(err)
		return entities.Application{}, err
	}
	queryResponse := fmt.Sprintf("SELECT * FROM %s  WHERE id = $1;", name_table_Application)
	rows, error := r.Collection.Query(queryResponse, id)

	if error != nil {
		log.Println(err)
		return entities.Application{}, error
	}
	if rows.Next() {
		return rowsToApplication(rows)
	}
	return entities.Application{}, errNoId
}

func rowsToApplication(rows *sql.Rows) (entities.Application, error) {
	var s entities.Application
	err := rows.Scan(&s.Id, &s.Label, &s.Description, &s.CreatedAt, &s.UpdatedAt)
	if err != nil {
		log.Println(err)
		return entities.Application{}, err
	}
	return s, nil

}

func (r *ApplicationRepository) RemoveApplication(id uint32) error {
	queryCheck := fmt.Sprintf("SELECT * FROM %s  WHERE id = $1;", name_table_Application)
	rows, error := r.Collection.Query(queryCheck, id)
	if error != nil {
		log.Println(error)
		return error
	}
	if rows.Next() {
		queryCheck := fmt.Sprintf("DELETE FROM %s  WHERE id = $1;", name_table_Application)
		_, err := r.Collection.Query(queryCheck, id)
		if err != nil {
			log.Println(err)
			return err
		}
		return nil
	}
	return errNoId

}
