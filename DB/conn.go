package DB

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func loadDotEnv(key string) string {
	err := godotenv.Load(".env")

	if err != nil{
		panic(error(err))
	}
	return os.Getenv(key)
}

func ConnSql() *sql.DB {
	host := loadDotEnv("HOST")
	port := loadDotEnv("PORT")
	user := loadDotEnv("USER")
	password := loadDotEnv("PASSWORD")
	dbname := loadDotEnv("DBNAME")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(error(err))
	}
	return db
}