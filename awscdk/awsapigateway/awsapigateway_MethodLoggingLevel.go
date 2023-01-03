package awsapigateway


// Example:
//   api := apigateway.NewRestApi(this, jsii.String("books"), &restApiProps{
//   	deployOptions: &stageOptions{
//   		loggingLevel: apigateway.methodLoggingLevel_INFO,
//   		dataTraceEnabled: jsii.Boolean(true),
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

