package route

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title       string
	Author      string
	Link        string
	Description string
	Published   string
	Doi         string
}

type Articles []*Article

func GetArticlesAll(c *gin.Context) {
	// Connect DB
	db, err := connectDB()
	fmt.Println("TETS")
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	// Close DB
	defer closeDB(db)

	// Get all articles data
	var articles Articles
	db.Find(&articles)

	// Send JSON as a response
	c.JSON(http.StatusOK, articles)
}

func PostArticle(c *gin.Context) {
	// Connect DB
	db, err := connectDB()
	fmt.Println("TETS")
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	// Close DB
	defer closeDB(db)

	var json Article
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&json); err.Error != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
	}

	c.String(http.StatusOK, "OK")
}

func GetAction(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	message := name + " is " + action
	c.String(http.StatusOK, message)
}

func GetWelcome(c *gin.Context) {
	// /welcome?firstname=Jane&lastname=Doe
	firstname := c.DefaultQuery("firstname", "Guest") // If not exists, return Guest in this case
	lastname := c.Query("lastname")                   // if not exists, return "" empty string
	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
}
