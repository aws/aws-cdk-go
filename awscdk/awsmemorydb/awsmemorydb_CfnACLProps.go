package awsmemorydb

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnACL`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnACLProps := &cfnACLProps{
//   	aclName: jsii.String("aclName"),
//
//   	// the properties below are optional
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	userNames: []*string{
//   		jsii.String("userNames"),
//   	},
//   }
//
type CfnACLProps struct {
	// The name of the Access Control List.
	AclName *string `field:"required" json:"aclName" yaml:"aclName"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The list of users that belong to the Access Control List.
	UserNames *[]*string `field:"optional" json:"userNames" yaml:"userNames"`
}

