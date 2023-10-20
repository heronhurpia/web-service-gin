package server

import (
	"database/sql"
	//"errors"
	"fmt"
	//"log"
	//"os"
	//"path/filepath"

	_ "github.com/lib/pq"
	//"encoding/json"
	//"github.com/heronhurpia/web-service-gin/helper" 
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
	rows, err := db.Query(`SELECT "id","description" FROM "sat"."comments" limit 20`)
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

	// for rows.Next() {
	// // 	//var name_en string
	//  	var description string
	//  	var id int
	//  	err = rows.Scan(&id, &description)
	//  	CheckError(err)

	// 	dat =: helper.Comment{
	// 	 	Id : id
	// 	 	Description : description
	// 	}
	// 	fmt.Print(dat)
	// // 	data := fmt.Sprintf("%d. %s\n", id, description)
	// // 	fmt.Print(data)
	// // 	//_, err := file.WriteString(data)
	// // 	CheckError(err)
	// }

	//fmt.Printf("t1: %T\n", rows)
	//resp, err := json.Marshal(rows)
	CheckError(err)

	var logs = []log{
		{ID: "1", Description: "Blue Train"},
		{ID: "2", Description: "Jeru"},
		{ID: "3", Description: "Sarah Vaughan and Clifford Brown"},
	}
	
	return logs ;
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}