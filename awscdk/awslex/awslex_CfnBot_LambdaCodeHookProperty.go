package awslex


// Specifies a Lambda function that verifies requests to a bot or fulfills the user's request to a bot.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaCodeHookProperty := &lambdaCodeHookProperty{
//   	codeHookInterfaceVersion: jsii.String("codeHookInterfaceVersion"),
//   	lambdaArn: jsii.String("lambdaArn"),
//   }
//
type CfnBot_LambdaCodeHookProperty struct {
	// Specifies the version of the request-response that you want Amazon Lex to use to invoke your Lambda function.
	CodeHookInterfaceVersion *string `field:"required" json:"codeHookInterfaceVersion" yaml:"codeHookInterfaceVersion"`
	// Specifies the Amazon Resource Name (ARN) of the Lambda function.
	LambdaArn *string `field:"required" json:"lambdaArn" yaml:"lambdaArn"`
}

