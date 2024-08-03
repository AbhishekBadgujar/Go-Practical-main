package routes

import (
	"database/sql"
	"net/http"

	model "github.com/AbhishekBadgujar/Go-Practical_main/model"
	"github.com/gin-gonic/gin"
)

var connectionUrl = "root:1234@tcp(127.0.0.1:3306)/userinfo" //Can store credentials in .env file for further readability,modification and data security
var Db *sql.DB

//Task 1 - Get Person Details from Id passed in Template Parameter

func getPerson(c *gin.Context) {
	CORSMiddleware(c)
	id := c.Param("person_id")
	var person model.Person
	Db, _ := sql.Open("mysql", connectionUrl)
	err := Db.QueryRow("SELECT id, Name, age FROM Person  WHERE id = ?", id).Scan(&person.Id, &person.Name, &person.Age)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"message": "Person is not found!"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error in getting person details"})
		}
		return
	}
	c.JSON(http.StatusOK, person)
}

// Task2 - Create POST API
func createPerson(c *gin.Context) {
	CORSMiddleware(c)
	var req model.CreatePersonRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	Db, _ := sql.Open("mysql", connectionUrl) //Connecting to DB
	tx, err := Db.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "DB connect error"})
		return
	}

	res, _ := tx.Exec("INSERT INTO Person (name, age) VALUES (?, ?)", req.Name, req.Age) // Inserting for person table from payload

	personID, _ := res.LastInsertId()

	_, _ = tx.Exec("INSERT INTO Phone (person_id, number) VALUES (?, ?)", personID, req.Number) // Inserting for Phone table from payload

	_, _ = tx.Exec("INSERT INTO Address ( city, state, street1, street2, zip_code) VALUES (?, ?, ?, ?, ?)",
		req.City, req.State, req.Street1, req.Street2, req.ZipCode) // Inserting for Address table from payload

	err = tx.Commit() //Commiting all insert transaction changes
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error in transaction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Person is Added!"})
}
