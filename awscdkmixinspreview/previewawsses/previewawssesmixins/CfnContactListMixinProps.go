package previewawssesmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnContactListPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnContactListMixinProps := &CfnContactListMixinProps{
//   	ContactListName: jsii.String("contactListName"),
//   	Description: jsii.String("description"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Topics: []interface{}{
//   		&TopicProperty{
//   			DefaultSubscriptionStatus: jsii.String("defaultSubscriptionStatus"),
//   			Description: jsii.String("description"),
//   			DisplayName: jsii.String("displayName"),
//   			TopicName: jsii.String("topicName"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-contactlist.html
//
type CfnContactListMixinProps struct {
	// The name of the contact list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-contactlist.html#cfn-ses-contactlist-contactlistname
	//
	ContactListName *string `field:"optional" json:"contactListName" yaml:"contactListName"`
	// A description of what the contact list is about.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-contactlist.html#cfn-ses-contactlist-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The tags associated with a contact list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-contactlist.html#cfn-ses-contactlist-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// An interest group, theme, or label within a list.
	//
	// A contact list can have multiple topics.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-contactlist.html#cfn-ses-contactlist-topics
	//
	Topics interface{} `field:"optional" json:"topics" yaml:"topics"`
}

