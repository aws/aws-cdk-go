package awsiot


// Describes an action to invoke a Lambda function.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaActionProperty := &lambdaActionProperty{
//   	functionArn: jsii.String("functionArn"),
//   }
//
type CfnTopicRule_LambdaActionProperty struct {
	// The ARN of the Lambda function.
	FunctionArn *string `field:"optional" json:"functionArn" yaml:"functionArn"`
}

