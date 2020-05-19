package Aux

import "github.com/google/uuid"

func GenerateUUID() (string) {
	id, err := uuid.NewUUID()
	if err != nil {
		panic(error(err))
	}
	return id.String()
}