package errors

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

type (
	// Error will be the implementation of error built in interface
	Error struct {
		Cause  error
		Action string
		Status int
		Code   string
	}
	// Response the struct returned to the client in case of error
	Response struct {
		Error ResponseData `json:"error"`
	}
	// ResponseData the struct that actual contains the information of error
	ResponseData struct {
		Code   string `json:"code,omitempty"`
		Action string `json:"action,omitempty"`
	}
)

// Error implementation of Error method of built in struct
func (f Error) Error() string {
	cause := ""
	if f.Cause != nil {
		cause = f.Cause.Error()
	}
	return fmt.Sprintf("error calling mutants service : %s", cause)
}

// EchoErrorHandler the handler to send the custom struct in case of error
func EchoErrorHandler() func(err error, c echo.Context) {
	return func(err error, c echo.Context) {
		statusCode := http.StatusInternalServerError
		resp := Response{
			Error: ResponseData{},
		}
		// Catch 404s or other routing errors
		if he, ok := err.(*echo.HTTPError); ok {
			statusCode = he.Code
			msg := fmt.Sprintf("%s", he.Message)
			err = Error{
				Action: msg,
				Status: statusCode,
				Code:   "6560b7ad-251f-412e-b459-611b43783ad8",
			}
		}
		if e, ok := err.(Error); ok {
			statusCode = e.Status
			resp.Error.Code = e.Code
			resp.Error.Action = e.Action
		}
		err = c.JSON(statusCode, resp)
		if err != nil {
			log.Printf("couldn't parse response: %v", err)
		}
	}
}
