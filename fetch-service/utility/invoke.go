package utility

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

type Payload_Body struct {
	Body string `json:"body"`
}

func InvokeLambda(res *http.Response) error {

	res_body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return err
	}

	defer res.Body.Close()

	var weatherData map[string]interface{}
	if err := json.Unmarshal(res_body, &weatherData); err != nil {
		log.Println("Error parsing JSON:", err)
		return err
	}

	sess, err := session.NewSession()
	if err != nil {
		log.Println("Error in creating session : ", err.Error())
		return err
	}

	client := lambda.New(sess)

	//! DO NOT REMOVE
	data, err := json.Marshal(weatherData)
	if err != nil {
		log.Println("Error in marshalling data : ", err.Error())
		return err
	}

	log.Println("Data : ", string(data))
	body := Payload_Body{Body: string(data)}
	payload, err := json.Marshal(body)
	if err != nil {
		log.Println("Error in marshalling payload : ", err.Error())
		return err
	}
	//!
	log.Println("Payload : ", string(payload))

	//* Invoking notify function
	input := &lambda.InvokeInput{
		FunctionName: aws.String(os.Getenv("NOTIFY_FUNC_ARN")),
		Payload:        payload,
		InvocationType: aws.String("Event"),
	}

	result, err := client.Invoke(input)
	if err != nil {
		log.Println("Error invoking Lambda function:", err)
		return err
	}

	log.Println("Lambda function invoked successfully:", result)

	return nil
}
