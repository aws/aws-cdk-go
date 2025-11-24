package mixinsawsnotifications


// Properties for CfnEventRulePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnEventRuleMixinProps := &CfnEventRuleMixinProps{
//   	EventPattern: jsii.String("eventPattern"),
//   	EventType: jsii.String("eventType"),
//   	NotificationConfigurationArn: jsii.String("notificationConfigurationArn"),
//   	Regions: []*string{
//   		jsii.String("regions"),
//   	},
//   	Source: jsii.String("source"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-notifications-eventrule.html
//
type CfnEventRuleMixinProps struct {
	// An additional event pattern used to further filter the events this `EventRule` receives.
	//
	// For more information, see [Amazon EventBridge event patterns](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-event-patterns.html) in the *Amazon EventBridge User Guide.*
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-notifications-eventrule.html#cfn-notifications-eventrule-eventpattern
	//
	EventPattern *string `field:"optional" json:"eventPattern" yaml:"eventPattern"`
	// The event type this rule should match with the EventBridge events.
	//
	// It must match with atleast one of the valid EventBridge event types. For example, Amazon EC2 Instance State change Notification and Amazon CloudWatch State Change. For more information, see [Event delivery from AWS services](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-service-event.html#eb-service-event-delivery-level) in the *Amazon EventBridge User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-notifications-eventrule.html#cfn-notifications-eventrule-eventtype
	//
	EventType *string `field:"optional" json:"eventType" yaml:"eventType"`
	// The ARN for the `NotificationConfiguration` associated with this `EventRule` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-notifications-eventrule.html#cfn-notifications-eventrule-notificationconfigurationarn
	//
	NotificationConfigurationArn *string `field:"optional" json:"notificationConfigurationArn" yaml:"notificationConfigurationArn"`
	// A list of AWS Regions that send events to this `EventRule` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-notifications-eventrule.html#cfn-notifications-eventrule-regions
	//
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
	// The event source this rule should match with the EventBridge event sources.
	//
	// It must match with atleast one of the valid EventBridge event sources. Only AWS service sourced events are supported. For example, `aws.ec2` and `aws.cloudwatch` . For more information, see [Event delivery from AWS services](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-service-event.html#eb-service-event-delivery-level) in the *Amazon EventBridge User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-notifications-eventrule.html#cfn-notifications-eventrule-source
	//
	Source *string `field:"optional" json:"source" yaml:"source"`
}

