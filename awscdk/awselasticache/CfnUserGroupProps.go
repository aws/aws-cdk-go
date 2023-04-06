package awselasticache

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnUserGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnUserGroupProps := &CfnUserGroupProps{
//   	Engine: jsii.String("engine"),
//   	UserGroupId: jsii.String("userGroupId"),
//   	UserIds: []*string{
//   		jsii.String("userIds"),
//   	},
//
//   	// the properties below are optional
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnUserGroupProps struct {
	// The current supported value is redis.
	Engine *string `field:"required" json:"engine" yaml:"engine"`
	// The ID of the user group.
	UserGroupId *string `field:"required" json:"userGroupId" yaml:"userGroupId"`
	// The list of user IDs that belong to the user group.
	//
	// A user named `default` must be included.
	UserIds *[]*string `field:"required" json:"userIds" yaml:"userIds"`
	// `AWS::ElastiCache::UserGroup.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

