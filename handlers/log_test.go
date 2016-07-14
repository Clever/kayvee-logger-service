package handlers

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/Clever/kayvee-logger-service/restapi/operations"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLogHandlerWritesEventToOutput(t *testing.T) {
	output := &bytes.Buffer{}
	handler := NewLogHandler(output)

	inputData := map[string]interface{}{
		"source": "fabulaws",
		"level":  "error",
		"title":  "instance_starting",
		"msg":    "salt run failed",
	}
	formattedInput, err := json.Marshal(inputData)
	require.NoError(t, err)

	responder := handler.Handle(operations.LogParams{Event: string(formattedInput) + "\n"})

	assert.Equal(t, operations.NewLogOK(), responder)

	expectedOutputData := inputData
	expectedOutputData["via_kayvee_logger_service"] = true
	expectedOutput, err := json.Marshal(expectedOutputData)

	assert.Equal(t, string(expectedOutput)+"\n", string(output.Bytes()))
}

func TestLogHandlerBlocksMalformedEvents(t *testing.T) {
	output := &bytes.Buffer{}
	handler := NewLogHandler(output)

	responder := handler.Handle(operations.LogParams{Event: `["just", "a list", "of errors"]`})

	assert.Equal(t, operations.NewLogBadRequest(), responder)
	assert.Equal(t, "", string(output.Bytes()))
}

func TestLogHandlerBlocksEventsWithMissingFields(t *testing.T) {
	output := &bytes.Buffer{}
	handler := NewLogHandler(output)

	responder := handler.Handle(operations.LogParams{Event: `{"source":"app1", "title":"event3"}`})

	assert.Equal(t, operations.NewLogBadRequest(), responder)
	assert.Equal(t, "", string(output.Bytes()))
}
