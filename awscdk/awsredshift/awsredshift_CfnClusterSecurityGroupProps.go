package awsredshift

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnClusterSecurityGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnClusterSecurityGroupProps := &cfnClusterSecurityGroupProps{
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
type CfnClusterSecurityGroupProps struct {
	// A description for the security group.
	Description *string `field:"required" json:"description" yaml:"description"`
	// Specifies an arbitrary set of tags (key–value pairs) to associate with this security group.
	//
	// Use tags to manage your resources.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

