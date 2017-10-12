package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sendgrid/rest"
	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

var sendgridkey = "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"

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
		// log.Fatal("$PORT must be set")
	}
	router := gin.New()
	router.Use(gin.Logger())
	router.GET("/sendmail/:to", sendmail)

	router.Run(":" + port)
}

func sendmail(c *gin.Context) {
	to := c.Param("to")
	res, err := sendmailWithSendGrid(to)
	log.Println(res)
	if err != nil {
		c.JSON(400, nil)
	} else {
		log.Println(res.Body)
		c.JSON(res.StatusCode, res.Body)
	}

}

func sendmailWithSendGrid(sendTo string) (*rest.Response, error) {
	from := mail.NewEmail("Example User", "test@example.com")
	subject := "Sending with SendGrid is Fun"
	to := mail.NewEmail("Example User", sendTo)
	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(sendgridkey)
	return client.Send(message)
}
