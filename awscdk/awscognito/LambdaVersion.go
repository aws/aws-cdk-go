package awscognito


// The user pool trigger version of the request that Amazon Cognito sends to your Lambda function.
//
// Example:
//   var userpool userPool
//   var preTokenGenerationFn function
//
//
//   userpool.AddTrigger(cognito.UserPoolOperation_PRE_TOKEN_GENERATION_CONFIG(), preTokenGenerationFn, cognito.LambdaVersion_V2_0)
//
type LambdaVersion string

const (
	// V1_0 trigger.
	LambdaVersion_V1_0 LambdaVersion = "V1_0"
	// V2_0 trigger.
	//
	// This is supported only for PRE_TOKEN_GENERATION trigger.
	LambdaVersion_V2_0 LambdaVersion = "V2_0"
)

