package main

import (
	"bytes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

var e error

func sendToWebhook(d map[string]string, c chan int) {

	_d := converter(d)
	if e != nil {
		c <- 500
		return
	}
	resp, e := http.Post("https://webhook.site/7bca21aa-8fe9-4b1e-87a2-74aa653a12c3", "application/json", bytes.NewBuffer(_d))
	if e != nil {
		c <- 500
		return
	}
	c <- resp.StatusCode

}

func worker(jobs <-chan map[string]string, results chan<- int) {
	for j := range jobs {
		ch := make(chan int)
		go sendToWebhook(j, ch)
		_status := <-ch
		results <- _status
	}
}

func form(c *gin.Context) {
	d := map[string]string{}

	err := c.ShouldBindJSON(&d)

	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	jobs <- d

	_status := <-results

	c.JSON(_status, gin.H{"status": _status})
}

var jobs = make(chan map[string]string, 100)
var results = make(chan int, 100)

func main() {
	r := gin.Default()

	for i := 0; i < 100; i++ {
		go worker(jobs, results)
	}

	r.GET("/ping", ping)
	r.POST("/form", form)

	r.Run(":8000")

}
