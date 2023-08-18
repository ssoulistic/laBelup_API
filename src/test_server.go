package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 1. 고객의 ip, 의뢰한 작업(ip 의뢰) 데이터.
type dataset struct {
	request_data string `json:"target_ip"`
	client       string `json:"client_ip"`
}

var sample = []dataset{
	{request_data: "1", client: "1-1"},
	{request_data: "2", client: "2-1"},
	{request_data: "3", client: "3-1"},
}

// 2. GET 요청 -> /데이터베이스 검색 + ip
func getip(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, sample)
}

func main() {
	router := gin.Default()
	router.GET("/result", getip)
	router.GET("/result/:fromdb", getipfromdb)
	router.POST("/result", postip)
	router.Run("localhost:8080")
}

// 3. 결과를 토대로 자동으로 정보를 반환
func postip(c *gin.Context) {
	var newip dataset
	if err := c.BindJSON(&newip); err != nil {
		return
	}
	sample = append(sample, newip)
	c.IndentedJSON(http.StatusCreated, newip)
}

func getipfromdb(c *gin.Context) {
	id := c.Param("request_data")

	for _, a := range sample {
		if a.request_data == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ip not found"})
}
