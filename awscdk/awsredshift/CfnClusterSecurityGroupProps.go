package awsredshift

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnClusterSecurityGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnClusterSecurityGroupProps := &CfnClusterSecurityGroupProps{
//   	Description: jsii.String("description"),
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
type CfnClusterSecurityGroupProps struct {
	// A description for the security group.
	Description *string `field:"required" json:"description" yaml:"description"`
	// Specifies an arbitrary set of tags (key–value pairs) to associate with this security group.
	//
	// Use tags to manage your resources.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

