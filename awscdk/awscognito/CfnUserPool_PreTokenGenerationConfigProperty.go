package awscognito


// The properties of a pre token generation Lambda trigger.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   preTokenGenerationConfigProperty := &PreTokenGenerationConfigProperty{
//   	LambdaArn: jsii.String("lambdaArn"),
//   	LambdaVersion: jsii.String("lambdaVersion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-pretokengenerationconfig.html
//
type CfnUserPool_PreTokenGenerationConfigProperty struct {
	// The Amazon Resource Name (ARN) of the function that you want to assign to your Lambda trigger.
	//
	// This parameter and the `PreTokenGeneration` property of `LambdaConfig` have the same value. For new instances of pre token generation triggers, set `LambdaArn` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-pretokengenerationconfig.html#cfn-cognito-userpool-pretokengenerationconfig-lambdaarn
	//
	LambdaArn *string `field:"optional" json:"lambdaArn" yaml:"lambdaArn"`
	// The user pool trigger version of the request that Amazon Cognito sends to your Lambda function.
	//
	// Higher-numbered versions add fields that support new features.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-pretokengenerationconfig.html#cfn-cognito-userpool-pretokengenerationconfig-lambdaversion
	//
	LambdaVersion *string `field:"optional" json:"lambdaVersion" yaml:"lambdaVersion"`
}

