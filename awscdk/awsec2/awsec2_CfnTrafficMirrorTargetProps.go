package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnTrafficMirrorTarget`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTrafficMirrorTargetProps := &cfnTrafficMirrorTargetProps{
//   	description: jsii.String("description"),
//   	gatewayLoadBalancerEndpointId: jsii.String("gatewayLoadBalancerEndpointId"),
//   	networkInterfaceId: jsii.String("networkInterfaceId"),
//   	networkLoadBalancerArn: jsii.String("networkLoadBalancerArn"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnTrafficMirrorTargetProps struct {
	// The description of the Traffic Mirror target.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ID of the Gateway Load Balancer endpoint.
	GatewayLoadBalancerEndpointId *string `field:"optional" json:"gatewayLoadBalancerEndpointId" yaml:"gatewayLoadBalancerEndpointId"`
	// The network interface ID that is associated with the target.
	NetworkInterfaceId *string `field:"optional" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// The Amazon Resource Name (ARN) of the Network Load Balancer that is associated with the target.
	NetworkLoadBalancerArn *string `field:"optional" json:"networkLoadBalancerArn" yaml:"networkLoadBalancerArn"`
	// The tags to assign to the Traffic Mirror target.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

