package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnRouteTable`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRouteTableProps := &cfnRouteTableProps{
//   	vpcId: jsii.String("vpcId"),
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
type CfnRouteTableProps struct {
	// The ID of the VPC.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// Any tags assigned to the route table.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

