package awsec2alpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Properties to define VPC [disable-awslint:from-method].
//
// Example:
//   stack := awscdk.Newstack()
//   myVpc := awsec2alpha.NewVpcV2(this, jsii.String("Vpc"), &VpcV2Props{
//   	PrimaryAddressBlock: awsec2alpha.IpAddresses_Ipv4(jsii.String("10.1.0.0/16")),
//   	SecondaryAddressBlocks: []iIpAddresses{
//   		awsec2alpha.IpAddresses_AmazonProvidedIpv6(&SecondaryAddressProps{
//   			CidrBlockName: jsii.String("AmazonProvided"),
//   		}),
//   	},
//   })
//
//   eigw := awsec2alpha.NewEgressOnlyInternetGateway(this, jsii.String("EIGW"), &EgressOnlyInternetGatewayProps{
//   	Vpc: myVpc,
//   })
//
//   routeTable := awsec2alpha.NewRouteTable(this, jsii.String("RouteTable"), &RouteTableProps{
//   	Vpc: myVpc,
//   })
//
//   routeTable.AddRoute(jsii.String("EIGW"), jsii.String("::/0"), map[string]iRouteTarget{
//   	"gateway": eigw,
//   })
//
// Experimental.
type VpcV2Props struct {
	// The default tenancy of instances launched into the VPC.
	//
	// By setting this to dedicated tenancy, instances will be launched on
	// hardware dedicated to a single AWS customer, unless specifically specified
	// at instance launch time. Please note, not all instance types are usable
	// with Dedicated tenancy.
	// Default: DefaultInstanceTenancy.Default (shared) tenancy
	//
	// Experimental.
	DefaultInstanceTenancy awsec2.DefaultInstanceTenancy `field:"optional" json:"defaultInstanceTenancy" yaml:"defaultInstanceTenancy"`
	// Indicates whether the instances launched in the VPC get DNS hostnames.
	// Default: true.
	//
	// Experimental.
	EnableDnsHostnames *bool `field:"optional" json:"enableDnsHostnames" yaml:"enableDnsHostnames"`
	// Indicates whether the DNS resolution is supported for the VPC.
	// Default: true.
	//
	// Experimental.
	EnableDnsSupport *bool `field:"optional" json:"enableDnsSupport" yaml:"enableDnsSupport"`
	// A must IPv4 CIDR block for the VPC.
	// See: https://docs.aws.amazon.com/vpc/latest/userguide/vpc-cidr-blocks.html
	//
	// Default: - Ipv4 CIDR Block ('10.0.0.0/16')
	//
	// Experimental.
	PrimaryAddressBlock IIpAddresses `field:"optional" json:"primaryAddressBlock" yaml:"primaryAddressBlock"`
	// The secondary CIDR blocks associated with the VPC.
	//
	// Can be  IPv4 or IPv6, two IPv4 ranges must follow RFC#1918 convention
	// For more information,.
	// See: https://docs.aws.amazon.com/vpc/latest/userguide/vpc-cidr-blocks.html#vpc-resize}.
	//
	// Default: - No secondary IP address.
	//
	// Experimental.
	SecondaryAddressBlocks *[]IIpAddresses `field:"optional" json:"secondaryAddressBlocks" yaml:"secondaryAddressBlocks"`
	// Physical name for the VPC.
	// Default: - autogenerated by CDK.
	//
	// Experimental.
	VpcName *string `field:"optional" json:"vpcName" yaml:"vpcName"`
}

