package awsec2alpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Options to define EgressOnlyInternetGateway for VPC.
//
// Example:
//   stack := awscdk.Newstack()
//   myVpc := awsec2alpha.NewVpcV2(this, jsii.String("Vpc"), &VpcV2Props{
//   	PrimaryAddressBlock: awsec2alpha.IpAddresses_Ipv4(jsii.String("10.1.0.0/16")),
//   	SecondaryAddressBlocks: []IIpAddresses{
//   		awsec2alpha.IpAddresses_AmazonProvidedIpv6(&SecondaryAddressProps{
//   			CidrBlockName: jsii.String("AmazonProvided"),
//   		}),
//   	},
//   })
//   routeTable := awsec2alpha.NewRouteTable(this, jsii.String("RouteTable"), &RouteTableProps{
//   	Vpc: myVpc,
//   })
//   subnet := awsec2alpha.NewSubnetV2(this, jsii.String("Subnet"), &SubnetV2Props{
//   	Vpc: myVpc,
//   	AvailabilityZone: jsii.String("eu-west-2a"),
//   	Ipv4CidrBlock: awsec2alpha.NewIpCidr(jsii.String("10.0.0.0/24")),
//   	Ipv6CidrBlock: awsec2alpha.NewIpCidr(jsii.String("2001:db8:1::/64")),
//   	SubnetType: awscdk.SubnetType_PRIVATE,
//   })
//
//   myVpc.AddEgressOnlyInternetGateway(&EgressOnlyInternetGatewayOptions{
//   	Subnets: []SubnetSelection{
//   		&SubnetSelection{
//   			SubnetType: awscdk.SubnetType_PRIVATE,
//   		},
//   	},
//   	Destination: jsii.String("::/60"),
//   })
//
// Experimental.
type EgressOnlyInternetGatewayOptions struct {
	// Destination Ipv6 address for EGW route.
	// Default: - '::/0' all Ipv6 traffic.
	//
	// Experimental.
	Destination *string `field:"optional" json:"destination" yaml:"destination"`
	// The resource name of the egress-only internet gateway.
	//
	// Provided name will be used for tagging.
	// Default: - no name tag associated and provisioned without a resource name.
	//
	// Experimental.
	EgressOnlyInternetGatewayName *string `field:"optional" json:"egressOnlyInternetGatewayName" yaml:"egressOnlyInternetGatewayName"`
	// List of subnets where route to EGW will be added.
	// Default: - no route created.
	//
	// Experimental.
	Subnets *[]*awsec2.SubnetSelection `field:"optional" json:"subnets" yaml:"subnets"`
}

