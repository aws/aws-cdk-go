package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// An abstract Provider of IpAddresses.
//
// Note this is specific to the IPv4 CIDR.
//
// Example:
//   vpc := ec2.NewVpc(this, jsii.String("TheVPC"), &VpcProps{
//   	// 'IpAddresses' configures the IP range and size of the entire VPC.
//   	// The IP space will be divided based on configuration for the subnets.
//   	IpAddresses: ec2.IpAddresses_Cidr(jsii.String("10.0.0.0/21")),
//
//   	// 'maxAzs' configures the maximum number of availability zones to use.
//   	// If you want to specify the exact availability zones you want the VPC
//   	// to use, use `availabilityZones` instead.
//   	MaxAzs: jsii.Number(3),
//
//   	// 'subnetConfiguration' specifies the "subnet groups" to create.
//   	// Every subnet group will have a subnet for each AZ, so this
//   	// configuration will create `3 groups × 3 AZs = 9` subnets.
//   	SubnetConfiguration: []subnetConfiguration{
//   		&subnetConfiguration{
//   			// 'subnetType' controls Internet access, as described above.
//   			SubnetType: ec2.SubnetType_PUBLIC,
//
//   			// 'name' is used to name this particular subnet group. You will have to
//   			// use the name for subnet selection if you have more than one subnet
//   			// group of the same type.
//   			Name: jsii.String("Ingress"),
//
//   			// 'cidrMask' specifies the IP addresses in the range of of individual
//   			// subnets in the group. Each of the subnets in this group will contain
//   			// `2^(32 address bits - 24 subnet bits) - 2 reserved addresses = 254`
//   			// usable IP addresses.
//   			//
//   			// If 'cidrMask' is left out the available address space is evenly
//   			// divided across the remaining subnet groups.
//   			CidrMask: jsii.Number(24),
//   		},
//   		&subnetConfiguration{
//   			CidrMask: jsii.Number(24),
//   			Name: jsii.String("Application"),
//   			SubnetType: ec2.SubnetType_PRIVATE_WITH_EGRESS,
//   		},
//   		&subnetConfiguration{
//   			CidrMask: jsii.Number(28),
//   			Name: jsii.String("Database"),
//   			SubnetType: ec2.SubnetType_PRIVATE_ISOLATED,
//
//   			// 'reserved' can be used to reserve IP address space. No resources will
//   			// be created for this subnet, but the IP range will be kept available for
//   			// future creation of this subnet, or even for future subdivision.
//   			Reserved: jsii.Boolean(true),
//   		},
//   	},
//   })
//
type IpAddresses interface {
}

// The jsii proxy struct for IpAddresses
type jsiiProxy_IpAddresses struct {
	_ byte // padding
}

// Used to provide centralized Ip Address Management services for your VPC.
//
// Uses VPC CIDR allocations from AWS IPAM
//
// Note this is specific to the IPv4 CIDR.
// See: https://docs.aws.amazon.com/vpc/latest/ipam/what-it-is-ipam.html
//
func IpAddresses_AwsIpamAllocation(props *AwsIpamProps) IIpAddresses {
	_init_.Initialize()

	if err := validateIpAddresses_AwsIpamAllocationParameters(props); err != nil {
		panic(err)
	}
	var returns IIpAddresses

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.IpAddresses",
		"awsIpamAllocation",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Used to provide local Ip Address Management services for your VPC.
//
// VPC CIDR is supplied at creation and subnets are calculated locally
//
// Note this is specific to the IPv4 CIDR.
func IpAddresses_Cidr(cidrBlock *string) IIpAddresses {
	_init_.Initialize()

	if err := validateIpAddresses_CidrParameters(cidrBlock); err != nil {
		panic(err)
	}
	var returns IIpAddresses

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.IpAddresses",
		"cidr",
		[]interface{}{cidrBlock},
		&returns,
	)

	return returns
}

