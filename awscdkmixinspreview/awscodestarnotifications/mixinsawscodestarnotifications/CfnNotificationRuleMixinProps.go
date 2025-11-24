package mixinsawscodestarnotifications


// Properties for CfnNotificationRulePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnNotificationRuleMixinProps := &CfnNotificationRuleMixinProps{
//   	CreatedBy: jsii.String("createdBy"),
//   	DetailType: jsii.String("detailType"),
//   	EventTypeId: jsii.String("eventTypeId"),
//   	EventTypeIds: []*string{
//   		jsii.String("eventTypeIds"),
//   	},
//   	Name: jsii.String("name"),
//   	Resource: jsii.String("resource"),
//   	Status: jsii.String("status"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	TargetAddress: jsii.String("targetAddress"),
//   	Targets: []interface{}{
//   		&TargetProperty{
//   			TargetAddress: jsii.String("targetAddress"),
//   			TargetType: jsii.String("targetType"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarnotifications-notificationrule.html
//
type CfnNotificationRuleMixinProps struct {
	// The name or email alias of the person who created the notification rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarnotifications-notificationrule.html#cfn-codestarnotifications-notificationrule-createdby
	//
	CreatedBy *string `field:"optional" json:"createdBy" yaml:"createdBy"`
	// The level of detail to include in the notifications for this resource.
	//
	// `BASIC` will include only the contents of the event as it would appear in Amazon CloudWatch. `FULL` will include any supplemental information provided by AWS CodeStar Notifications and/or the service for the resource for which the notification is created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarnotifications-notificationrule.html#cfn-codestarnotifications-notificationrule-detailtype
	//
	DetailType *string `field:"optional" json:"detailType" yaml:"detailType"`
	// The event type associated with this notification rule.
	//
	// For a complete list of event types and IDs, see [Notification concepts](https://docs.aws.amazon.com/dtconsole/latest/userguide/concepts.html#concepts-api) in the *Developer Tools Console User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarnotifications-notificationrule.html#cfn-codestarnotifications-notificationrule-eventtypeid
	//
	EventTypeId *string `field:"optional" json:"eventTypeId" yaml:"eventTypeId"`
	// A list of event types associated with this notification rule.
	//
	// For a complete list of event types and IDs, see [Notification concepts](https://docs.aws.amazon.com/dtconsole/latest/userguide/concepts.html#concepts-api) in the *Developer Tools Console User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarnotifications-notificationrule.html#cfn-codestarnotifications-notificationrule-eventtypeids
	//
	EventTypeIds *[]*string `field:"optional" json:"eventTypeIds" yaml:"eventTypeIds"`
	// The name for the notification rule.
	//
	// Notification rule names must be unique in your AWS account .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarnotifications-notificationrule.html#cfn-codestarnotifications-notificationrule-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) of the resource to associate with the notification rule.
	//
	// Supported resources include pipelines in AWS CodePipeline , repositories in AWS CodeCommit , and build projects in AWS CodeBuild .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarnotifications-notificationrule.html#cfn-codestarnotifications-notificationrule-resource
	//
	Resource *string `field:"optional" json:"resource" yaml:"resource"`
	// The status of the notification rule.
	//
	// The default value is `ENABLED` . If the status is set to `DISABLED` , notifications aren't sent for the notification rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarnotifications-notificationrule.html#cfn-codestarnotifications-notificationrule-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// A list of tags to apply to this notification rule.
	//
	// Key names cannot start with " `aws` ".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarnotifications-notificationrule.html#cfn-codestarnotifications-notificationrule-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The Amazon Resource Name (ARN) of the Amazon  topic or  client.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarnotifications-notificationrule.html#cfn-codestarnotifications-notificationrule-targetaddress
	//
	TargetAddress *string `field:"optional" json:"targetAddress" yaml:"targetAddress"`
	// A list of Amazon Resource Names (ARNs) of Amazon  topics and  clients to associate with the notification rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarnotifications-notificationrule.html#cfn-codestarnotifications-notificationrule-targets
	//
	Targets interface{} `field:"optional" json:"targets" yaml:"targets"`
}

