package server

import (
	"database/sql"
	//"errors"
	"fmt"
	//"log"
	//"os"
	//"path/filepath"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	// user     = "postgres"
	// password = "password"
	// dbname   = "cardapio_db"
	password = "U3t8QYg"
	user     = "sat"
	dbname   = "portal"
)

func GetAllChannels() []log {

	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")

//	rows, err := db.Query(`SELECT "id","name_en" FROM "public"."dishes" limit 5`)
	rows, err := db.Query(`SELECT "id","description","commentable_type","commentable_id","updated_at","created_at" FROM "sat"."comments" order by id desc limit 5`)
	CheckError(err)
	defer rows.Close()

	/*
		// Cria subdiretório
		subdir := "sample"
		if _, err := os.Stat(subdir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(subdir, os.ModePerm)
			if err != nil {
				log.Println(err)
			}
		}

		// Cria arquivo destino dos dados
		cwd, _ := os.Getwd()
		path := filepath.Join(cwd, subdir, "test-file.txt")
		newFilePath := filepath.FromSlash(path)
		file, err := os.Create(newFilePath)
		CheckError(err)
		defer file.Close()
	*/

	fmt.Println("rows scan")
	var logs []log
	for rows.Next() {

		var id int
		var description string
		var commentable_type string
		var commentable_id	int
		var updated_at	string
		var created_at	string

		err = rows.Scan(&id, &description,&commentable_type,&commentable_id,&updated_at,&created_at)
	  	CheckError(err)

		var l log 
		l.ID = id 
		l.Description = description
		l.Commentable_type = commentable_type
		l.Commentable_id = commentable_id
		l.Created_at = created_at
		logs = append(logs,l)
	}

	return logs ;
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}