package awsapigateway


// Example:
//   api := apigateway.NewRestApi(this, jsii.String("books"), &RestApiProps{
//   	DeployOptions: &StageOptions{
//   		LoggingLevel: apigateway.MethodLoggingLevel_INFO,
//   		DataTraceEnabled: jsii.Boolean(true),
//   	},
//   })
//
// Experimental.
type MethodLoggingLevel string

const (
	// Experimental.
	MethodLoggingLevel_OFF MethodLoggingLevel = "OFF"
	// Experimental.
	MethodLoggingLevel_ERROR MethodLoggingLevel = "ERROR"
	// Experimental.
	MethodLoggingLevel_INFO MethodLoggingLevel = "INFO"
)

