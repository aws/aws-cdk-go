package awsses

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnContactList`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnContactListProps := &cfnContactListProps{
//   	contactListName: jsii.String("contactListName"),
//   	description: jsii.String("description"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	topics: []interface{}{
//   		&topicProperty{
//   			defaultSubscriptionStatus: jsii.String("defaultSubscriptionStatus"),
//   			displayName: jsii.String("displayName"),
//   			topicName: jsii.String("topicName"),
//
//   			// the properties below are optional
//   			description: jsii.String("description"),
//   		},
//   	},
//   }
//
type CfnContactListProps struct {
	// The name of the contact list.
	ContactListName *string `field:"optional" json:"contactListName" yaml:"contactListName"`
	// A description of what the contact list is about.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The tags associated with a contact list.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// An interest group, theme, or label within a list.
	//
	// A contact list can have multiple topics.
	Topics interface{} `field:"optional" json:"topics" yaml:"topics"`
}

