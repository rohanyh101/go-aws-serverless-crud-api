package handlers

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/roh4nyh/go_serverless/pkg/errors"
	"github.com/roh4nyh/go_serverless/pkg/user"
)

type ErrorBody struct {
	ErrorMsg *string `json:"error,omitempty"`
}

func GetUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {
	email := req.QueryStringParameters["email"]

	// If email is provided, fetch single user by email
	if len(email) > 0 {
		result, err := user.FetchUser(email, tableName, dynaClient)
		if err != nil {
			return apiResponse(http.StatusInternalServerError, ErrorBody{aws.String(err.Error())})
		}

		if result.Email == "" {
			return apiResponse(http.StatusNotFound, ErrorBody{aws.String(errors.ErrorUserDoesNotExist.Error())})
		}

		return apiResponse(http.StatusOK, result)
	}

	// If email is not provided, fetch all users
	result, err := user.FetchUsers(tableName, dynaClient)
	if err != nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}

	return apiResponse(http.StatusOK, result)
}

func CreateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {
	result, err := user.CreateUser(req, tableName, dynaClient)
	if err != nil {
		return apiResponse(http.StatusInternalServerError, ErrorBody{aws.String(err.Error())})
	}

	return apiResponse(http.StatusCreated, result)
}

func UpdateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {
	result, err := user.UpdateUser(req, tableName, dynaClient)
	if err != nil {
		return apiResponse(http.StatusInternalServerError, ErrorBody{aws.String(err.Error())})
	}

	return apiResponse(http.StatusCreated, result)
}

func DeleteUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {
	err := user.DeleteUser(req, tableName, dynaClient)
	if err != nil {
		return apiResponse(http.StatusInternalServerError, ErrorBody{aws.String(err.Error())})
	}

	return apiResponse(http.StatusNoContent, nil)
}

func UnhandledMethod() (*events.APIGatewayProxyResponse, error) {
	return apiResponse(http.StatusMethodNotAllowed, ErrorBody{aws.String(errors.ErrorMethodNotAllowed.Error())})
}
