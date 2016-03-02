package handlers

import (
	"encoding/json"
	"io"
	"log"
	"strings"

	"github.com/Clever/kayvee-logger-service/restapi/operations"
	"github.com/go-swagger/go-swagger/httpkit/middleware"
	"gopkg.in/Clever/kayvee-go.v2/validator"
)

// LogHandler is an implementation of `operations.LogHandler` that prints out
// a given kayvee event string.
type LogHandler struct {
	logger *log.Logger
}

// NewLogHandler returns a new `LogHandler` that prints log lines to the given
// output.
func NewLogHandler(output io.Writer) operations.LogHandler {
	return LogHandler{
		logger: log.New(output, "", 0),
	}
}

// Handle logs the request payload to the configured output.
func (handler LogHandler) Handle(params operations.LogParams) middleware.Responder {
	event := strings.TrimSpace(params.Event)
	err := validator.ValidateJSONFormat(event)
	if err != nil {
		return operations.NewLogBadRequest()
	}

	// Add a `via_kayvee_logger_service` field to identify logs proxied through
	// this service.
	var kayveeData map[string]interface{}
	json.Unmarshal([]byte(event), &kayveeData)
	kayveeData["via_kayvee_logger_service"] = true
	finalOutput, _ := json.Marshal(kayveeData)

	handler.logger.Println(string(finalOutput))

	return operations.NewLogOK()
}
