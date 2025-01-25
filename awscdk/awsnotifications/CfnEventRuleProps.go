package awsnotifications


// Properties for defining a `CfnEventRule`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEventRuleProps := &CfnEventRuleProps{
//   	EventType: jsii.String("eventType"),
//   	NotificationConfigurationArn: jsii.String("notificationConfigurationArn"),
//   	Regions: []*string{
//   		jsii.String("regions"),
//   	},
//   	Source: jsii.String("source"),
//
//   	// the properties below are optional
//   	EventPattern: jsii.String("eventPattern"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-notifications-eventrule.html
//
type CfnEventRuleProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-notifications-eventrule.html#cfn-notifications-eventrule-eventtype
	//
	EventType *string `field:"required" json:"eventType" yaml:"eventType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-notifications-eventrule.html#cfn-notifications-eventrule-notificationconfigurationarn
	//
	NotificationConfigurationArn *string `field:"required" json:"notificationConfigurationArn" yaml:"notificationConfigurationArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-notifications-eventrule.html#cfn-notifications-eventrule-regions
	//
	Regions *[]*string `field:"required" json:"regions" yaml:"regions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-notifications-eventrule.html#cfn-notifications-eventrule-source
	//
	Source *string `field:"required" json:"source" yaml:"source"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-notifications-eventrule.html#cfn-notifications-eventrule-eventpattern
	//
	EventPattern *string `field:"optional" json:"eventPattern" yaml:"eventPattern"`
}

