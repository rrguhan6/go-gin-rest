package main

import (
	"bytes"
	"encoding/json"
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

func convertToNestedData(d DATA) NestedData {
	return NestedData{
		Event:           d.Ev,
		EventType:       d.Et,
		AppID:           d.ID,
		UserID:          d.UID,
		MessageID:       d.Mid,
		PageTitle:       d.T,
		PageURL:         d.P,
		BrowserLanguage: d.L,
		ScreenSize:      d.Sc,
		Attributes: Attributes{
			FormVarient: ValueType{d.Atrv1, d.Atrt1},
			Ref:         ValueType{d.Atrv2, d.Atrt2},
		},
		Traits: Traits{
			Name:  ValueType{d.Uatrv1, d.Uatrt1},
			Email: ValueType{d.Uatrv2, d.Uatrt2},
			Age:   ValueType{d.Uatrv3, d.Uatrt3},
		},
	}
}

var e error

func sendToWebhook(d DATA, c chan int) {

	_d := convertToNestedData(d)
	jsonValue, e := json.Marshal(conventionalMarshaller{_d})
	if e != nil {
		c <- 500
		return
	}
	resp, e := http.Post("https://webhook.site/7dce8115-29d8-4f04-bba7-1e5f8850c66a", "application/json", bytes.NewBuffer(jsonValue))
	if e != nil {
		c <- 500
		return
	}
	c <- resp.StatusCode

}

func form(c *gin.Context) {
	var d DATA
	err := c.ShouldBindJSON(&d)

	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ch := make(chan int)
	go sendToWebhook(d, ch)

	_status := <-ch

	c.JSON(_status, gin.H{"status": _status})
}

func main() {
	r := gin.Default()

	r.GET("/ping", ping)
	r.POST("/form", form)

	r.Run(":8000")
}
