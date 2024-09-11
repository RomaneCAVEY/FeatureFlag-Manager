package infrastructure

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/RomaneCAVEY/FeatureFlag-Manager/domain/entities"
)

/*No sql injection possible because the user don't write in our application*/

const name_table string = "feature_flags"

var errorNoId = errors.New("no feature-flag with this id")
var errNoApplication = errors.New("no feature-flag with this application")
var errDuplicateKeyFeatureFlag = errors.New("Label already exists for this application. Find another label")


type FeatureFlagRepository struct {
	Collection *sql.DB
}


func (r *FeatureFlagRepository) Save(f entities.FeatureFlag) (entities.FeatureFlag, error) {
	var projects string = strings.Join(f.Projects, " , ")
	var owners string = strings.Join(f.Owners, " , ")

	query := fmt.Sprintf("INSERT INTO %s (slug, label, IsEnabled, application, projects, owners, description, createdat, updatedat,createdby,updatedby) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10,$11);", name_table)

	_, err := r.Collection.Exec(query, f.Slug, f.Label, f.IsEnabled, f.Application, projects, owners, f.Description, f.CreatedAt, f.CreatedAt, f.CreatedBy, f.UpdatedBy)
	if err != nil {
		log.Println(errDuplicateKeyFeatureFlag)
		return entities.FeatureFlag{}, errDuplicateKeyFeatureFlag
	}
	queryResponse := fmt.Sprintf("SELECT id, slug, label, isEnabled, application, projects, owners, description, createdat, updatedat, createdby,updatedby FROM %s  WHERE label = $1 AND application = $2 ;", name_table)
	rows, err := r.Collection.Query(queryResponse, f.Label, f.Application)
	if err != nil {
		errorSelect := errors.New("error with Db")
		log.Println(errorSelect)
		return entities.FeatureFlag{}, errorSelect

	}
	if rows.Next() {
		return rowsToFeatureFlag(rows)
	}
	return f, nil
}

func (r *FeatureFlagRepository) FindAll(start int, end int) ([]entities.FeatureFlag, int, error) {
	limit := end - start
	var flags = []entities.FeatureFlag{}
	query := fmt.Sprintf("SELECT id, slug, label, isEnabled, application, projects, owners, description, createdat, updatedat, createdby,updatedby FROM %s OFFSET $1 LIMIT $2;", name_table)
	rows, err := r.Collection.Query(query, start, limit)

	if err != nil {
		log.Println(err)
		return nil, 0, err
	}
	for rows.Next() {
		f, err := rowsToFeatureFlag(rows)
		if err != nil {
			log.Println(err)
			return flags, 0, err
		}
		flags = append(flags, f)
	}
	var count int
	query_count := fmt.Sprintf("SELECT COUNT (*) FROM %s ;", name_table)
	row_count, err := r.Collection.Query(query_count)
	if row_count.Next() {
		row_count.Scan(&count)
	}
	return flags, count, nil
}

func (r *FeatureFlagRepository) FindByApplication(application string, start int, end int) ([]entities.FeatureFlag, int, error) {
	var flags = []entities.FeatureFlag{}
	limit := end - start
	queryResponse := fmt.Sprintf("SELECT  id, slug, label, isEnabled, application, projects, owners, description, createdat, updatedat, createdby,updatedby FROM %s  WHERE application = $1 OFFSET $2 LIMIT $3;", name_table)
	rows, err := r.Collection.Query(queryResponse, application, start, limit)
	if err != nil {
		log.Println(err)
		return flags, 0, err
	}
	for rows.Next() {
		f, err := rowsToFeatureFlag(rows)
		if err != nil {
			log.Println(err)
			return flags, 0, err
		}
		flags = append(flags, f)
	}
	var count int
	query_count := fmt.Sprintf("SELECT COUNT (*) FROM %s WHERE application = $1 OFFSET $2 LIMIT $3;", name_table)
	row_count, err := r.Collection.Query(query_count, application, start, limit)
	if err != nil {
		log.Println(err)
		return flags, 0, err
	}
	if row_count.Next() {
		row_count.Scan(&count)
	}

	return flags, count, nil
}

func (r *FeatureFlagRepository) FindById(id uint32) (entities.FeatureFlag, error) {
	flag := entities.FeatureFlag{}
	queryResponse := fmt.Sprintf("SELECT id, slug, label, isEnabled, application, projects, owners, description, createdat, updatedat, createdby, updatedby FROM %s  WHERE id = $1 ;", name_table)
	rows, err := r.Collection.Query(queryResponse, id)
	if err != nil {
		log.Println(err)
		return flag, err
	}
	if rows.Next() {
		f, err := rowsToFeatureFlag(rows)
		if err != nil {
			log.Println(err)
			return flag, err
		}
		return f, nil
	}

	return flag, nil
}

func (r *FeatureFlagRepository) SaveChangesFeatureFlag(id uint32, label string, isEnabled bool, name string) (entities.FeatureFlag, error) {
	queryCheck := fmt.Sprintf("UPDATE %s SET label = $1, IsEnabled = $2, updatedat=$3, updatedby=$4 WHERE id = $5;", name_table)

	_, err := r.Collection.Query(queryCheck, label, isEnabled, time.Now(), name, id)

	if err != nil {
		log.Println(err)
		return entities.FeatureFlag{}, err
	}
	query := fmt.Sprintf("SELECT id, slug, label, isEnabled, application, projects, owners, description, createdat, updatedat,createdby, updatedby FROM  %s WHERE id = $1;", name_table)

	rows, error := r.Collection.Query(query, id)

	if error != nil {
		log.Println(err)
		return entities.FeatureFlag{}, error
	}
	if rows.Next() {
		return rowsToFeatureFlag(rows)
	}
	return entities.FeatureFlag{}, errorNoId
}

func (r *FeatureFlagRepository) RemoveFeatureFlag(id uint32) error {
	queryCheck := fmt.Sprintf("SELECT id FROM %s  WHERE id = $1;", name_table)

	rows, err := r.Collection.Query(queryCheck, id)

	if err != nil {
		log.Println(err)
		return err
	}
	if rows.Next() {
		queryCheck := fmt.Sprintf("DELETE FROM %s WHERE id = $1;", name_table)

		_, err := r.Collection.Query(queryCheck, id)
		if err != nil {
			log.Println(err)
			return err
		}
		return nil
	}
	return errorNoId

}

func rowsToFeatureFlag(rows *sql.Rows) (entities.FeatureFlag, error) {
	var f entities.FeatureFlag
	var projects string
	var owners string
	err := rows.Scan(&f.Id, &f.Slug, &f.Label, &f.IsEnabled, &f.Application, &projects, &owners, &f.Description, &f.CreatedAt, &f.UpdatedAt, &f.CreatedBy, &f.UpdatedBy)

	if err != nil {
		log.Println(err)
		return entities.FeatureFlag{}, err
	}
	f.Projects = strings.Split(projects, ",")
	f.Owners = strings.Split(owners, ",")
	return f, nil

}
