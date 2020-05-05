package Models

import (
	"squad-manager/Aux"
	"squad-manager/DB"
	"time"
)

type Deploy struct {
	ID   string `json:id`
	Requirements_id string `json:requirements_id`
	Date  time.Time    `json:date`
}

func InsertDeploy(deploy Deploy) error{
	db := DB.ConnSql()

	insertDeploy := `
INSERT INTO deploy (deploy_id, deploy_date, requirements_id)
VALUES ($1,$2,$3)
`
	deploy.ID = Aux.GenerateUUID()
	_, err := db.Exec(insertDeploy, deploy.ID, deploy.Date, deploy.Requirements_id)
	return err
}