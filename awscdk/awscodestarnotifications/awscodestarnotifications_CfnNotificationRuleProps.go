package awscodestarnotifications


// Properties for defining a `CfnNotificationRule`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnNotificationRuleProps := &cfnNotificationRuleProps{
//   	detailType: jsii.String("detailType"),
//   	eventTypeIds: []*string{
//   		jsii.String("eventTypeIds"),
//   	},
//   	name: jsii.String("name"),
//   	resource: jsii.String("resource"),
//   	targets: []interface{}{
//   		&targetProperty{
//   			targetAddress: jsii.String("targetAddress"),
//   			targetType: jsii.String("targetType"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	createdBy: jsii.String("createdBy"),
//   	eventTypeId: jsii.String("eventTypeId"),
//   	status: jsii.String("status"),
//   	tags: tags,
//   	targetAddress: jsii.String("targetAddress"),
//   }
//
type CfnNotificationRuleProps struct {
	// The level of detail to include in the notifications for this resource.
	//
	// `BASIC` will include only the contents of the event as it would appear in Amazon CloudWatch. `FULL` will include any supplemental information provided by AWS CodeStar Notifications and/or the service for the resource for which the notification is created.
	DetailType *string `field:"required" json:"detailType" yaml:"detailType"`
	// A list of event types associated with this notification rule.
	//
	// For a complete list of event types and IDs, see [Notification concepts](https://docs.aws.amazon.com/codestar-notifications/latest/userguide/concepts.html#concepts-api) in the *Developer Tools Console User Guide* .
	EventTypeIds *[]*string `field:"required" json:"eventTypeIds" yaml:"eventTypeIds"`
	// The name for the notification rule.
	//
	// Notification rule names must be unique in your AWS account .
	Name *string `field:"required" json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) of the resource to associate with the notification rule.
	//
	// Supported resources include pipelines in AWS CodePipeline , repositories in AWS CodeCommit , and build projects in AWS CodeBuild .
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	// A list of Amazon Resource Names (ARNs) of AWS Chatbot topics and AWS Chatbot clients to associate with the notification rule.
	Targets interface{} `field:"required" json:"targets" yaml:"targets"`
	// `AWS::CodeStarNotifications::NotificationRule.CreatedBy`.
	CreatedBy *string `field:"optional" json:"createdBy" yaml:"createdBy"`
	// `AWS::CodeStarNotifications::NotificationRule.EventTypeId`.
	EventTypeId *string `field:"optional" json:"eventTypeId" yaml:"eventTypeId"`
	// The status of the notification rule.
	//
	// The default value is `ENABLED` . If the status is set to `DISABLED` , notifications aren't sent for the notification rule.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// A list of tags to apply to this notification rule.
	//
	// Key names cannot start with " `aws` ".
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// `AWS::CodeStarNotifications::NotificationRule.TargetAddress`.
	TargetAddress *string `field:"optional" json:"targetAddress" yaml:"targetAddress"`
}

