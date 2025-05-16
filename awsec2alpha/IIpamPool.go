package awsec2alpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Definition used to add or create a new IPAM pool.
// Experimental.
type IIpamPool interface {
	// Function to associate a IPv6 address with IPAM pool.
	// Experimental.
	ProvisionCidr(id *string, options *IpamPoolCidrProvisioningOptions) awsec2.CfnIPAMPoolCidr
	// Pool CIDR for IPv6 to be provisioned with Public IP source set to 'Amazon'.
	// Experimental.
	IpamCidrs() *[]awsec2.CfnIPAMPoolCidr
	// Pool CIDR for IPv4 to be provisioned using IPAM Required to check for subnet IP range is within the VPC range.
	// Experimental.
	IpamIpv4Cidrs() *[]*string
	// Pool ID to be passed to the VPC construct.
	// Experimental.
	IpamPoolId() *string
}

// The jsii proxy for IIpamPool
type jsiiProxy_IIpamPool struct {
	_ byte // padding
}

func (i *jsiiProxy_IIpamPool) ProvisionCidr(id *string, options *IpamPoolCidrProvisioningOptions) awsec2.CfnIPAMPoolCidr {
	if err := i.validateProvisionCidrParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsec2.CfnIPAMPoolCidr

	_jsii_.Invoke(
		i,
		"provisionCidr",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IIpamPool) IpamCidrs() *[]awsec2.CfnIPAMPoolCidr {
	var returns *[]awsec2.CfnIPAMPoolCidr
	_jsii_.Get(
		j,
		"ipamCidrs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIpamPool) IpamIpv4Cidrs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipamIpv4Cidrs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIpamPool) IpamPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipamPoolId",
		&returns,
	)
	return returns
}

