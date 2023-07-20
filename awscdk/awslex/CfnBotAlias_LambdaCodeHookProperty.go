package awslex


// Specifies a Lambda function that verifies requests to a bot or fulfills the user's request to a bot.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaCodeHookProperty := &LambdaCodeHookProperty{
//   	CodeHookInterfaceVersion: jsii.String("codeHookInterfaceVersion"),
//   	LambdaArn: jsii.String("lambdaArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-botalias-lambdacodehook.html
//
type CfnBotAlias_LambdaCodeHookProperty struct {
	// The version of the request-response that you want Amazon Lex to use to invoke your Lambda function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-botalias-lambdacodehook.html#cfn-lex-botalias-lambdacodehook-codehookinterfaceversion
	//
	CodeHookInterfaceVersion *string `field:"required" json:"codeHookInterfaceVersion" yaml:"codeHookInterfaceVersion"`
	// The Amazon Resource Name (ARN) of the Lambda function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-botalias-lambdacodehook.html#cfn-lex-botalias-lambdacodehook-lambdaarn
	//
	LambdaArn *string `field:"required" json:"lambdaArn" yaml:"lambdaArn"`
}

