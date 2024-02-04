package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Manas8803/Cloudy-Messenger/fetch-service/utility"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type ReqData struct {
	Email string `json:"email"`
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var data ReqData
	var res events.APIGatewayProxyResponse

	//* Unmarshal request body
	err := json.Unmarshal([]byte(request.Body), &data)
	if err != nil {
		utility.RespondWithError(&res, http.StatusBadRequest, "Bad Request : Unmarshalling request body")
		return res, err
	}

	//* Check for valid email
	email := data.Email
	if !utility.IsValidEmail(email) {
		utility.RespondWithError(&res, http.StatusBadRequest, "Bad Request : Invalid email")
		return res, err
	}

	//* Invoke notify-service
	err = utility.InvokeLambda(email)
	if err != nil {
		utility.RespondWithError(&res, http.StatusInternalServerError, "Internal Server Error : Invoking Lambda")
		return res, err
	}

	utility.RespondWithJSON(&res, http.StatusOK, map[string]interface{}{"message": "Daily forecast sent!"})
	return res, nil
}
func main() {
	lambda.Start(Handler)
}
