package awsses


// An interest group, theme, or label within a list.
//
// Lists can have multiple topics.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   topicProperty := &topicProperty{
//   	defaultSubscriptionStatus: jsii.String("defaultSubscriptionStatus"),
//   	displayName: jsii.String("displayName"),
//   	topicName: jsii.String("topicName"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   }
//
type CfnContactList_TopicProperty struct {
	// The default subscription status to be applied to a contact if the contact has not noted their preference for subscribing to a topic.
	DefaultSubscriptionStatus *string `field:"required" json:"defaultSubscriptionStatus" yaml:"defaultSubscriptionStatus"`
	// The name of the topic the contact will see.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The name of the topic.
	TopicName *string `field:"required" json:"topicName" yaml:"topicName"`
	// A description of what the topic is about, which the contact will see.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

