package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// An abstract Provider of IpAddresses.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var pool cfnIPAMPool
//
//
//   ec2.NewVpc(stack, jsii.String("TheVPC"), &vpcProps{
//   	ipAddresses: ec2.ipAddresses.awsIpamAllocation(&awsIpamProps{
//   		ipv4IpamPoolId: pool.ref,
//   		ipv4NetmaskLength: jsii.Number(18),
//   		defaultSubnetIpv4NetmaskLength: jsii.Number(24),
//   	}),
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
// Uses VPC Cidr allocations from AWS IPAM.
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
// VPC Cidr is supplied at creation and subnets are calculated locally.
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

