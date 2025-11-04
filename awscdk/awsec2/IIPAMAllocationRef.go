package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IPAMAllocation.
// Experimental.
type IIPAMAllocationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a IPAMAllocation resource.
	// Experimental.
	IpamAllocationRef() *IPAMAllocationReference
}

// The jsii proxy for IIPAMAllocationRef
type jsiiProxy_IIPAMAllocationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IIPAMAllocationRef) IpamAllocationRef() *IPAMAllocationReference {
	var returns *IPAMAllocationReference
	_jsii_.Get(
		j,
		"ipamAllocationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIPAMAllocationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIPAMAllocationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

