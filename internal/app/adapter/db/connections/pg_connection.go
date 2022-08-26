package connections

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

func Connect() *sql.DB {
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&TimeZone=Asia/Shanghai", os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("DATABASE_NAME"))
	result, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = result.Ping()
	if err != nil {
		panic(err)
	}
	return result
}
