package awslambda


// This field takes in 2 values either Text or JSON.
//
// By setting this value to Text,
// will result in the current structure of logs format, whereas, by setting this value to JSON,
// Lambda will print the logs as Structured JSON Logs, with the corresponding timestamp and log level
// of each event. Selecting ‘JSON’ format will only allow customer’s to have different log level
// Application log level and the System log level.
type LogFormat string

const (
	// Lambda Logs text format.
	LogFormat_TEXT LogFormat = "TEXT"
	// Lambda structured logging in Json format.
	LogFormat_JSON LogFormat = "JSON"
)

