package awscloudwatch


// A notification rule: which alert states trigger it, and where the notification is sent.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   notificationRuleProperty := &NotificationRuleProperty{
//   	Target: &TargetProperty{
//   		Arn: jsii.String("arn"),
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		Metadata: map[string]*string{
//   			"metadataKey": jsii.String("metadata"),
//   		},
//   	},
//   	Trigger: &TriggerProperty{
//   		StateValues: []*string{
//   			jsii.String("stateValues"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alert-notificationrule.html
//
type CfnAlert_NotificationRuleProperty struct {
	// Where the notification is sent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alert-notificationrule.html#cfn-cloudwatch-alert-notificationrule-target
	//
	Target interface{} `field:"required" json:"target" yaml:"target"`
	// The alert states that trigger the notification rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alert-notificationrule.html#cfn-cloudwatch-alert-notificationrule-trigger
	//
	Trigger interface{} `field:"required" json:"trigger" yaml:"trigger"`
}

