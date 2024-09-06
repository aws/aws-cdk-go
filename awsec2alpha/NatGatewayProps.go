package awsec2alpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Properties to define a NAT gateway.
//
// Example:
//   myVpc := vpc_v2.NewVpcV2(this, jsii.String("Vpc"))
//   routeTable := vpc_v2.NewRouteTable(this, jsii.String("RouteTable"), &RouteTableProps{
//   	Vpc: myVpc,
//   })
//   subnet := vpc_v2.NewSubnetV2(this, jsii.String("Subnet"), &SubnetV2Props{
//   	Vpc: myVpc,
//   	AvailabilityZone: jsii.String("eu-west-2a"),
//   	Ipv4CidrBlock: awsec2alpha.NewIpCidr(jsii.String("10.0.0.0/24")),
//   	SubnetType: ec2.SubnetType_PRIVATE_ISOLATED,
//   })
//
//   natgw := vpc_v2.NewNatGateway(this, jsii.String("NatGW"), &NatGatewayProps{
//   	Subnet: subnet,
//   	Vpc: myVpc,
//   	ConnectivityType: awsec2alpha.NatConnectivityType_PRIVATE,
//   	PrivateIpAddress: jsii.String("10.0.0.42"),
//   })
//   vpc_v2.NewRoute(this, jsii.String("NatGwRoute"), &RouteProps{
//   	RouteTable: RouteTable,
//   	Destination: jsii.String("0.0.0.0/0"),
//   	Target: map[string]iRouteTarget{
//   		"gateway": natgw,
//   	},
//   })
//
// Experimental.
type NatGatewayProps struct {
	// The subnet in which the NAT gateway is located.
	// Experimental.
	Subnet awsec2.ISubnet `field:"required" json:"subnet" yaml:"subnet"`
	// AllocationID of Elastic IP address that's associated with the NAT gateway.
	//
	// This property is required for a public NAT
	// gateway and cannot be specified with a private NAT gateway.
	// Default: attr.allocationID of a new Elastic IP created by default
	// //TODO: ADD L2 for elastic ip.
	//
	// Experimental.
	AllocationId *string `field:"optional" json:"allocationId" yaml:"allocationId"`
	// Indicates whether the NAT gateway supports public or private connectivity.
	// Default: public.
	//
	// Experimental.
	ConnectivityType NatConnectivityType `field:"optional" json:"connectivityType" yaml:"connectivityType"`
	// The maximum amount of time to wait before forcibly releasing the IP addresses if connections are still in progress.
	// Default: 350 seconds.
	//
	// Experimental.
	MaxDrainDuration awscdk.Duration `field:"optional" json:"maxDrainDuration" yaml:"maxDrainDuration"`
	// The resource name of the NAT gateway.
	// Default: none.
	//
	// Experimental.
	NatGatewayName *string `field:"optional" json:"natGatewayName" yaml:"natGatewayName"`
	// The private IPv4 address to assign to the NAT gateway.
	//
	// If you don't provide an
	// address, a private IPv4 address will be automatically assigned.
	// Default: none.
	//
	// Experimental.
	PrivateIpAddress *string `field:"optional" json:"privateIpAddress" yaml:"privateIpAddress"`
	// Secondary EIP allocation IDs.
	// See: https://docs.aws.amazon.com/vpc/latest/userguide/vpc-nat-gateway.html#nat-gateway-creating
	//
	// Default: none.
	//
	// Experimental.
	SecondaryAllocationIds *[]*string `field:"optional" json:"secondaryAllocationIds" yaml:"secondaryAllocationIds"`
	// The number of secondary private IPv4 addresses you want to assign to the NAT gateway.
	//
	// `SecondaryPrivateIpAddressCount` and `SecondaryPrivateIpAddresses` cannot be
	// set at the same time.
	// See: https://docs.aws.amazon.com/vpc/latest/userguide/vpc-nat-gateway.html#nat-gateway-creating
	//
	// Default: none.
	//
	// Experimental.
	SecondaryPrivateIpAddressCount *float64 `field:"optional" json:"secondaryPrivateIpAddressCount" yaml:"secondaryPrivateIpAddressCount"`
	// Secondary private IPv4 addresses.
	//
	// `SecondaryPrivateIpAddressCount` and `SecondaryPrivateIpAddresses` cannot be
	// set at the same time.
	// See: https://docs.aws.amazon.com/vpc/latest/userguide/vpc-nat-gateway.html#nat-gateway-creating
	//
	// Default: none.
	//
	// Experimental.
	SecondaryPrivateIpAddresses *[]*string `field:"optional" json:"secondaryPrivateIpAddresses" yaml:"secondaryPrivateIpAddresses"`
	// The ID of the VPC in which the NAT gateway is located.
	// Default: none.
	//
	// Experimental.
	Vpc IVpcV2 `field:"optional" json:"vpc" yaml:"vpc"`
}

