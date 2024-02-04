package utility

import (
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

type Data struct {
	Email string `json:"email"`
}
type Payload_Body struct {
	Body string `json:"body"`
}

func InvokeLambda(email string) error {
	sess, err := session.NewSession()
	if err != nil {
		log.Println("Error in creating session : ", err.Error())
		return err
	}

	client := lambda.New(sess)
	data, err := json.Marshal(Data{Email: email})
	if err != nil {
		log.Println("Error in marshalling data : ", err.Error())
		return err
	}

	body := Payload_Body{Body: string(data)}
	payload, err := json.Marshal(body)
	if err != nil {
		log.Println("Error in marshalling payload : ", err.Error())
		return err
	}

	input := &lambda.InvokeInput{
		FunctionName:   aws.String(os.Getenv("NOTIFY_FUNC_ARN")),
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
