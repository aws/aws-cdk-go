package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnLocalGatewayRouteTable`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLocalGatewayRouteTableProps := &CfnLocalGatewayRouteTableProps{
//   	LocalGatewayId: jsii.String("localGatewayId"),
//
//   	// the properties below are optional
//   	Mode: jsii.String("mode"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnLocalGatewayRouteTableProps struct {
	// The ID of the local gateway.
	LocalGatewayId *string `field:"required" json:"localGatewayId" yaml:"localGatewayId"`
	// The mode of the local gateway route table.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// The tags assigned to the local gateway route table.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

