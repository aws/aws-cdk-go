package awsredshift

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnClusterSubnetGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnClusterSubnetGroupProps := &CfnClusterSubnetGroupProps{
//   	Description: jsii.String("description"),
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
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
type CfnClusterSubnetGroupProps struct {
	// A description for the subnet group.
	Description *string `field:"required" json:"description" yaml:"description"`
	// An array of VPC subnet IDs.
	//
	// A maximum of 20 subnets can be modified in a single request.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// Specifies an arbitrary set of tags (key–value pairs) to associate with this subnet group.
	//
	// Use tags to manage your resources.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

