package Models

import (
	"database/sql"
	"log"
	"squad-manager/Aux"
	"squad-manager/DB"
)

type Dev struct {
	ID   string `json:id`
	Name string `json:name`
	Age  int    `json:age`
}

func InsertDev(dev Dev) error {
	db := DB.ConnSql()
	defer db.Close()

	insertDeveloper := `
INSERT INTO devs (dev_id,name, age)
VALUES ($1, $2, $3)
`

	dev.ID = Aux.GenerateUUID()
	_, err := db.Exec(insertDeveloper, dev.ID, dev.Name, dev.Age)
	return err
}

func SearchAllDevs() ([]Dev,error) {
	db := DB.ConnSql()
	defer db.Close()

	searchDevs := "SELECT * FROM devs"
	err := db.QueryRow(searchDevs)
	if err != nil {
		log.Fatal(err)
	}

	if err == sql.ErrNoRows {
		log.Fatal("No results found")
	}

}
