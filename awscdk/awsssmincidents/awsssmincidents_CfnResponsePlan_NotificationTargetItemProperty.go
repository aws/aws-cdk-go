package awsssmincidents


// The SNS topic that's used by AWS Chatbot to notify the incidents chat channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   notificationTargetItemProperty := &notificationTargetItemProperty{
//   	snsTopicArn: jsii.String("snsTopicArn"),
//   }
//
type CfnResponsePlan_NotificationTargetItemProperty struct {
	// The Amazon Resource Name (ARN) of the SNS topic.
	SnsTopicArn *string `field:"optional" json:"snsTopicArn" yaml:"snsTopicArn"`
}

