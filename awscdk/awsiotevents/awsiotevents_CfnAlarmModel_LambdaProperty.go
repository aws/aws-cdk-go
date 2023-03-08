package awsiotevents


// Calls a Lambda function, passing in information about the detector model instance and the event that triggered the action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaProperty := &lambdaProperty{
//   	functionArn: jsii.String("functionArn"),
//
//   	// the properties below are optional
//   	payload: &payloadProperty{
//   		contentExpression: jsii.String("contentExpression"),
//   		type: jsii.String("type"),
//   	},
//   }
//
type CfnAlarmModel_LambdaProperty struct {
	// The ARN of the Lambda function that is executed.
	FunctionArn *string `field:"required" json:"functionArn" yaml:"functionArn"`
	// You can configure the action payload when you send a message to a Lambda function.
	Payload interface{} `field:"optional" json:"payload" yaml:"payload"`
}

