package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"time"
)

var (
	db *sqlx.DB
)

type Image struct {
	Name string
	Data []byte
}

func main() {

	db_host := "db00"
	db_port := "3306"
	db_user := "isucon"
	db_password := ":isucon"

	dsn := fmt.Sprintf("%s%s@tcp(%s:%s)/isubata?parseTime=true&loc=Local&charset=utf8mb4",
		db_user, db_password, db_host, db_port)

	log.Printf("Connecting to db: %q", dsn)
	db, connectErr := sqlx.Connect("mysql", dsn)
	log.Println(connectErr)
	for {
		err := db.Ping()
		if err == nil {
			break
		}
		log.Println("error = ", err)
		time.Sleep(time.Second * 3)
	}

	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(5 * time.Minute)
	log.Printf("Succeeded to connect db.")

	images := []Image{}

	err := db.Select(&images, "select name, data from image group by name, data;")

	if err != nil {
		log.Println(err)
		return
	}

	image_dir := "/home/isucon/isubata/webapp/public/icons"
	if _, err := os.Stat(image_dir); os.IsNotExist(err) {
		err = os.Mkdir(image_dir, 0775)
	}

	if err != nil {
		log.Println(err)
		return
	}

	for i := 0; i < len(images); i++ {

		file, err := os.Create(fmt.Sprintf("%s/%s", image_dir, images[i].Name))

		if err != nil {
			fmt.Println(err)
		}

		defer file.Close()

		file.Write(images[i].Data)
	}

	// select min(id) as id, name, data from image group by name, data having count(*) > 1;

}
