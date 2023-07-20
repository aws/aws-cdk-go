package awsiotevents


// Information required to publish the Amazon SNS message.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   snsProperty := &SnsProperty{
//   	TargetArn: jsii.String("targetArn"),
//
//   	// the properties below are optional
//   	Payload: &PayloadProperty{
//   		ContentExpression: jsii.String("contentExpression"),
//   		Type: jsii.String("type"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-sns.html
//
type CfnDetectorModel_SnsProperty struct {
	// The ARN of the Amazon SNS target where the message is sent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-sns.html#cfn-iotevents-detectormodel-sns-targetarn
	//
	TargetArn *string `field:"required" json:"targetArn" yaml:"targetArn"`
	// You can configure the action payload when you send a message as an Amazon SNS push notification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-sns.html#cfn-iotevents-detectormodel-sns-payload
	//
	Payload interface{} `field:"optional" json:"payload" yaml:"payload"`
}

