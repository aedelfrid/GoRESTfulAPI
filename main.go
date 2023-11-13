package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

type user struct {
	ID string `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
	OwnedGames struct {*game} `json:"ownedGames"`
	Friends struct {*user} `json:"friends"`
	//authTokens?
}

type game struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Prices struct {*gamePlatformPrices} `json:"prices"`
}

type platform struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Games struct {*gamePlatformPrices} `json:"games"`
}

type gamePlatformPrices struct {
	ID string `json:"id"`
	GameID string `json:"gameId"`
	Platform struct {*platform} `json:"platform"`
	CrossplayEnabled bool `json:"crossplayEnabled"`
	Price float64 `json:"price"`
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

var users = []user{
	{ID: "1", Username: "asone0", Email: "btocque0@twitpic.com", Password: "jS8+e1YDh0"},
	{ID: "2", Username: "cashington1", Email: "tsuggitt1@privacy.gov.au", Password: "hS1$v<YH{?`?ED'"},
	{ID: "3", Username: "ogoracci2", Email: "gpoint2@mediafire.com", Password: "lK6_q%&0I<N3o4G7"},
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)

	router.Run("localhost:8000")
}

