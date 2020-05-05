package Models

import (
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
	var arr = []Dev{}
	db := DB.ConnSql()
	defer db.Close()

	searchDevs := "SELECT * FROM devs"
	rows, err := db.Query(searchDevs)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for rows.Next() {
		newDev := Dev{}
		if err := rows.Scan(&newDev.ID,&newDev.Name,&newDev.Age); err != nil {
			log.Fatal(err)
			return nil, err
		}
		arr = append(arr,newDev)
	}

	err = rows.Close()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return arr, nil
}
