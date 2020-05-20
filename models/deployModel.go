package models

import (
	"squad-manager/db"
	"time"
)

// Deploy is a struct to represent a deploy
type Deploy struct {
	ID             string    `json:"id"`
	RequirementsID string    `json:"requirements_id"`
	Date           time.Time `json:"date"`
}

// InsertDeploy is a function to insert a Deploy on the database
func InsertDeploy(deploy Deploy) error {
	db := db.ConnSQL()

	insertDeploy := `
INSERT INTO deploy (deploy_id, deploy_date, requirements_id)
VALUES ($1,$2,$3)
`
	deploy.ID = utilsFunc.GenerateUUID()
	_, err := db.Exec(insertDeploy, deploy.ID, deploy.Date, deploy.RequirementsID)
	return err
}
