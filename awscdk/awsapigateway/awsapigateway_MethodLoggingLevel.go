package awsapigateway


// Example:
//   api := apigateway.NewRestApi(this, jsii.String("books"), &restApiProps{
//   	cloudWatchRole: jsii.Boolean(true),
//   	deployOptions: &stageOptions{
//   		loggingLevel: apigateway.methodLoggingLevel_INFO,
//   		dataTraceEnabled: jsii.Boolean(true),
//   	},
//   })
//
type MethodLoggingLevel string

const (
	MethodLoggingLevel_OFF MethodLoggingLevel = "OFF"
	MethodLoggingLevel_ERROR MethodLoggingLevel = "ERROR"
	MethodLoggingLevel_INFO MethodLoggingLevel = "INFO"
)

