package Models

import (
	"squad-manager/Aux"
	"squad-manager/DB"
)

type Dev struct {
	id   string `json:id`
	Name string `json:name`
	Age  int    `json:age`
}

func InsertDev(dev Dev) error {
	db := DB.ConnSql()
	defer db.Close()

	insertDeveloper := `
	INSERT INTO devs (dev_id,name, age)
	VALUES ($1, $2, $3)`

	dev.id = Aux.GenerateUUID()
	_, err := db.Exec(insertDeveloper, dev.id, dev.Name, dev.Age)
	return err
}

//func SearchAllDevs() (devs []Dev, err error) {
//	db := DB.ConnSql()
//	defer db.Close()
//
//
//}
