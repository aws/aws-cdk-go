package awsec2alpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awsec2alpha/v2/internal"
)

// Placeholder to see what extra props we might need, will be added to original IVPC.
// Experimental.
type IVpcV2 interface {
	awsec2.IVpc
	// The primary IPv4 CIDR block associated with the VPC.
	//
	// Needed in order to validate the vpc range of subnet
	// current prop vpcCidrBlock refers to the token value
	// For more information, see the {@link https://docs.aws.amazon.com/vpc/latest/userguide/vpc-cidr-blocks.html#vpc-sizing-ipv4}.
	// Experimental.
	Ipv4CidrBlock() *string
	// The secondary CIDR blocks associated with the VPC.
	//
	// For more information, see the {@link https://docs.aws.amazon.com/vpc/latest/userguide/vpc-cidr-blocks.html#vpc-resize}.
	// Experimental.
	SecondaryCidrBlock() *[]awsec2.CfnVPCCidrBlock
}

// The jsii proxy for IVpcV2
type jsiiProxy_IVpcV2 struct {
	internal.Type__awsec2IVpc
}

func (j *jsiiProxy_IVpcV2) Ipv4CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv4CidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpcV2) SecondaryCidrBlock() *[]awsec2.CfnVPCCidrBlock {
	var returns *[]awsec2.CfnVPCCidrBlock
	_jsii_.Get(
		j,
		"secondaryCidrBlock",
		&returns,
	)
	return returns
}

