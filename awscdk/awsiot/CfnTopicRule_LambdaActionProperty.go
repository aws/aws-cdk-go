package awsiot


// Describes an action to invoke a Lambda function.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaActionProperty := &LambdaActionProperty{
//   	FunctionArn: jsii.String("functionArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-lambdaaction.html
//
type CfnTopicRule_LambdaActionProperty struct {
	// The ARN of the Lambda function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-lambdaaction.html#cfn-iot-topicrule-lambdaaction-functionarn
	//
	FunctionArn *string `field:"optional" json:"functionArn" yaml:"functionArn"`
}

