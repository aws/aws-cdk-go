package awslambda


// This field takes in 2 values either Text or JSON.
//
// By setting this value to Text,
// will result in the current structure of logs format, whereas, by setting this value to JSON,
// Lambda will print the logs as Structured JSON Logs, with the corresponding timestamp and log level
// of each event. Selecting ‘JSON’ format will only allow customer’s to have different log level
// Application log level and the System log level.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var logGroup iLogGroup
//
//
//   lambda.NewFunction(this, jsii.String("Lambda"), &FunctionProps{
//   	Code: lambda.NewInlineCode(jsii.String("foo")),
//   	Handler: jsii.String("index.handler"),
//   	Runtime: lambda.Runtime_NODEJS_18_X(),
//   	LoggingFormat: lambda.LoggingFormat_JSON,
//   	SystemLogLevel: lambda.SystemLogLevel_INFO,
//   	ApplicationLogLevel: lambda.ApplicationLogLevel_INFO,
//   	LogGroup: logGroup,
//   })
//
type LoggingFormat string

const (
	// Lambda Logs text format.
	LoggingFormat_TEXT LoggingFormat = "TEXT"
	// Lambda structured logging in Json format.
	LoggingFormat_JSON LoggingFormat = "JSON"
)

