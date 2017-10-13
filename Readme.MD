# Tutorial Deployments
## Heroku

### Require
* [Heroku CLI](https://devcenter.heroku.com/articles/heroku-cli)
* [godep](https://github.com/tools/godep)

### Initial project with example code

```go
package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sendgrid/rest"
	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

var sendgridkey = "SENDGRID_API_KEY"

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}
	router := gin.New()
	router.Use(gin.Logger())
	router.GET("/sendmail/:to", sendmail)

	router.Run(":" + port)
}

func sendmail(c *gin.Context) {
	to := c.Param("to")
	res, err := sendmailWithSendGrid(to)
	if err != nil {
		c.JSON(400, err.Error())
	} else {
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

```
### Save dependencies to your project with godep

```bash
go get github.com/gin-gonic/gin
go get github.com/sendgrid/sendgrid-go
go get github.com/sendgrid/rest
godep save ./...
```

### Define a Procfile
follow document [link.](https://devcenter.heroku.com/articles/getting-started-with-nodejs#define-a-procfile)

The `Procfile` in the example app you deployed looks like this:
```
web: sendgrid-go-heroku
```

create app in heroku with heroku-cli run `heroku apps:create sendgrid-go-heroku`
and deploy with simple git command like this:
```
git add .
git commit -m 'deploy to heroku'
git push heroku master
```

