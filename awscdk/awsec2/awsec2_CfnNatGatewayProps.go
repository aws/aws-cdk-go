package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnNatGateway`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnNatGatewayProps := &cfnNatGatewayProps{
//   	subnetId: jsii.String("subnetId"),
//
//   	// the properties below are optional
//   	allocationId: jsii.String("allocationId"),
//   	connectivityType: jsii.String("connectivityType"),
//   	privateIpAddress: jsii.String("privateIpAddress"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnNatGatewayProps struct {
	// The ID of the subnet in which the NAT gateway is located.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// [Public NAT gateway only] The allocation ID of the Elastic IP address that's associated with the NAT gateway.
	//
	// This property is required for a public NAT gateway and cannot be specified with a private NAT gateway.
	AllocationId *string `field:"optional" json:"allocationId" yaml:"allocationId"`
	// Indicates whether the NAT gateway supports public or private connectivity.
	//
	// The default is public connectivity.
	ConnectivityType *string `field:"optional" json:"connectivityType" yaml:"connectivityType"`
	// The private IPv4 address to assign to the NAT gateway.
	//
	// If you don't provide an address, a private IPv4 address will be automatically assigned.
	PrivateIpAddress *string `field:"optional" json:"privateIpAddress" yaml:"privateIpAddress"`
	// The tags for the NAT gateway.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

