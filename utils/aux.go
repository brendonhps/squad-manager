package utils

import "github.com/google/uuid"

// GenerateUUID is a function to generate a new UUID
func GenerateUUID() string {
	id, err := uuid.NewUUID()
	if err != nil {
		panic(error(err))
	}
	return id.String()
}
