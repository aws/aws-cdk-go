package awsses


// LambdaAction configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaActionConfig := &lambdaActionConfig{
//   	functionArn: jsii.String("functionArn"),
//
//   	// the properties below are optional
//   	invocationType: jsii.String("invocationType"),
//   	topicArn: jsii.String("topicArn"),
//   }
//
type LambdaActionConfig struct {
	// The Amazon Resource Name (ARN) of the AWS Lambda function.
	FunctionArn *string `field:"required" json:"functionArn" yaml:"functionArn"`
	// The invocation type of the AWS Lambda function.
	InvocationType *string `field:"optional" json:"invocationType" yaml:"invocationType"`
	// The Amazon Resource Name (ARN) of the Amazon SNS topic to notify when the Lambda action is executed.
	TopicArn *string `field:"optional" json:"topicArn" yaml:"topicArn"`
}

