package utility

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

func RespondWithJSON(res *events.APIGatewayProxyResponse, status int, data interface{}) {
	body, err := json.Marshal(data)
	if err != nil {
		log.Println("Error in marshalling body data : ", err)
	}
	res.Body = string(body)
}

func RespondWithError(res *events.APIGatewayProxyResponse, status int, message string) {
	response := map[string]string{"error": message}
	RespondWithJSON(res, status, response)
}
