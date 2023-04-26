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
//   cfnNatGatewayProps := &CfnNatGatewayProps{
//   	SubnetId: jsii.String("subnetId"),
//
//   	// the properties below are optional
//   	AllocationId: jsii.String("allocationId"),
//   	ConnectivityType: jsii.String("connectivityType"),
//   	MaxDrainDurationSeconds: jsii.Number(123),
//   	PrivateIpAddress: jsii.String("privateIpAddress"),
//   	SecondaryAllocationIds: []*string{
//   		jsii.String("secondaryAllocationIds"),
//   	},
//   	SecondaryPrivateIpAddressCount: jsii.Number(123),
//   	SecondaryPrivateIpAddresses: []*string{
//   		jsii.String("secondaryPrivateIpAddresses"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
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
	// The maximum amount of time to wait (in seconds) before forcibly releasing the IP addresses if connections are still in progress.
	//
	// Default value is 350 seconds.
	MaxDrainDurationSeconds *float64 `field:"optional" json:"maxDrainDurationSeconds" yaml:"maxDrainDurationSeconds"`
	// The private IPv4 address to assign to the NAT gateway.
	//
	// If you don't provide an address, a private IPv4 address will be automatically assigned.
	PrivateIpAddress *string `field:"optional" json:"privateIpAddress" yaml:"privateIpAddress"`
	// Secondary EIP allocation IDs.
	//
	// For more information, see [Create a NAT gateway](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-nat-gateway.html#nat-gateway-creating) in the *Amazon VPC User Guide* .
	SecondaryAllocationIds *[]*string `field:"optional" json:"secondaryAllocationIds" yaml:"secondaryAllocationIds"`
	// [Private NAT gateway only] The number of secondary private IPv4 addresses you want to assign to the NAT gateway.
	//
	// For more information about secondary addresses, see [Create a NAT gateway](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-nat-gateway.html#nat-gateway-creating) in the *Amazon Virtual Private Cloud User Guide* .
	//
	// > `SecondaryPrivateIpAddressCount` and `SecondaryPrivateIpAddresses` cannot be set at the same time.
	SecondaryPrivateIpAddressCount *float64 `field:"optional" json:"secondaryPrivateIpAddressCount" yaml:"secondaryPrivateIpAddressCount"`
	// Secondary private IPv4 addresses.
	//
	// For more information about secondary addresses, see [Create a NAT gateway](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-nat-gateway.html#nat-gateway-creating) in the *Amazon Virtual Private Cloud User Guide* .
	//
	// > `SecondaryPrivateIpAddressCount` and `SecondaryPrivateIpAddresses` cannot be set at the same time.
	SecondaryPrivateIpAddresses *[]*string `field:"optional" json:"secondaryPrivateIpAddresses" yaml:"secondaryPrivateIpAddresses"`
	// The tags for the NAT gateway.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

