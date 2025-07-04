package awsses


// LambdaAction configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaActionConfig := &LambdaActionConfig{
//   	FunctionArn: jsii.String("functionArn"),
//
//   	// the properties below are optional
//   	InvocationType: jsii.String("invocationType"),
//   	TopicArn: jsii.String("topicArn"),
//   }
//
type LambdaActionConfig struct {
	// The Amazon Resource Name (ARN) of the AWS Lambda function.
	FunctionArn *string `field:"required" json:"functionArn" yaml:"functionArn"`
	// The invocation type of the AWS Lambda function.
	// Default: 'Event'.
	//
	InvocationType *string `field:"optional" json:"invocationType" yaml:"invocationType"`
	// The Amazon Resource Name (ARN) of the Amazon SNS topic to notify when the Lambda action is executed.
	// Default: - No notification is sent to SNS.
	//
	TopicArn *string `field:"optional" json:"topicArn" yaml:"topicArn"`
}

