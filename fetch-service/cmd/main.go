package main

import (
	"context"
	"net/http"

	"github.com/Manas8803/Cloudy-Messenger/fetch-service/utility"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type ReqData struct {
	Email string `json:"email"`
}

func Handler(ctx context.Context, request events.CloudWatchEvent) (events.APIGatewayProxyResponse, error) {
	var res events.APIGatewayProxyResponse
	//TODO :  Api call
	
	//* Invoke notify-service
	err := utility.InvokeLambda()
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
