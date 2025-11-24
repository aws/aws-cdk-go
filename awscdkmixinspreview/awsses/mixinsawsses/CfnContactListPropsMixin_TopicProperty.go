package mixinsawsses


// An interest group, theme, or label within a list.
//
// Lists can have multiple topics.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   topicProperty := &TopicProperty{
//   	DefaultSubscriptionStatus: jsii.String("defaultSubscriptionStatus"),
//   	Description: jsii.String("description"),
//   	DisplayName: jsii.String("displayName"),
//   	TopicName: jsii.String("topicName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-contactlist-topic.html
//
type CfnContactListPropsMixin_TopicProperty struct {
	// The default subscription status to be applied to a contact if the contact has not noted their preference for subscribing to a topic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-contactlist-topic.html#cfn-ses-contactlist-topic-defaultsubscriptionstatus
	//
	DefaultSubscriptionStatus *string `field:"optional" json:"defaultSubscriptionStatus" yaml:"defaultSubscriptionStatus"`
	// A description of what the topic is about, which the contact will see.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-contactlist-topic.html#cfn-ses-contactlist-topic-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the topic the contact will see.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-contactlist-topic.html#cfn-ses-contactlist-topic-displayname
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The name of the topic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-contactlist-topic.html#cfn-ses-contactlist-topic-topicname
	//
	TopicName *string `field:"optional" json:"topicName" yaml:"topicName"`
}

