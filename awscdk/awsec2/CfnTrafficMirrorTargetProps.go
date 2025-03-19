package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnTrafficMirrorTarget`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTrafficMirrorTargetProps := &CfnTrafficMirrorTargetProps{
//   	Description: jsii.String("description"),
//   	GatewayLoadBalancerEndpointId: jsii.String("gatewayLoadBalancerEndpointId"),
//   	NetworkInterfaceId: jsii.String("networkInterfaceId"),
//   	NetworkLoadBalancerArn: jsii.String("networkLoadBalancerArn"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrortarget.html
//
type CfnTrafficMirrorTargetProps struct {
	// The description of the Traffic Mirror target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrortarget.html#cfn-ec2-trafficmirrortarget-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ID of the Gateway Load Balancer endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrortarget.html#cfn-ec2-trafficmirrortarget-gatewayloadbalancerendpointid
	//
	GatewayLoadBalancerEndpointId *string `field:"optional" json:"gatewayLoadBalancerEndpointId" yaml:"gatewayLoadBalancerEndpointId"`
	// The network interface ID that is associated with the target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrortarget.html#cfn-ec2-trafficmirrortarget-networkinterfaceid
	//
	NetworkInterfaceId *string `field:"optional" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// The Amazon Resource Name (ARN) of the Network Load Balancer that is associated with the target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrortarget.html#cfn-ec2-trafficmirrortarget-networkloadbalancerarn
	//
	NetworkLoadBalancerArn *string `field:"optional" json:"networkLoadBalancerArn" yaml:"networkLoadBalancerArn"`
	// The tags to assign to the Traffic Mirror target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrortarget.html#cfn-ec2-trafficmirrortarget-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

