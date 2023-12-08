package main 

import (
	"fmt"
	"log"
	"os"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func init(){
	if err:= godotenv.Load("/app/.env"); err != nil{
		log.Fatal("No .env file found", err)
	}
}
var (
	str = os.Getenv("DBSTR")
)


func main(){
	// setting up db connection
	var connStr string
	connStr = str
	db, err := sql.Open("postgres", connStr)
	CheckError(err)

	defer db.Close()
	// setting up SQL Query
	getStmt := `select exists (
			select "mlra_name" 
			from public_test."geoIndicators" 
			where "mlra_name" is null
		)`
	// getting results from query
	rows, err := db.Query(getStmt)
	CheckError(err)

	defer rows.Close()
	// printing results with a forloop 
	for rows.Next(){
		var exists string 

		err = rows.Scan(&exists)
		CheckError(err)

		fmt.Printf("Does the mlra column contain nulls?: %s",exists)
	}
}

// error function to minimize error check repeats
func CheckError(err error) {
	if err != nil {
			panic(err)
	}
}