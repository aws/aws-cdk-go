package awsiotevents


// Information required to publish the Amazon SNS message.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   snsProperty := &snsProperty{
//   	targetArn: jsii.String("targetArn"),
//
//   	// the properties below are optional
//   	payload: &payloadProperty{
//   		contentExpression: jsii.String("contentExpression"),
//   		type: jsii.String("type"),
//   	},
//   }
//
type CfnDetectorModel_SnsProperty struct {
	// The ARN of the Amazon SNS target where the message is sent.
	TargetArn *string `field:"required" json:"targetArn" yaml:"targetArn"`
	// You can configure the action payload when you send a message as an Amazon SNS push notification.
	Payload interface{} `field:"optional" json:"payload" yaml:"payload"`
}

