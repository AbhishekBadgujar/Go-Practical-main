package main

import (
	"database/sql"
	"log"
	"os"

	routes "github.com/AbhishekBadgujar/Go-Practical_main/routes"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func init() {
	if _, err := os.Stat(".env"); err == nil { //os.Stat() if returns nil checks the existence of a file.
		log.Println("Loading the config file from .env file")
		err = godotenv.Load(".env") //Load environment variables
		if err != nil {
			log.Println("Error loading the config file")
		}
		log.Println("Successfully loaded the config file")
	}
	ConnectDb() //Connect to DB
}

func ConnectDb() {
	connectionUrl := os.Getenv("DBUsername") + ":" + os.Getenv("DBPassword") + "@tcp(127.0.0.1:3306)/" + os.Getenv("DBName") //storing sensitive DB info in .env file
	//var connectionUrl = "root:1234@tcp(127.0.0.1:3306)/userinfo" actual DB url
	Db, err := sql.Open("mysql", connectionUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer Db.Close()
	err = Db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to the database!")
}

func main() {
	//fmt.Println("This is my practical exam")
	routes.ClientRoutes()
}

//Ideally should use Makefile for testing and building application, but not able to do so due to time constraint
