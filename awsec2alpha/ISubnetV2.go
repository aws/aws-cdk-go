package awsec2alpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awsec2alpha/v2/internal"
)

// Interface with additional properties for SubnetV2.
// Experimental.
type ISubnetV2 interface {
	awsec2.ISubnet
	// The IPv6 CIDR block for this subnet.
	// Experimental.
	Ipv6CidrBlock() *string
}

// The jsii proxy for ISubnetV2
type jsiiProxy_ISubnetV2 struct {
	internal.Type__awsec2ISubnet
}

func (j *jsiiProxy_ISubnetV2) Ipv6CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6CidrBlock",
		&returns,
	)
	return returns
}

