package awsssmincidents


// The Amazon  topic that's used by  to notify the incidents chat channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   notificationTargetItemProperty := &NotificationTargetItemProperty{
//   	SnsTopicArn: jsii.String("snsTopicArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmincidents-responseplan-notificationtargetitem.html
//
type CfnResponsePlan_NotificationTargetItemProperty struct {
	// The Amazon Resource Name (ARN) of the Amazon  topic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmincidents-responseplan-notificationtargetitem.html#cfn-ssmincidents-responseplan-notificationtargetitem-snstopicarn
	//
	SnsTopicArn *string `field:"optional" json:"snsTopicArn" yaml:"snsTopicArn"`
}

