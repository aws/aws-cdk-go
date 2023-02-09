package awselasticache

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnSecurityGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSecurityGroupProps := &cfnSecurityGroupProps{
//   	description: jsii.String("description"),
//
//   	// the properties below are optional
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnSecurityGroupProps struct {
	// A description for the cache security group.
	Description *string `field:"required" json:"description" yaml:"description"`
	// A tag that can be added to an ElastiCache security group.
	//
	// Tags are composed of a Key/Value pair. You can use tags to categorize and track all your security groups. A tag with a null Value is permitted.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

