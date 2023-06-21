package awsapigateway


// Example:
//   api := apigateway.NewRestApi(this, jsii.String("books"), &RestApiProps{
//   	CloudWatchRole: jsii.Boolean(true),
//   	DeployOptions: &StageOptions{
//   		LoggingLevel: apigateway.MethodLoggingLevel_INFO,
//   		DataTraceEnabled: jsii.Boolean(true),
//   	},
//   })
//
type MethodLoggingLevel string

const (
	MethodLoggingLevel_OFF MethodLoggingLevel = "OFF"
	MethodLoggingLevel_ERROR MethodLoggingLevel = "ERROR"
	MethodLoggingLevel_INFO MethodLoggingLevel = "INFO"
)

