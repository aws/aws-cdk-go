package awsmedialive


// Properties for defining a `CfnEventBridgeRuleTemplate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEventBridgeRuleTemplateProps := &CfnEventBridgeRuleTemplateProps{
//   	EventType: jsii.String("eventType"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	EventTargets: []interface{}{
//   		&EventBridgeRuleTemplateTargetProperty{
//   			Arn: jsii.String("arn"),
//   		},
//   	},
//   	GroupIdentifier: jsii.String("groupIdentifier"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-eventbridgeruletemplate.html
//
type CfnEventBridgeRuleTemplateProps struct {
	// The type of event to match with the rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-eventbridgeruletemplate.html#cfn-medialive-eventbridgeruletemplate-eventtype
	//
	EventType *string `field:"required" json:"eventType" yaml:"eventType"`
	// A resource's name.
	//
	// Names must be unique within the scope of a resource type in a specific region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-eventbridgeruletemplate.html#cfn-medialive-eventbridgeruletemplate-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// A resource's optional description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-eventbridgeruletemplate.html#cfn-medialive-eventbridgeruletemplate-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The destinations that will receive the event notifications.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-eventbridgeruletemplate.html#cfn-medialive-eventbridgeruletemplate-eventtargets
	//
	EventTargets interface{} `field:"optional" json:"eventTargets" yaml:"eventTargets"`
	// An eventbridge rule template group's identifier.
	//
	// Can be either be its id or current name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-eventbridgeruletemplate.html#cfn-medialive-eventbridgeruletemplate-groupidentifier
	//
	GroupIdentifier *string `field:"optional" json:"groupIdentifier" yaml:"groupIdentifier"`
	// Represents the tags associated with a resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-eventbridgeruletemplate.html#cfn-medialive-eventbridgeruletemplate-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

