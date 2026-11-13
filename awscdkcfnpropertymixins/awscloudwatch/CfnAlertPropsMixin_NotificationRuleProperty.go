package awscloudwatch


// A notification rule: which alert states trigger it, and where the notification is sent.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   notificationRuleProperty := &NotificationRuleProperty{
//   	Target: &TargetProperty{
//   		Arn: jsii.String("arn"),
//   		Metadata: map[string]*string{
//   			"metadataKey": jsii.String("metadata"),
//   		},
//   		Type: jsii.String("type"),
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
type CfnAlertPropsMixin_NotificationRuleProperty struct {
	// Where the notification is sent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alert-notificationrule.html#cfn-cloudwatch-alert-notificationrule-target
	//
	Target interface{} `field:"optional" json:"target" yaml:"target"`
	// The alert states that trigger the notification rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alert-notificationrule.html#cfn-cloudwatch-alert-notificationrule-trigger
	//
	Trigger interface{} `field:"optional" json:"trigger" yaml:"trigger"`
}

