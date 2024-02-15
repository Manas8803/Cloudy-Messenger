package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Manas8803/Cloudy-Messenger/fetch-service/utility"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type ReqData struct {
	Email string `json:"email"`
}

func Handler(ctx context.Context, request events.CloudWatchEvent) (events.APIGatewayProxyResponse, error) {
	var res events.APIGatewayProxyResponse

	//* Get current date
	currentDate := getCurrentDate();

	//* Api call
	resp, err := http.Get("https://api.openweathermap.org/data/3.0/onecall/day_summary?lat=21.1458&lon=79.088860&date="+currentDate+"&appid="+os.Getenv("WEATHER_API_KEY"))
	if err != nil {
		log.Println("Error in fetching weather data : ", err.Error())
	}
	log.Println("Response : ", resp)

	//* Invoke notify-service
	err = utility.InvokeLambda(resp)
	if err != nil {
		utility.RespondWithError(&res, http.StatusInternalServerError, "Internal Server Error : Invoking Lambda")
		return res, err
	}

	utility.RespondWithJSON(&res, http.StatusOK, map[string]interface{}{"message": "Daily forecast sent!"})
	return res, nil
}

func getCurrentDate() string {
	currentTime := time.Now()

	currentDate := currentTime.Format("2006-01-02")
	return currentDate
}
func main() {
	lambda.Start(Handler)
}
